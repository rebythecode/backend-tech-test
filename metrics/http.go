package metrics

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type endpointMetrics struct {
	Success int   `json:"success"`
	Failed  int   `json:"failed"`
	Min     int64 `json:"min"`
	Max     int64 `json:"max"`
	Avg     int64 `json:"average"`
}

type MetricsHttpBody struct {
	Metrics map[string]endpointMetrics `json:"metrics"`
}

type MetricsServer struct {
	app    *MetricsApplication
	logger *zap.Logger
}

func NewMetricsServer(app *MetricsApplication, logger *zap.Logger) *MetricsServer {
	return &MetricsServer{
		app:    app,
		logger: logger,
	}
}

func (server *MetricsServer) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	trace, err := server.app.StartMetricsTrace(r.Context(), "Metrics")
	if err == nil {
		defer trace.Finish(&err)
	} else {
		server.logger.Error("creating trace",
			zap.Error(err))
	}

	allMetrics, err := server.app.RetrieveMetricsByEndpoint(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody := &MetricsHttpBody{
		Metrics: make(map[string]endpointMetrics),
	}

	for endpoint, metrics := range allMetrics {
		respBody.Metrics[endpoint] = endpointMetrics{
			Success: metrics.success,
			Failed:  metrics.failed,
			Min:     metrics.min,
			Max:     metrics.max,
			Avg:     metrics.avg,
		}
	}

	resp, err := json.Marshal(respBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(resp); err != nil {
		server.logger.Error("writing http response",
			zap.Error(err))
	}
}
