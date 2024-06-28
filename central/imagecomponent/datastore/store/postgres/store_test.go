// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type ImageComponentsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestImageComponentsStore(t *testing.T) {
	suite.Run(t, new(ImageComponentsStoreSuite))
}

func (s *ImageComponentsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ImageComponentsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE image_components CASCADE")
	s.T().Log("image_components", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ImageComponentsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ImageComponentsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	imageComponent := &storage.ImageComponent{}
	s.NoError(testutils.FullInit(imageComponent, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundImageComponent, exists, err := store.Get(ctx, imageComponent.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundImageComponent)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, imageComponent))
	foundImageComponent, exists, err = store.Get(ctx, imageComponent.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), imageComponent, foundImageComponent)

	imageComponentCount, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, imageComponentCount)
	imageComponentCount, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(imageComponentCount)

	imageComponentExists, err := store.Exists(ctx, imageComponent.GetId())
	s.NoError(err)
	s.True(imageComponentExists)
	s.NoError(store.Upsert(ctx, imageComponent))
	s.ErrorIs(store.Upsert(withNoAccessCtx, imageComponent), sac.ErrResourceAccessDenied)

	s.NoError(store.Delete(ctx, imageComponent.GetId()))
	foundImageComponent, exists, err = store.Get(ctx, imageComponent.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundImageComponent)
	s.NoError(store.Delete(withNoAccessCtx, imageComponent.GetId()))

	var imageComponents []*storage.ImageComponent
	var imageComponentIDs []string
	for i := 0; i < 200; i++ {
		imageComponent := &storage.ImageComponent{}
		s.NoError(testutils.FullInit(imageComponent, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		imageComponents = append(imageComponents, imageComponent)
		imageComponentIDs = append(imageComponentIDs, imageComponent.GetId())
	}

	s.NoError(store.UpsertMany(ctx, imageComponents))

	imageComponentCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, imageComponentCount)

	s.NoError(store.DeleteMany(ctx, imageComponentIDs))

	imageComponentCount, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, imageComponentCount)
}
