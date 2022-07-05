package ride

import (
	"context"
	"errors"

	back "github.com/HectorMRC/backend-tech-test"
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

func (app *RideApplication) Start(ctx context.Context, userID, vehicleID string) (*Ride, error) {
	app.logger.Info("processing a \"start\" request",
		zap.String("user_id", userID),
		zap.String("vehicle_id", vehicleID))

	if _, err := app.repo.FindActiveByUserOrVehicle(ctx, userID, vehicleID); !errors.Is(err, back.ErrNotFound) {
		return nil, back.ErrNotAvailable
	}

	newRide := NewRide(userID, vehicleID)
	if err := app.repo.Create(ctx, newRide); err != nil {
		return nil, err
	}

	return newRide, nil
}

func (app *RideApplication) Finish(ctx context.Context, rideID string) (*Ride, error) {
	app.logger.Info("processing a \"finish\" request",
		zap.String("ride_id", rideID))

	ride, err := app.repo.Find(ctx, rideID)
	if err != nil {
		return nil, err
	}

	if err := ride.Finish(app.unlockPrice, app.minutePrice); err != nil {
		return nil, err
	}

	return ride, nil
}
