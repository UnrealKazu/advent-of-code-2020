package main

import "testing"

func TestDetermineAdapterSeq_WithSmallExample_ShouldReturnCorrectNrofSequences(t *testing.T) {
	a := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	sum := DetermineAdapterSeq(a)

	exp := int64(8)
	if sum != exp {
		t.Errorf("Unexpected nrof possible sequences. Expected %d, got %d", exp, sum)
	}
}

func TestDetermineAdapterSeq_WithBigExample_ShouldReturnCorrectNrofSequences(t *testing.T) {
	a := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

	sum := DetermineAdapterSeq(a)

	exp := int64(19208)
	if sum != exp {
		t.Errorf("Unexpected nrof possible sequences. Expected %d, got %d", exp, sum)
	}
}
