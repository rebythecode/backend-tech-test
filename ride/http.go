package ride

import (
	"encoding/json"
	"errors"
	"net/http"

	back "github.com/HectorMRC/backend-tech-test"
	"github.com/HectorMRC/backend-tech-test/metrics"
	"go.uber.org/zap"
)

type RideHttpBody struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	VehicleID string `json:"vehicle_id"`
	Cost      int    `json:"cost,omitempty"`
}

type RideServer struct {
	rideApp    *RideApplication
	metricsApp *metrics.MetricsApplication
	logger     *zap.Logger
}

func NewRideServer(rideApp *RideApplication, metricsApp *metrics.MetricsApplication, logger *zap.Logger) *RideServer {
	return &RideServer{
		rideApp:    rideApp,
		metricsApp: metricsApp,
		logger:     logger,
	}
}

func (server *RideServer) RideStartHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RideHttpBody
	var err error

	trace, err := server.metricsApp.StartMetricsTrace(r.Context(), "StartRide")
	if err == nil {
		defer trace.Finish(&err)
	} else {
		server.logger.Error("creating trace",
			zap.Error(err))
	}

	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ride, err := server.rideApp.Start(r.Context(), reqBody.UserID, reqBody.VehicleID)
	if errors.Is(err, back.ErrNotAvailable) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody := &RideHttpBody{
		ID:        ride.id,
		UserID:    ride.userID,
		VehicleID: ride.vehicleID,
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

func (server *RideServer) RideFinishHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RideHttpBody
	var err error

	trace, err := server.metricsApp.StartMetricsTrace(r.Context(), "FinishRide")
	if err == nil {
		defer trace.Finish(&err)
	} else {
		server.logger.Error("creating trace",
			zap.Error(err))
	}

	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ride, err := server.rideApp.Finish(r.Context(), reqBody.ID)
	if errors.Is(err, back.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if errors.Is(err, ErrAlreadyFinished) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody := &RideHttpBody{
		ID:        ride.id,
		UserID:    ride.userID,
		VehicleID: ride.vehicleID,
		Cost:      ride.cost,
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
