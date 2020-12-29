package main

import "testing"

func TestPlay_WithMultipleGames_ReturnsCorrectValues(t *testing.T) {
	starting := []int{0, 3, 6}

	val := play(starting)
	exp := 436

	if val != 436 {
		t.Errorf("Unexpected result value from game. Expected %d, got %d", exp, val)
	}
}
