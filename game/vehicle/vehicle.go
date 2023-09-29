// vehicle/vehicle.go

// vehicle/vehicle.go

package vehicle

import (
	"errors"

	"github.com/smettu1/game/enums"
	"github.com/smettu1/game/island"
)

type Vehicle interface {
	CanMove(destination *island.Island) bool
	MoveTo(destination *island.Island) error
}

type HotBalloon struct{}

func (h HotBalloon) CanMove(destination *island.Island) bool {
	// Hot air balloons can move to any island that is on a flight path.
	for _, edge := range destination.Paths {
		if edge.AllowedPaths == enums.Flight {
			return true
		}
	}
	return false
}

func (h HotBalloon) MoveTo(destination *island.Island) error {
	// Check if the hot air balloon can move to the destination.
	if h.CanMove(destination) {
		return nil // Successfully moved
	}
	return errors.New("Hot air balloon cannot move to this island")
}

type Car struct{}

func (c Car) CanMove(destination *island.Island) bool {
	// Cars can move to any island connected by a bridge.
	for _, edge := range destination.Paths {
		if edge.AllowedPaths == enums.Bridge {
			return true
		}
	}
	return false
}

func (c Car) MoveTo(destination *island.Island) error {
	// Check if the car can move to the destination.
	if c.CanMove(destination) {
		return nil // Successfully moved
	}
	return errors.New("Car cannot move to this island")
}
