package ride

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"

	back "github.com/HectorMRC/backend-tech-test"
)

type InMemoryRideRepository struct {
	instances map[string]*Ride
	sequence  int64
	mu        sync.RWMutex
}

func NewInMemoryRideRepository() *InMemoryRideRepository {
	return &InMemoryRideRepository{
		instances: make(map[string]*Ride),
	}
}

func (repo *InMemoryRideRepository) Find(ctx context.Context, id string) (*Ride, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if ride, exists := repo.instances[id]; exists {
		return ride, nil
	}

	return nil, back.ErrNotFound
}

func (repo *InMemoryRideRepository) FindActiveByUserOrVehicle(ctx context.Context, userID, vehicleID string) (*Ride, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, ride := range repo.instances {
		if ride.end == nil &&
			(ride.userID == userID || ride.vehicleID == vehicleID) {
			return ride, nil
		}
	}

	return nil, back.ErrNotFound
}

func (repo *InMemoryRideRepository) Create(ctx context.Context, ride *Ride) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	id := atomic.AddInt64(&repo.sequence, 1)
	ride.id = strconv.FormatInt(id, 10)
	repo.instances[ride.id] = ride
	return nil
}
