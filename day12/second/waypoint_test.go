package main

import "testing"

func TestRotate_With90Degrees_ShouldProduceCorrectCoordinates(t *testing.T) {
	w := NewWaypoint()

	w.EastWest = 10
	w.NorthSouth = 4

	w.Rotate(-90)

	ewExp := 4
	nsExp := -10

	if ewExp != int(w.EastWest) || nsExp != int(w.NorthSouth) {
		t.Errorf("Unexpected coordinates. Expected (%d,%d), got (%v,%v)", ewExp, nsExp, w.EastWest, w.NorthSouth)
	}
}
