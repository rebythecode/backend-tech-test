package ride

import (
	"context"

	"go.uber.org/zap"
)

type RideRepository interface {
	Find(ctx context.Context, rideId string) (*Ride, error)
	FindActiveByUserOrVehicle(ctx context.Context, userID, vehicleID string) (*Ride, error)
	Create(ctx context.Context, ride *Ride) error
}

type RideApplication struct {
	unlockPrice int
	minutePrice int
	repo        RideRepository
	logger      *zap.Logger
}

func NewRideApplication(unlockPrice, minutePrice int, repo RideRepository, logger *zap.Logger) *RideApplication {
	return &RideApplication{
		unlockPrice: unlockPrice,
		minutePrice: minutePrice,
		repo:        repo,
		logger:      logger,
	}
}

func (app *RideApplication) Start(userID, vehicleID string) (*Ride, error) {
	return nil, nil
}

func (app *RideApplication) Finish(rideID string) (*Ride, error) {
	return nil, nil
}
