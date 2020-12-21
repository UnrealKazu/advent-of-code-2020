package main

import (
	"testing"
)

func TestDetermineSeatPlan_WithExample_ShouldReturnCorrectNrofOccupiedSiets(t *testing.T) {
	m := [][]string{
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
		{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
		{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
		{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
		{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"},
		{"L", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
		{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
	}

	s := New(m)

	pl := s.DetermineSeatPlan()

	act := 0

	for _, pl1 := range pl {
		for _, pl2 := range pl1 {
			if pl2 == "#" {
				act++
			}
		}
	}

	exp := 26

	if act != exp {
		t.Errorf("Unexpected number of occupied seats. Expected %d, got %d", exp, act)
	}
}

func TestGetSeatCheckPositions_WithCenterSeatAndFirstRound_ShouldReturnCorrectPositions(t *testing.T) {
	pos := GetSeatCheckPositions(1, 6, 6, 3, 3)

	exp := []Position{
		{2, 2}, {2, 3}, {2, 4},
		{3, 2}, {3, 4},
		{4, 2}, {4, 3}, {4, 4},
	}

	valid := true
	for i, p := range exp {
		if pos[i].I != p.I || pos[i].J != p.J {
			valid = false
		}
	}

	if !valid {
		t.Error("Unexpected array of positions returned")
	}
}

func TestGetSeatCheckPositions_WithCenterSeatAndSecondRound_ShouldReturnCorrectPositions(t *testing.T) {
	pos := GetSeatCheckPositions(2, 6, 6, 3, 3)

	exp := []Position{
		{1, 1}, {1, 3}, {1, 5},
		{3, 1}, {3, 5},
		{5, 1}, {5, 3}, {5, 5},
	}

	valid := true
	for i, p := range exp {
		if pos[i].I != p.I || pos[i].J != p.J {
			valid = false
		}
	}

	if !valid {
		t.Error("Unexpected array of positions returned")
	}
}

func TestGetSeatCheckPositions_WithCornerSeatAndSixthRound_ShouldReturnCorrectPositions(t *testing.T) {
	pos := GetSeatCheckPositions(6, 6, 6, 6, 0)

	exp := []Position{
		{0, 0},
		{0, 6},
		{6, 6},
	}

	valid := true
	for i, p := range exp {
		if pos[i].I != p.I || pos[i].J != p.J {
			valid = false
		}
	}

	if !valid {
		t.Error("Unexpected array of positions returned")
	}
}

func TestGetSeatCheckPositions_WithOutOfBoundsRound_ShouldReturnNoPositions(t *testing.T) {
	pos := GetSeatCheckPositions(200, 6, 6, 6, 0)

	if len(pos) != 0 {
		t.Errorf("Unexpected number of positions received. Expected %d, got %d", 0, len(pos))
	}
}
