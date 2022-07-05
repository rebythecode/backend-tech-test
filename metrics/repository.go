package metrics

import (
	"context"
	"sync"
)

type InMemoryMetricsRepository struct {
	instances sync.Map
}

func NewInMemoryMetricsRepository() *InMemoryMetricsRepository {
	return &InMemoryMetricsRepository{}
}

func (repo *InMemoryMetricsRepository) Save(ctx context.Context, m *Metrics) error {
	repo.instances.Store(m, nil)
	return nil
}

func (repo *InMemoryMetricsRepository) Retrieve(ctx context.Context) ([]*EndpointMetrics, error) {
	return nil, nil
}
