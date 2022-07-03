package ride

const (
	RIDE_STATUS_ACTIVE uint8 = iota
	RIDE_STATUS_FINISHED
)

type Ride struct {
	user    string
	vehicle string
	status  uint8
	cost    uint32
}
