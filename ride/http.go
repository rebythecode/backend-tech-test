package ride

import (
	"encoding/json"
	"errors"
	"net/http"

	back "github.com/HectorMRC/backend-tech-test"
	"go.uber.org/zap"
)

type RideHttpBody struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	VehicleID string `json:"vehicle_id"`
	Cost      int    `json:"cost,omitempty"`
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
	var reqBody RideHttpBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ride, err := server.app.Start(r.Context(), reqBody.UserID, reqBody.VehicleID)
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
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ride, err := server.app.Finish(r.Context(), reqBody.ID)
	if errors.Is(err, back.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
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
