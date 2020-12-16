package main

import "testing"

func TestReadInput(t *testing.T) {
	slope := readInput()

	if len(slope) != 323 {
		t.Errorf("Unexpected number of slope lines. Expected %d, found %d", 323, len(slope))
	}

	if len(slope[0]) != 31 {
		t.Errorf("Unexpected number of slope columns. Expected %d, found %d", 31, len(slope[0]))
	}
}
