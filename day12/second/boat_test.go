package main

import "testing"

func TestBoat_WithExample_ShouldHaveCorrectDistances(t *testing.T) {
	s := New()

	dirs := []string{"F10", "N3", "F7", "R90", "F11"}

	for _, d := range dirs {
		ProcessInstruction(s, d)
	}

	exp := 286
	act := int(s.GetManhattanDistance())

	if act != exp {
		t.Errorf("Unexpected Manhattan Distance. Expected %d, got %d", exp, act)
	}
}
