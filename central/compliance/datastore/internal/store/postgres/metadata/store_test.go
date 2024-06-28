// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ComplianceRunMetadataStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceRunMetadataStore(t *testing.T) {
	suite.Run(t, new(ComplianceRunMetadataStoreSuite))
}

func (s *ComplianceRunMetadataStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceRunMetadataStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_run_metadata CASCADE")
	s.T().Log("compliance_run_metadata", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ComplianceRunMetadataStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceRunMetadataStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceRunMetadata := &storage.ComplianceRunMetadata{}
	s.NoError(testutils.FullInit(complianceRunMetadata, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceRunMetadata, exists, err := store.Get(ctx, complianceRunMetadata.GetRunId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceRunMetadata)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceRunMetadata))
	foundComplianceRunMetadata, exists, err = store.Get(ctx, complianceRunMetadata.GetRunId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), complianceRunMetadata, foundComplianceRunMetadata)

	complianceRunMetadataCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, complianceRunMetadataCount)
	complianceRunMetadataCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(complianceRunMetadataCount)

	complianceRunMetadataExists, err := store.Exists(ctx, complianceRunMetadata.GetRunId())
	s.NoError(err)
	s.True(complianceRunMetadataExists)
	s.NoError(store.Upsert(ctx, complianceRunMetadata))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceRunMetadata), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, complianceRunMetadata.GetRunId()))
	foundComplianceRunMetadata, exists, err = store.Get(ctx, complianceRunMetadata.GetRunId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceRunMetadata)
	s.NoError(store.Delete(withNoAccessCtx, complianceRunMetadata.GetRunId()))

	var complianceRunMetadatas []*storage.ComplianceRunMetadata
	var complianceRunMetadataIDs []string
	for i := 0; i < 200; i++ {
		complianceRunMetadata := &storage.ComplianceRunMetadata{}
		s.NoError(testutils.FullInit(complianceRunMetadata, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceRunMetadatas = append(complianceRunMetadatas, complianceRunMetadata)
		complianceRunMetadataIDs = append(complianceRunMetadataIDs, complianceRunMetadata.GetRunId())
	}

	s.NoError(store.UpsertMany(ctx, complianceRunMetadatas))

	complianceRunMetadataCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, complianceRunMetadataCount)

	s.NoError(store.DeleteMany(ctx, complianceRunMetadataIDs))

	complianceRunMetadataCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, complianceRunMetadataCount)
}

const (
	withAllAccess                = "AllAccess"
	withNoAccess                 = "NoAccess"
	withAccess                   = "Access"
	withAccessToCluster          = "AccessToCluster"
	withNoAccessToCluster        = "NoAccessToCluster"
	withAccessToDifferentCluster = "AccessToDifferentCluster"
	withAccessToDifferentNs      = "AccessToDifferentNs"
)

var (
	withAllAccessCtx = sac.WithAllAccess(context.Background())
)

type testCase struct {
	context                context.Context
	expectedObjIDs         []string
	expectedIdentifiers    []string
	expectedMissingIndices []int
	expectedObjects        []*storage.ComplianceRunMetadata
	expectedWriteError     error
}

func (s *ComplianceRunMetadataStoreSuite) getTestData(access ...storage.Access) (*storage.ComplianceRunMetadata, *storage.ComplianceRunMetadata, map[string]testCase) {
	objA := &storage.ComplianceRunMetadata{}
	s.NoError(testutils.FullInit(objA, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	objB := &storage.ComplianceRunMetadata{}
	s.NoError(testutils.FullInit(objB, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))

	testCases := map[string]testCase{
		withAllAccess: {
			context:                sac.WithAllAccess(context.Background()),
			expectedObjIDs:         []string{objA.GetRunId(), objB.GetRunId()},
			expectedIdentifiers:    []string{objA.GetRunId(), objB.GetRunId()},
			expectedMissingIndices: []int{},
			expectedObjects:        []*storage.ComplianceRunMetadata{objA, objB},
			expectedWriteError:     nil,
		},
		withNoAccess: {
			context:                sac.WithNoAccess(context.Background()),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceRunMetadata{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withNoAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(uuid.Nil.String()),
				)),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceRunMetadata{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccess: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
				)),
			expectedObjIDs:         []string{objA.GetRunId()},
			expectedIdentifiers:    []string{objA.GetRunId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceRunMetadata{objA},
			expectedWriteError:     nil,
		},
		withAccessToCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
				)),
			expectedObjIDs:         []string{objA.GetRunId()},
			expectedIdentifiers:    []string{objA.GetRunId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceRunMetadata{objA},
			expectedWriteError:     nil,
		},
		withAccessToDifferentCluster: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys("caaaaaaa-bbbb-4011-0000-111111111111"),
				)),
			expectedObjIDs:         []string{},
			expectedIdentifiers:    []string{},
			expectedMissingIndices: []int{0, 1},
			expectedObjects:        []*storage.ComplianceRunMetadata{},
			expectedWriteError:     sac.ErrResourceAccessDenied,
		},
		withAccessToDifferentNs: {
			context: sac.WithGlobalAccessScopeChecker(context.Background(),
				sac.AllowFixedScopes(
					sac.AccessModeScopeKeys(access...),
					sac.ResourceScopeKeys(targetResource),
					sac.ClusterScopeKeys(objA.GetClusterId()),
					sac.NamespaceScopeKeys("unknown ns"),
				)),
			expectedObjIDs:         []string{objA.GetRunId()},
			expectedIdentifiers:    []string{objA.GetRunId()},
			expectedMissingIndices: []int{1},
			expectedObjects:        []*storage.ComplianceRunMetadata{objA},
			expectedWriteError:     nil,
		},
	}

	return objA, objB, testCases
}

