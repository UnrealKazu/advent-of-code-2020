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
	EastWestDistance   int
	NorthSouthDistance int
}

// New returns a pointer to a new Ship struct
func New() *Ship {
	return &Ship{
		Direction:          East,
		EastWestDistance:   0,
		NorthSouthDistance: 0,
	}
}

// Rotate will rotate the ship, given the orientation and degree of the turn
func (s *Ship) Rotate(or Orientation, deg int) {
	if deg%90 != 0 {
		panic("Unexpected rotation found")
	}

	var steps int

	if or == Right {
		steps = deg / 90
	} else {
		// turning left is just the inverse of turning right
		steps = (360 - deg) / 90
	}

	s.Direction = Direction(((int(s.Direction) + steps) % 4))
}

// MoveInDirection will move the ship in the given direction, by the given units
func (s *Ship) MoveInDirection(dir Direction, units int) {
	if units < 0 {
		// we need absolute numbers
		units *= -1
	}

	if dir == North {
		s.NorthSouthDistance += units
	} else if dir == South {
		s.NorthSouthDistance -= units
	} else if dir == East {
		s.EastWestDistance += units
	} else if dir == West {
		s.EastWestDistance -= units
	}
}

// MoveForward will move the ship in the direction it's already facing, by the given units
func (s *Ship) MoveForward(units int) {
	s.MoveInDirection(s.Direction, units)
}

// GetManhattanDistance returns the total Manhattan Distance the ship has travelled
func (s *Ship) GetManhattanDistance() int {
	sum := 0

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
