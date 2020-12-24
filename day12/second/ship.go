package main

// Direction indicates the direction the ship is facing, given the compass rose
type Direction int

const (
	// North faces the ship north
	North Direction = 0
	// East faces the ship east
	East Direction = 1
	// South faces the ship south
	South Direction = 2
	// West faces the ship west
	West Direction = 3
)

// Orientation is used when turning the ship, either left or right
type Orientation int

const (
	// Left orientation indicates turning the ship left
	Left Orientation = iota
	// Right orientiation indicates turning the ship right
	Right
)

// Ship represents the ship on the map
type Ship struct {
	Direction          Direction
	Waypoint           *Waypoint
	EastWestDistance   float64
	NorthSouthDistance float64
}

// New returns a pointer to a new Ship struct
func New() *Ship {
	return &Ship{
		Direction:          East,
		Waypoint:           NewWaypoint(),
		EastWestDistance:   0,
		NorthSouthDistance: 0,
	}
}

// MoveForward will move the ship to the waypoint, times the multiplier
func (s *Ship) MoveForward(multiplier float64) {
	s.EastWestDistance += s.Waypoint.EastWest * multiplier
	s.NorthSouthDistance += s.Waypoint.NorthSouth * multiplier
}

// GetManhattanDistance returns the total Manhattan Distance the ship has travelled
func (s *Ship) GetManhattanDistance() float64 {
	sum := float64(0)

	if s.NorthSouthDistance < 0 {
		sum += s.NorthSouthDistance * -1
	} else {
		sum += s.NorthSouthDistance
	}

	if s.EastWestDistance < 0 {
		sum += s.EastWestDistance * -1
	} else {
		sum += s.EastWestDistance
	}

	return sum
}
