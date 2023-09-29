package game

import (
	"github.com/smettu1/game/enums"
	"github.com/smettu1/game/island"
	"github.com/smettu1/game/player"
)

type Game struct {
	Players     []*player.Player
	StartIsland *island.Island
	EndIsland   *island.Island
	Winner      *player.Player
}

func NewGame(players []*player.Player, startIsland *island.Island, endIsland *island.Island) *Game {
	return &Game{
		Players:     players,
		StartIsland: startIsland,
		EndIsland:   endIsland,
	}
}

func (g *Game) MovePlayer(player *player.Player, destination *island.Island) bool {
	// Check if the player's current island has a valid path to the destination using the chosen vehicle.
	for _, edge := range player.CurrentIsland.Paths {
		if edge.Destination == destination && edge.AllowedPaths == enums.ConnectionType(player.TypeOfPath) {
			// Move the player to the destination island.
			player.CurrentIsland = destination
			player.NumberOfMoves++
			return true
		}
	}
	return false
}

func (g *Game) PossibleMovesForPlayer(player *player.Player) []*island.Island {
	// Find all the islands connected to the player's current island using their chosen vehicle.
	possibleMoves := []*island.Island{}
	for _, edge := range player.CurrentIsland.Paths {
		if edge.AllowedPaths == enums.ConnectionType(player.TypeOfPath) {
			possibleMoves = append(possibleMoves, edge.Destination)
		}
	}
	return possibleMoves
}

func (g *Game) IsRaceFinished() bool {
	// Check if any player has reached the end island.
	for _, player := range g.Players {
		if player.CurrentIsland == g.EndIsland {
			g.Winner = player
			return true
		}
	}
	return false
}
