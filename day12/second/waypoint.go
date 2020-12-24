package main

import "math"

// Waypoint indicates a waypoint to be used as orientation for a ship
type Waypoint struct {
	NorthSouth float64
	EastWest   float64
}

// NewWaypoint provides a pointer to a new waypoint to be used
func NewWaypoint() *Waypoint {
	return &Waypoint{
		NorthSouth: 1,
		EastWest:   10,
	}
}

// Move will move the waypoint in the given direction, by the given units
func (w *Waypoint) Move(dir Direction, units float64) {
	if units < 0 {
		// we need absolute numbers
		units *= -1
	}

	if dir == North {
		w.NorthSouth += units
	} else if dir == South {
		w.NorthSouth -= units
	} else if dir == East {
		w.EastWest += units
	} else if dir == West {
		w.EastWest -= units
	}
}

// Rotate will rotate the waypoint on the 0,0 axis by the given degrees
func (w *Waypoint) Rotate(deg float64) {
	rad := (deg / 360) * math.Pi * 2

	nEW := (w.EastWest * math.Cos(rad)) - (w.NorthSouth * math.Sin(rad))
	nNS := (w.EastWest * math.Sin(rad)) + (w.NorthSouth * math.Cos(rad))

	w.EastWest = nEW
	w.NorthSouth = nNS
}
