// enums/enums.go
package enums

// enums for ConnectionType and VehicleType
type ConnectionType int

const (
	Flight ConnectionType = 1
	Bridge ConnectionType = 2
)

type VehicleType int

const (
	HotBalloon VehicleType = 1
	Car        VehicleType = 2
)
