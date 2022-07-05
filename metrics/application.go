package metrics

import (
	"context"

	"go.uber.org/zap"
)

type MetricsRepository interface {
	Save(ctx context.Context, l *Metrics) error
	Retrieve(ctx context.Context) ([]*EndpointMetrics, error)
}

type MetricsApplication struct {
	repo   MetricsRepository
	logger *zap.Logger
}

func NewMetricsApplication(repo MetricsRepository, logger *zap.Logger) *MetricsApplication {
	return &MetricsApplication{
		repo:   repo,
		logger: logger,
	}
}
