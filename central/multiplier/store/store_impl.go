package store

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/dberrors"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/secondarykey"
	"github.com/stackrox/rox/pkg/uuid"
)

type storeImpl struct {
	*bolt.DB
}

func (b *storeImpl) getMultiplier(id string, bucket *bolt.Bucket) (multiplier *v1.Multiplier, exists bool, err error) {
	multiplier = new(v1.Multiplier)
	val := bucket.Get([]byte(id))
	if val == nil {
		return
	}
	exists = true
	err = proto.Unmarshal(val, multiplier)
	return
}

// GetMultiplier returns multiplier with given id.
func (b *storeImpl) GetMultiplier(id string) (multiplier *v1.Multiplier, exists bool, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "Multiplier")
	err = b.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(multiplierBucket))
		multiplier, exists, err = b.getMultiplier(id, bucket)
		return err
	})
	return
}

// GetMultipliers retrieves multipliers from bolt
func (b *storeImpl) GetMultipliers() ([]*v1.Multiplier, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetMany, "Multiplier")
	var multipliers []*v1.Multiplier
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(multiplierBucket))
		return b.ForEach(func(k, v []byte) error {
			var multiplier v1.Multiplier
			if err := proto.Unmarshal(v, &multiplier); err != nil {
				return err
			}
			multipliers = append(multipliers, &multiplier)
			return nil
		})
	})
	return multipliers, err
}

// AddMultiplier adds a multiplier into bolt
func (b *storeImpl) AddMultiplier(multiplier *v1.Multiplier) (string, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Add, "Multiplier")
	multiplier.Id = uuid.NewV4().String()
	err := b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(multiplierBucket))
		_, exists, err := b.getMultiplier(multiplier.GetId(), bucket)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("Multiplier %s (%s) cannot be added because it already exists", multiplier.GetId(), multiplier.GetName())
		}
		if err := secondarykey.CheckUniqueKeyExistsAndInsert(tx, multiplierBucket, multiplier.GetId(), multiplier.GetName()); err != nil {
			return fmt.Errorf("Could not add multiplier due to name validation: %s", err)
		}
		bytes, err := proto.Marshal(multiplier)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(multiplier.GetId()), bytes)
	})
	return multiplier.Id, err
}

// UpdateMultiplier upserts a multiplier into bolt
func (b *storeImpl) UpdateMultiplier(multiplier *v1.Multiplier) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Update, "Multiplier")
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(multiplierBucket))
		// If the update is changing the name, check if the name has already been taken
		if val, _ := secondarykey.GetCurrentUniqueKey(tx, multiplierBucket, multiplier.GetId()); val != multiplier.GetName() {
			if err := secondarykey.UpdateUniqueKey(tx, multiplierBucket, multiplier.GetId(), multiplier.GetName()); err != nil {
				return fmt.Errorf("Could not update multiplier due to name validation: %s", err)
			}
		}
		bytes, err := proto.Marshal(multiplier)
		if err != nil {
			return err
		}
		return b.Put([]byte(multiplier.GetId()), bytes)
	})
}

// RemoveMultiplier removes a multiplier from bolt
func (b *storeImpl) RemoveMultiplier(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Remove, "Multiplier")
	return b.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(multiplierBucket))
		key := []byte(id)
		if exists := b.Get(key) != nil; !exists {
			return dberrors.ErrNotFound{Type: "Multiplier", ID: string(key)}
		}
		if err := secondarykey.RemoveUniqueKey(tx, multiplierBucket, id); err != nil {
			return err
		}
		return b.Delete(key)
	})
}
