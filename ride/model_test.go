package ride

import (
	"errors"
	"testing"
	"time"

	back "github.com/HectorMRC/backend-tech-test"
)

func TestSetEndtime(t *testing.T) {
	t.Parallel()
	ride := NewRide("TestUser", "TestVehicle", time.Now())
	want := time.Now()
	if err := ride.SetEndtime(want); err != nil {
		t.Fatalf("got error when setting endtime: %s", err.Error())
	}

	if got := ride.end; !got.Equal(want) {
		t.Fatalf("got endtime = %v, want = %v", got, want)
	}
}

func TestSetEndtime_whenAlreadyFinished(t *testing.T) {
	t.Parallel()
	ride := NewRide("TestUser", "TestVehicle", time.Now())

	endtime := time.Now()
	ride.end = &endtime

	if err := ride.SetEndtime(endtime); !errors.Is(err, back.ErrNotAvailable) {
		t.Fatalf("got error = %s, want = %s", err.Error(), back.ErrNotAvailable)
	}
}

func TestSetCost(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		baseCost   int
		minuteCost int
		duration   time.Duration
		want       int
	}{
		{"set cost when no minutes", 100, 18, 0, 100},
		{"set cost when only one minute", 100, 18, 1 * time.Minute, 118},
		{"set cost when five minutes", 100, 18, 5 * time.Minute, 190},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ride := NewRide("TestUser", "TestVehicle", time.Now())
			if err := ride.SetEndtime(ride.start.Add(tt.duration)); err != nil {
				t.Fatalf("got error when setting endtime: %s", err.Error())
			}

			if err := ride.SetCost(tt.baseCost, tt.minuteCost); err != nil {
				t.Fatalf("got error when setting cost: %s", err.Error())
			}

			if got := ride.cost; got != tt.want {
				t.Fatalf("got cost = %d, want = %d", got, tt.want)
			}
		})
	}
}

func TestSetCost_whenNoFinished(t *testing.T) {
	t.Parallel()
	ride := NewRide("TestUser", "TestVehicle", time.Now())
	if err := ride.SetCost(100, 18); !errors.Is(err, back.ErrNotAvailable) {
		t.Fatalf("got error = %s, want = %s", err.Error(), back.ErrNotAvailable)
	}
}
