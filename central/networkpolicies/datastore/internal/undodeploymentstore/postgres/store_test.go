// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type NetworkpoliciesundodeploymentsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestNetworkpoliciesundodeploymentsStore(t *testing.T) {
	suite.Run(t, new(NetworkpoliciesundodeploymentsStoreSuite))
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.Require().NoError(err)

	Destroy(ctx, pool)

	s.pool = pool
	s.store = New(ctx, pool)
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) TearDownTest() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *NetworkpoliciesundodeploymentsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	networkPolicyApplicationUndoDeploymentRecord := &storage.NetworkPolicyApplicationUndoDeploymentRecord{}
	s.NoError(testutils.FullInit(networkPolicyApplicationUndoDeploymentRecord, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err := store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoDeploymentRecord)

	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoDeploymentRecord))
	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoDeploymentRecord, foundNetworkPolicyApplicationUndoDeploymentRecord)

	networkPolicyApplicationUndoDeploymentRecordCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, networkPolicyApplicationUndoDeploymentRecordCount)

	networkPolicyApplicationUndoDeploymentRecordExists, err := store.Exists(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(networkPolicyApplicationUndoDeploymentRecordExists)
	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoDeploymentRecord))

	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoDeploymentRecord, foundNetworkPolicyApplicationUndoDeploymentRecord)

	s.NoError(store.Delete(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId()))
	foundNetworkPolicyApplicationUndoDeploymentRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoDeploymentRecord.GetDeploymentId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoDeploymentRecord)

	var networkPolicyApplicationUndoDeploymentRecords []*storage.NetworkPolicyApplicationUndoDeploymentRecord
	for i := 0; i < 200; i++ {
		networkPolicyApplicationUndoDeploymentRecord := &storage.NetworkPolicyApplicationUndoDeploymentRecord{}
		s.NoError(testutils.FullInit(networkPolicyApplicationUndoDeploymentRecord, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkPolicyApplicationUndoDeploymentRecords = append(networkPolicyApplicationUndoDeploymentRecords, networkPolicyApplicationUndoDeploymentRecord)
	}

	s.NoError(store.UpsertMany(ctx, networkPolicyApplicationUndoDeploymentRecords))

	networkPolicyApplicationUndoDeploymentRecordCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, networkPolicyApplicationUndoDeploymentRecordCount)
}
