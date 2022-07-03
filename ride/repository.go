package ride

import "context"

type InMemoryRideRepository struct {
}

func (repo *InMemoryRideRepository) Find(ctx context.Context, rideId string) (*Ride, error) {
	return nil, nil
}

func (repo *InMemoryRideRepository) Create(ctx context.Context, ride *Ride) error {
	return nil
}

func (repo *InMemoryRideRepository) Save(ctx context.Context, ride *Ride) error {
	return nil
}
