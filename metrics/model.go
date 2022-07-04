package metrics

import "time"

type Metrics struct {
	endpoint string
	err      error
	start    time.Duration
	end      time.Duration
}

type EndpointMetrics struct {
	endpoint string
	success  uint
	failed   uint
	min      time.Duration
	max      time.Duration
	avg      time.Duration
}
