package metrics

import "time"

type Trace struct {
	endpoint string
	err      error
	start    time.Time
	end      *time.Time
}

type Metrics struct {
	success int
	failed  int
	min     time.Duration
	max     time.Duration
	avg     time.Duration
}

func NewTrace(endpoint string) *Trace {
	return &Trace{
		endpoint: endpoint,
		start:    time.Now(),
	}
}

func (trace *Trace) Finish(err *error) {
	if err != nil {
		trace.err = *err
	}

	endTime := time.Now()
	trace.end = &endTime
}
