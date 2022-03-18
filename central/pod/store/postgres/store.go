// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
)

const (
	baseTable  = "pods"
	countStmt  = "SELECT COUNT(*) FROM pods"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM pods WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM pods WHERE Id = $1"
	deleteStmt  = "DELETE FROM pods WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM pods"
	getIDsStmt  = "SELECT Id FROM pods"
	getManyStmt = "SELECT serialized FROM pods WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM pods WHERE Id = ANY($1::text[])"
)

func init() {
	globaldb.RegisterTable(baseTable, "Pod")
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.Pod, bool, error)
	Upsert(ctx context.Context, obj *storage.Pod) error
	UpsertMany(ctx context.Context, objs []*storage.Pod) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.Pod, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.Pod) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTablePods(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists pods (
    Id varchar,
    Name varchar,
    DeploymentId varchar,
    Namespace varchar,
    ClusterId varchar,
    Started timestamp,
    serialized bytea,
    PRIMARY KEY(Id)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			panic(err)
		}
	}

	createTablePodsLiveInstances(ctx, db)
	createTablePodsTerminatedInstances(ctx, db)
}

func createTablePodsLiveInstances(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists pods_LiveInstances (
    pods_Id varchar,
    idx integer,
    InstanceId_ContainerRuntime integer,
    InstanceId_Id varchar,
    InstanceId_Node varchar,
    ContainingPodId varchar,
    ContainerName varchar,
    ContainerIps text[],
    Started timestamp,
    ImageDigest varchar,
    Finished timestamp,
    ExitCode integer,
    TerminationReason varchar,
    PRIMARY KEY(pods_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (pods_Id) REFERENCES pods(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists podsLiveInstances_idx on pods_LiveInstances using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			panic(err)
		}
	}

}

func createTablePodsTerminatedInstances(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists pods_TerminatedInstances (
    pods_Id varchar,
    idx integer,
    PRIMARY KEY(pods_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (pods_Id) REFERENCES pods(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists podsTerminatedInstances_idx on pods_TerminatedInstances using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			panic(err)
		}
	}

	createTablePodsTerminatedInstancesInstances(ctx, db)
}

func createTablePodsTerminatedInstancesInstances(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists pods_TerminatedInstances_Instances (
    pods_Id varchar,
    pods_TerminatedInstances_idx integer,
    idx integer,
    InstanceId_ContainerRuntime integer,
    InstanceId_Id varchar,
    InstanceId_Node varchar,
    ContainingPodId varchar,
    ContainerName varchar,
    ContainerIps text[],
    Started timestamp,
    ImageDigest varchar,
    Finished timestamp,
    ExitCode integer,
    TerminationReason varchar,
    PRIMARY KEY(pods_Id, pods_TerminatedInstances_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (pods_Id, pods_TerminatedInstances_idx) REFERENCES pods_TerminatedInstances(pods_Id, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists podsTerminatedInstancesInstances_idx on pods_TerminatedInstances_Instances using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			panic(err)
		}
	}

}

func insertIntoPods(ctx context.Context, tx pgx.Tx, obj *storage.Pod) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetName(),
		obj.GetDeploymentId(),
		obj.GetNamespace(),
		obj.GetClusterId(),
		pgutils.NilOrStringTimestamp(obj.GetStarted()),
		serialized,
	}

	finalStr := "INSERT INTO pods (Id, Name, DeploymentId, Namespace, ClusterId, Started, serialized) VALUES($1, $2, $3, $4, $5, $6, $7) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, DeploymentId = EXCLUDED.DeploymentId, Namespace = EXCLUDED.Namespace, ClusterId = EXCLUDED.ClusterId, Started = EXCLUDED.Started, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetLiveInstances() {
		if err := insertIntoPodsLiveInstances(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from pods_LiveInstances where pods_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetLiveInstances()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetTerminatedInstances() {
		if err := insertIntoPodsTerminatedInstances(ctx, tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from pods_TerminatedInstances where pods_Id = $1 AND idx >= $2"
	_, err = tx.Exec(ctx, query, obj.GetId(), len(obj.GetTerminatedInstances()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoPodsLiveInstances(ctx context.Context, tx pgx.Tx, obj *storage.ContainerInstance, pods_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pods_Id,
		idx,
		obj.GetInstanceId().GetContainerRuntime(),
		obj.GetInstanceId().GetId(),
		obj.GetInstanceId().GetNode(),
		obj.GetContainingPodId(),
		obj.GetContainerName(),
		obj.GetContainerIps(),
		pgutils.NilOrStringTimestamp(obj.GetStarted()),
		obj.GetImageDigest(),
		pgutils.NilOrStringTimestamp(obj.GetFinished()),
		obj.GetExitCode(),
		obj.GetTerminationReason(),
	}

	finalStr := "INSERT INTO pods_LiveInstances (pods_Id, idx, InstanceId_ContainerRuntime, InstanceId_Id, InstanceId_Node, ContainingPodId, ContainerName, ContainerIps, Started, ImageDigest, Finished, ExitCode, TerminationReason) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT(pods_Id, idx) DO UPDATE SET pods_Id = EXCLUDED.pods_Id, idx = EXCLUDED.idx, InstanceId_ContainerRuntime = EXCLUDED.InstanceId_ContainerRuntime, InstanceId_Id = EXCLUDED.InstanceId_Id, InstanceId_Node = EXCLUDED.InstanceId_Node, ContainingPodId = EXCLUDED.ContainingPodId, ContainerName = EXCLUDED.ContainerName, ContainerIps = EXCLUDED.ContainerIps, Started = EXCLUDED.Started, ImageDigest = EXCLUDED.ImageDigest, Finished = EXCLUDED.Finished, ExitCode = EXCLUDED.ExitCode, TerminationReason = EXCLUDED.TerminationReason"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoPodsTerminatedInstances(ctx context.Context, tx pgx.Tx, obj *storage.Pod_ContainerInstanceList, pods_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pods_Id,
		idx,
	}

	finalStr := "INSERT INTO pods_TerminatedInstances (pods_Id, idx) VALUES($1, $2) ON CONFLICT(pods_Id, idx) DO UPDATE SET pods_Id = EXCLUDED.pods_Id, idx = EXCLUDED.idx"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetInstances() {
		if err := insertIntoPodsTerminatedInstancesInstances(ctx, tx, child, pods_Id, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from pods_TerminatedInstances_Instances where pods_Id = $1 AND pods_TerminatedInstances_idx = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, pods_Id, idx, len(obj.GetInstances()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoPodsTerminatedInstancesInstances(ctx context.Context, tx pgx.Tx, obj *storage.ContainerInstance, pods_Id string, pods_TerminatedInstances_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pods_Id,
		pods_TerminatedInstances_idx,
		idx,
		obj.GetInstanceId().GetContainerRuntime(),
		obj.GetInstanceId().GetId(),
		obj.GetInstanceId().GetNode(),
		obj.GetContainingPodId(),
		obj.GetContainerName(),
		obj.GetContainerIps(),
		pgutils.NilOrStringTimestamp(obj.GetStarted()),
		obj.GetImageDigest(),
		pgutils.NilOrStringTimestamp(obj.GetFinished()),
		obj.GetExitCode(),
		obj.GetTerminationReason(),
	}

	finalStr := "INSERT INTO pods_TerminatedInstances_Instances (pods_Id, pods_TerminatedInstances_idx, idx, InstanceId_ContainerRuntime, InstanceId_Id, InstanceId_Node, ContainingPodId, ContainerName, ContainerIps, Started, ImageDigest, Finished, ExitCode, TerminationReason) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) ON CONFLICT(pods_Id, pods_TerminatedInstances_idx, idx) DO UPDATE SET pods_Id = EXCLUDED.pods_Id, pods_TerminatedInstances_idx = EXCLUDED.pods_TerminatedInstances_idx, idx = EXCLUDED.idx, InstanceId_ContainerRuntime = EXCLUDED.InstanceId_ContainerRuntime, InstanceId_Id = EXCLUDED.InstanceId_Id, InstanceId_Node = EXCLUDED.InstanceId_Node, ContainingPodId = EXCLUDED.ContainingPodId, ContainerName = EXCLUDED.ContainerName, ContainerIps = EXCLUDED.ContainerIps, Started = EXCLUDED.Started, ImageDigest = EXCLUDED.ImageDigest, Finished = EXCLUDED.Finished, ExitCode = EXCLUDED.ExitCode, TerminationReason = EXCLUDED.TerminationReason"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTablePods(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.Pod) error {
	conn, release := s.acquireConn(ctx, ops.Get, "Pod")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoPods(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.Pod) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "Pod")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.Pod) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "Pod")

	return s.upsert(ctx, objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "Pod")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "Pod")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.Pod, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "Pod")

	conn, release := s.acquireConn(ctx, ops.Get, "Pod")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.Pod
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "Pod")

	conn, release := s.acquireConn(ctx, ops.Remove, "Pod")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.PodIDs")

	rows, err := s.db.Query(ctx, getIDsStmt)
	if err != nil {
		return nil, pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.Pod, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "Pod")

	conn, release := s.acquireConn(ctx, ops.GetMany, "Pod")
	defer release()

	rows, err := conn.Query(ctx, getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	elems := make([]*storage.Pod, 0, len(ids))
	foundSet := make(map[string]struct{})
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		var msg storage.Pod
		if err := proto.Unmarshal(data, &msg); err != nil {
			return nil, nil, err
		}
		foundSet[msg.GetId()] = struct{}{}
		elems = append(elems, &msg)
	}
	missingIndices := make([]int, 0, len(ids)-len(foundSet))
	for i, id := range ids {
		if _, ok := foundSet[id]; !ok {
			missingIndices = append(missingIndices, i)
		}
	}
	return elems, missingIndices, nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ctx context.Context, ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "Pod")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "Pod")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.Pod) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.Pod
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTablePods(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS pods CASCADE")
	dropTablePodsLiveInstances(ctx, db)
	dropTablePodsTerminatedInstances(ctx, db)

}

func dropTablePodsLiveInstances(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS pods_LiveInstances CASCADE")

}

func dropTablePodsTerminatedInstances(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS pods_TerminatedInstances CASCADE")
	dropTablePodsTerminatedInstancesInstances(ctx, db)

}

func dropTablePodsTerminatedInstancesInstances(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS pods_TerminatedInstances_Instances CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTablePods(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
