package ride

const (
	RIDE_STATUS_ACTIVE uint8 = iota
	RIDE_STATUS_FINISHED
)

type Ride struct {
	id        string
	userID    string
	vehicleID string
	cost      uint
	status    uint8
}
