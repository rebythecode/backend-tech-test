package metrics

import (
	"context"
	"sync"
	"time"
)

type InMemoryMetricsRepository struct {
	instances sync.Map
}

func NewInMemoryMetricsRepository() *InMemoryMetricsRepository {
	return &InMemoryMetricsRepository{}
}

func (repo *InMemoryMetricsRepository) Create(ctx context.Context, trace *Trace) error {
	repo.instances.Store(trace, nil)
	return nil
}

func (repo *InMemoryMetricsRepository) RetrieveMetricsByEndpoint(ctx context.Context) (map[string]*Metrics, error) {
	allMetrics := make(map[string]*Metrics)
	repo.instances.Range(func(v, _ any) bool {
		trace := v.(*Trace)
		if trace.end == nil {
			return true
		}

		duration := trace.end.Sub(trace.start)

		metrics, exists := allMetrics[trace.endpoint]
		if !exists {
			metrics = &Metrics{
				min: duration,
				max: duration,
			}

			allMetrics[trace.endpoint] = metrics
		}

		metrics.avg += duration
		if trace.err != nil {
			metrics.failed++
		} else {
			metrics.success++
		}

		if metrics.min > duration {
			metrics.min = duration
		} else if metrics.max < duration {
			metrics.max = duration
		}

		return true
	})

	for _, metrics := range allMetrics {
		metrics.avg /= time.Duration(metrics.failed + metrics.success)
	}

	return allMetrics, nil
}
