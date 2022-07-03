package ride

import (
	"context"

	"go.uber.org/zap"
)

type RideRepository interface {
	Find(ctx context.Context, rideId string) (*Ride, error)
	Create(ctx context.Context, ride *Ride) error
	Save(ctx context.Context, ride *Ride) error
}

type DirectoryApplication struct {
	repo   RideRepository
	logger *zap.Logger
}