func (s *ComplianceRunMetadataStoreSuite) TestSACUpsert() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.Upsert(testCase.context, obj), testCase.expectedWriteError)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACUpsertMany() {
	obj, _, testCases := s.getTestData(storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			assert.ErrorIs(t, s.store.UpsertMany(testCase.context, []*storage.ComplianceRunMetadata{obj}), testCase.expectedWriteError)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACCount() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			expectedCount := len(testCase.expectedObjects)
			count, err := s.store.Count(testCase.context, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, expectedCount, count)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACWalk() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers := []string{}
			getIDs := func(obj *storage.ComplianceRunMetadata) error {
				identifiers = append(identifiers, obj.GetRunId())
				return nil
			}
			err := s.store.Walk(testCase.context, getIDs)
			assert.NoError(t, err)
			assert.ElementsMatch(t, testCase.expectedIdentifiers, identifiers)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACGetIDs() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			identifiers, err := s.store.GetIDs(testCase.context)
			assert.NoError(t, err)
			assert.ElementsMatch(t, testCase.expectedObjIDs, identifiers)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACExists() {
	objA, _, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			exists, err := s.store.Exists(testCase.context, objA.GetRunId())
			assert.NoError(t, err)

			// Assumption from the test case structure: objA is always in the visible list
			// in the first position.
			expectedFound := len(testCase.expectedObjects) > 0
			assert.Equal(t, expectedFound, exists)
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACGet() {
	objA, _, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, exists, err := s.store.Get(testCase.context, objA.GetRunId())
			assert.NoError(t, err)

			// Assumption from the test case structure: objA is always in the visible list
			// in the first position.
			expectedFound := len(testCase.expectedObjects) > 0
			assert.Equal(t, expectedFound, exists)
			if expectedFound {
				protoassert.Equal(t, objA, actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACDelete() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS)

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.Delete(testCase.context, objA.GetRunId()))
			assert.NoError(t, s.store.Delete(testCase.context, objB.GetRunId()))

			count, err := s.store.Count(withAllAccessCtx, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, 2-len(testCase.expectedObjects), count)

			// Ensure objects allowed by test scope were actually deleted
			for _, obj := range testCase.expectedObjects {
				found, err := s.store.Exists(withAllAccessCtx, obj.GetRunId())
				assert.NoError(t, err)
				assert.False(t, found)
			}
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACDeleteMany() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS)
	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			s.SetupTest()

			s.NoError(s.store.Upsert(withAllAccessCtx, objA))
			s.NoError(s.store.Upsert(withAllAccessCtx, objB))

			assert.NoError(t, s.store.DeleteMany(testCase.context, []string{
				objA.GetRunId(),
				objB.GetRunId(),
			}))

			count, err := s.store.Count(withAllAccessCtx, search.EmptyQuery())
			assert.NoError(t, err)
			assert.Equal(t, 2-len(testCase.expectedObjects), count)

			// Ensure objects allowed by test scope were actually deleted
			for _, obj := range testCase.expectedObjects {
				found, err := s.store.Exists(withAllAccessCtx, obj.GetRunId())
				assert.NoError(t, err)
				assert.False(t, found)
			}
		})
	}
}

func (s *ComplianceRunMetadataStoreSuite) TestSACGetMany() {
	objA, objB, testCases := s.getTestData(storage.Access_READ_ACCESS)
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objA))
	s.Require().NoError(s.store.Upsert(withAllAccessCtx, objB))

	for name, testCase := range testCases {
		s.T().Run(fmt.Sprintf("with %s", name), func(t *testing.T) {
			actual, missingIndices, err := s.store.GetMany(testCase.context, []string{objA.GetRunId(), objB.GetRunId()})
			assert.NoError(t, err)
			protoassert.SlicesEqual(t, testCase.expectedObjects, actual)
			assert.Equal(t, testCase.expectedMissingIndices, missingIndices)
		})
	}

	s.T().Run("with no identifiers", func(t *testing.T) {
		actual, missingIndices, err := s.store.GetMany(withAllAccessCtx, []string{})
		assert.Nil(t, err)
		assert.Nil(t, actual)
		assert.Nil(t, missingIndices)
	})
}
