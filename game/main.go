package main

import (
	"fmt"

	"github.com/smettu1/game/game"
	"github.com/smettu1/game/island"
	"github.com/smettu1/game/player"
	"github.com/smettu1/game/vehicle"
)

func main() {
	// Create islands
	startIsland := island.NewIsland("Start Island")
	endIsland := island.NewIsland("End Island")
	startIsland.AddConnection(&island.Edge{Destination: endIsland})
	// Create players with their chosen vehicles
	player1 := player.NewPlayer("Player 1", &vehicle.HotBalloon{}, startIsland)
	player2 := player.NewPlayer("Player 2", &vehicle.Car{}, startIsland)
	// Initialize the game with players and islands
	g := game.NewGame([]*player.Player{player1, player2}, startIsland, endIsland)
	fmt.Println("Game initialized!")
	// Add player1 and player2 steps below.
	g.MovePlayer(player1, startIsland)
}
