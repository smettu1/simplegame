// island/island.go
package island

import "github.com/smettu1/game/enums"

// Edge represents a connection between islands.
type Edge struct {
	AllowedPaths enums.ConnectionType
	Destination  *Island
}

// Island represents an island in the game.
type Island struct {
	Name  string
	Paths []*Edge
}

// NewIsland creates a new island with the given name.
func NewIsland(name string) *Island {
	return &Island{
		Name:  name,
		Paths: make([]*Edge, 0),
	}
}

// AddConnection adds a connection to another island.
func (i *Island) AddConnection(conn *Edge) {
	i.Paths = append(i.Paths, conn)
}
