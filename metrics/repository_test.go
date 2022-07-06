package metrics

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetrieveMetricsByEndpoint(t *testing.T) {
	t.Parallel()

	repo := NewInMemoryMetricsRepository()

	trace1 := NewTrace("Example1")
	endtime1 := trace1.start.Add(10 * time.Second)
	trace1.end = &endtime1
	if err := repo.Create(context.TODO(), trace1); err != nil {
		t.Fatalf("got error when creating trace: %s", err.Error())
	}

	trace2 := NewTrace("Example1")
	endtime2 := trace2.start.Add(50 * time.Second)
	trace2.end = &endtime2
	trace2.err = errors.New("example")
	if err := repo.Create(context.TODO(), trace2); err != nil {
		t.Fatalf("got error when creating trace: %s", err.Error())
	}

	trace3 := NewTrace("Example2")
	endtime3 := trace3.start.Add(4 * time.Second)
	trace3.end = &endtime3
	if err := repo.Create(context.TODO(), trace3); err != nil {
		t.Fatalf("got error when creating trace: %s", err.Error())
	}

	trace4 := NewTrace("Example3")
	if err := repo.Create(context.TODO(), trace4); err != nil {
		t.Fatalf("got error when creating trace: %s", err.Error())
	}

	allMetrics, err := repo.RetrieveMetricsByEndpoint(context.TODO())
	if err != nil {
		t.Fatalf("got error when retrieving metrics: %s", err.Error())
	}

	if len(allMetrics) != 2 {
		t.Fatalf("got metrics lenght = %d, want = %d", len(allMetrics), 2)
	}

	metric, exists := allMetrics["Example1"]
	if !exists {
		t.Fatalf("got no metrics for endpoint Example1")
	}

	var nano int64 = 1000000000
	if metric.avg/nano != 30 {
		t.Fatalf("got avg = %d, want = %d", metric.avg, 30)
	}

	if metric.min/nano != 10 {
		t.Fatalf("got min = %d, want = %d", metric.min, 10)
	}

	if metric.max/nano != 50 {
		t.Fatalf("got max = %d, want = %d", metric.max, 50)
	}

	if metric.success != 1 {
		t.Fatalf("got success = %d, want = %d", metric.success, 2)
	}

	if metric.failed != 1 {
		t.Fatalf("got failed = %d, want = %d", metric.failed, 0)
	}

	metric, exists = allMetrics["Example2"]
	if !exists {
		t.Fatalf("got no metrics for endpoint Example2")
	}
}
