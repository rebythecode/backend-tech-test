package ride

import (
	"context"
	"errors"
	"testing"
	"time"

	back "github.com/HectorMRC/backend-tech-test"
	"go.uber.org/zap"
)

type RideRepositoryMock struct {
	find                      func(ctx context.Context, rideId string) (*Ride, error)
	findActiveByUserOrVehicle func(ctx context.Context, userID, vehicleID string) (*Ride, error)
	create                    func(ctx context.Context, ride *Ride) error
}

func (mock *RideRepositoryMock) Find(ctx context.Context, rideId string) (*Ride, error) {
	if mock.find != nil {
		return mock.find(ctx, rideId)
	}

	return nil, back.ErrNotFound
}

func (mock *RideRepositoryMock) FindActiveByUserOrVehicle(ctx context.Context, userID, vehicleID string) (*Ride, error) {
	if mock.findActiveByUserOrVehicle != nil {
		return mock.findActiveByUserOrVehicle(ctx, userID, vehicleID)
	}

	return nil, back.ErrNotFound
}

func (mock *RideRepositoryMock) Create(ctx context.Context, ride *Ride) error {
	if mock.create != nil {
		return mock.create(ctx, ride)
	}

	return nil
}

func TestStart(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := &RideRepositoryMock{}
	app := NewRideApplication(0, 0, repo, logger)

	before := time.Now().Unix()
	ride, err := app.Start(context.TODO(), "TestUser", "TestVehicle")
	after := time.Now().Unix()

	if err != nil {
		t.Fatalf("got error when starting ride: %s", err.Error())
	}

	if got := ride.start.Unix(); before > got || after < got {
		t.Fatalf("got start = %v, want > %v && < %v", got, before, after)
	}
}

func TestStart_whenAlreadyOnRide(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := &RideRepositoryMock{
		findActiveByUserOrVehicle: func(ctx context.Context, userID, vehicleID string) (*Ride, error) {
			return nil, nil
		},
	}

	app := NewRideApplication(0, 0, repo, logger)
	if _, err := app.Start(context.TODO(), "TestUser", "TestVehicle"); !errors.Is(err, back.ErrNotAvailable) {
		t.Fatalf("got error = %s, want = %s", err.Error(), back.ErrNotAvailable)
	}
}

func TestFinish(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := &RideRepositoryMock{
		find: func(ctx context.Context, id string) (*Ride, error) {
			return NewRide("TestUser", "TestVehicle", time.Now()), nil
		},
	}

	app := NewRideApplication(100, 0, repo, logger)

	before := time.Now().Unix()
	ride, err := app.Finish(context.TODO(), "TestRide")
	after := time.Now().Unix()

	if err != nil {
		t.Fatalf("got error when finishing ride: %s", err.Error())
	}

	if ride.end == nil {
		t.Fatal("got no endtime")
	}

	if got := ride.end.Unix(); before > got || after < got {
		t.Fatalf("got end = %v, want > %v && < %v", got, before, after)
	}

	if ride.cost != 100 {
		t.Fatalf("got cost = %d, want = %d", ride.cost, 100)
	}
}

func TestFinish_whenRideDoesNotExists(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := &RideRepositoryMock{}
	app := NewRideApplication(100, 0, repo, logger)
	if _, err := app.Finish(context.TODO(), "TestRide"); !errors.Is(err, back.ErrNotFound) {
		t.Fatalf("got error = %s, want = %s", err.Error(), back.ErrNotFound)
	}
}

func TestFinish_whenRideAlreadyFinished(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo := &RideRepositoryMock{
		find: func(ctx context.Context, id string) (*Ride, error) {
			ride := NewRide("TestUser", "TestVehicle", time.Now())
			ride.SetEndtime(time.Now())
			return ride, nil
		},
	}

	app := NewRideApplication(100, 0, repo, logger)
	if _, err := app.Finish(context.TODO(), "TestRide"); !errors.Is(err, back.ErrNotAvailable) {
		t.Fatalf("got error = %s, want = %s", err.Error(), back.ErrNotFound)
	}
}
