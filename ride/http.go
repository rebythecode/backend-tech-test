package ride

import (
	"net/http"

	"go.uber.org/zap"
)

type RideHttpBody struct {
	UserID    string  `json:"user_id"`
	VehicleID string  `json:"vehicle_id"`
	Cost      float32 `json:"cost,omitempty"`
}

type RideServer struct {
	app    *RideApplication
	logger *zap.Logger
}

func NewRideServer(app *RideApplication, logger *zap.Logger) *RideServer {
	return &RideServer{
		app:    app,
		logger: logger,
	}
}

func (server *RideServer) RideStartHandler(w http.ResponseWriter, r *http.Request) {
}

func (server *RideServer) RideFinishHandler(w http.ResponseWriter, r *http.Request) {

}
