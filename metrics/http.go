package metrics

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type MetricsHttpBody struct {
	Metrics []struct {
		Endpoint string        `json:"endpoint"`
		Success  uint          `json:"success"`
		Failed   uint          `json:"failed"`
		Min      time.Duration `json:"min"`
		Max      time.Duration `json:"max"`
		Avg      time.Duration `json:"average"`
	} `json:"metrics"`
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
}
