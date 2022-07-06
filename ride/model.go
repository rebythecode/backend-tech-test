package ride

import (
	"time"

	back "github.com/HectorMRC/backend-tech-test"
)

type Ride struct {
	id        string
	userID    string
	vehicleID string
	start     time.Time
	end       *time.Time
	cost      int
}

func NewRide(userID, vehicleID string) *Ride {
	return &Ride{
		userID:    userID,
		vehicleID: vehicleID,
		start:     time.Now(),
	}
}

func (ride *Ride) SetEndtime(t time.Time) error {
	if ride.end != nil {
		return back.ErrNotAvailable
	}

	ride.end = &t
	return nil
}

func (ride *Ride) SetCost(baseCost, minuteCost int) error {
	if ride.end == nil {
		return back.ErrNotAvailable
	}

	rideDuration := ride.end.Sub(ride.start)
	ride.cost = baseCost + minuteCost*int(rideDuration.Minutes())
	return nil
}
