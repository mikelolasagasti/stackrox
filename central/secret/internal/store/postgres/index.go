// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"reflect"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	metrics "github.com/stackrox/rox/central/metrics"
	mappings "github.com/stackrox/rox/central/secret/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/walker"
	search "github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
	"github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_SECRETS, walker.Walk(reflect.TypeOf((*storage.Secret)(nil)), baseTable))
}

func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "Secret")

	return postgres.RunCountRequest(v1.SearchCategory_SECRETS, q, b.db, mappings.OptionsMap)
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "Secret")

	return postgres.RunSearchRequest(v1.SearchCategory_SECRETS, q, b.db, mappings.OptionsMap)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddSecret(deployment *storage.Secret) error {
	return nil
}

func (b *indexerImpl) AddSecrets(_ []*storage.Secret) error {
	return nil
}

func (b *indexerImpl) DeleteSecret(id string) error {
	return nil
}

func (b *indexerImpl) DeleteSecrets(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
