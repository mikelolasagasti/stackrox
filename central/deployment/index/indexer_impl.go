package index

import (
	"time"

	"github.com/blevesearch/bleve"
	"github.com/stackrox/rox/central/deployment/index/mappings"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/api/v1"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/blevesearch"
)

// AlertIndex provides storage functionality for alerts.
type indexerImpl struct {
	index bleve.Index
}

type deploymentWrapper struct {
	*v1.Deployment `json:"deployment"`
	Type           string `json:"type"`
}

// AddDeployment adds the deployment to the index
func (b *indexerImpl) AddDeployment(deployment *v1.Deployment) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "Deployment")
	return b.index.Index(deployment.GetId(), &deploymentWrapper{Type: v1.SearchCategory_DEPLOYMENTS.String(), Deployment: deployment})
}

// AddDeployments adds the deployments to the index
func (b *indexerImpl) AddDeployments(deployments []*v1.Deployment) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "Deployment")

	batch := b.index.NewBatch()
	for _, deployment := range deployments {
		batch.Index(deployment.GetId(), &deploymentWrapper{Type: v1.SearchCategory_DEPLOYMENTS.String(), Deployment: deployment})
	}
	return b.index.Batch(batch)
}

// DeleteDeployment deletes the deployment from the index
func (b *indexerImpl) DeleteDeployment(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "Deployment")
	return b.index.Delete(id)
}

// Search takes a Query and finds any matches
func (b *indexerImpl) Search(q *v1.Query) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "Deployment")
	return blevesearch.RunSearchRequest(v1.SearchCategory_DEPLOYMENTS, q, b.index, mappings.OptionsMap)
}
