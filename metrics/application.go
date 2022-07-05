package metrics

import (
	"context"

	"go.uber.org/zap"
)

type MetricsRepository interface {
	Create(ctx context.Context, t *Trace) error
	RetrieveMetricsByEndpoint(ctx context.Context) (map[string]*Metrics, error)
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

func (app *MetricsApplication) StartMetricsTrace(ctx context.Context, endpoint string) (*Trace, error) {
	trace := NewTrace(endpoint)
	if err := app.repo.Create(ctx, trace); err != nil {
		return nil, err
	}

	return trace, nil
}

func (app *MetricsApplication) RetrieveMetricsByEndpoint(ctx context.Context) (map[string]*Metrics, error) {
	app.logger.Info("processing a \"metrics\" request")

	return app.repo.RetrieveMetricsByEndpoint(ctx)
}
