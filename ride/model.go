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

func (ride *Ride) Finish(baseCost, minuteCost int) error {
	if ride.end != nil {
		return back.ErrNotAvailable
	}

	endTime := time.Now()
	rideDuration := endTime.Sub(ride.start)

	ride.end = &endTime
	ride.cost = baseCost + minuteCost*int(rideDuration.Minutes())
	return nil
}
