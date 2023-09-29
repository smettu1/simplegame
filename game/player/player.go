// player/player.go
package player

import (
	"github.com/smettu1/game/enums"
	"github.com/smettu1/game/island"
	"github.com/smettu1/game/vehicle"
)

// Player represents a player in the game.
type Player struct {
	Name            string
	ModeOfTransport vehicle.Vehicle
	MyMove          bool
	NumberOfMoves   int
	CurrentIsland   *island.Island
	TypeOfPath      enums.ConnectionType
}

// NewPlayer creates a new player with the specified name and vehicle and island.
func NewPlayer(name string, transport vehicle.Vehicle, island *island.Island) *Player {
	return &Player{
		Name:            name,
		ModeOfTransport: transport,
		MyMove:          false,
		NumberOfMoves:   0,
		CurrentIsland:   island,
	}
}
