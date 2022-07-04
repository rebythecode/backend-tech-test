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

func (repo *InMemoryMetricsRepository) Create(ctx context.Context, l *Metrics) error {
	repo.instances.Store(l, nil)
	return nil
}

func (repo *InMemoryMetricsRepository) Average(ctx context.Context) (*EndpointMetrics, error) {
	return nil, nil
}
