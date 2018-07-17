package store

import (
	"time"

	"bitbucket.org/stack-rox/apollo/central/metrics"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
)

type storeImpl struct {
	*bolt.DB
}

// GetNamespace returns namespace with given id.
func (b *storeImpl) GetNamespace(id string) (namespace *v1.Namespace, exists bool, err error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Get", "Namespace")
	namespace = new(v1.Namespace)
	err = b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(namespaceBucket))
		val := b.Get([]byte(id))
		if val == nil {
			return nil
		}
		exists = true
		return proto.Unmarshal(val, namespace)
	})

	return
}

// GetNamespaces retrieves namespaces matching the request from bolt
func (b *storeImpl) GetNamespaces() ([]*v1.Namespace, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "GetMany", "Namespace")
	var namespaces []*v1.Namespace
	err := b.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(namespaceBucket))
		return b.ForEach(func(k, v []byte) error {
			var namespace v1.Namespace
			if err := proto.Unmarshal(v, &namespace); err != nil {
				return err
			}
			namespaces = append(namespaces, &namespace)
			return nil
		})
	})
	return namespaces, err
}

// AddNamespace adds a namespace to bolt
func (b *storeImpl) AddNamespace(namespace *v1.Namespace) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Update", "Namespace")
	return b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(namespaceBucket))
		bytes, err := proto.Marshal(namespace)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(namespace.GetId()), bytes)
	})
}

// UpdateNamespace updates a namespace to bolt
func (b *storeImpl) UpdateNamespace(namespace *v1.Namespace) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Update", "Namespace")
	return b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(namespaceBucket))
		bytes, err := proto.Marshal(namespace)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(namespace.GetId()), bytes)
	})
}

// RemoveNamespace removes a namespace.
func (b *storeImpl) RemoveNamespace(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), "Remove", "Namespace")
	return b.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(namespaceBucket))
		return bucket.Delete([]byte(id))
	})
}
