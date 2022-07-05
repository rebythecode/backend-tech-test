package metrics

import "time"

type Metrics struct {
	endpoint string
	err      error
	start    time.Time
	end      time.Time
}

type EndpointMetrics struct {
	endpoint string
	success  int64
	failed   int64
	min      time.Duration
	max      time.Duration
	avg      time.Duration
}
