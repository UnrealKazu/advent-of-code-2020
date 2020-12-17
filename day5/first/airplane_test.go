package main

import (
	"fmt"
	"testing"
)

func TestProcessLine_ShouldReturnCorrectSeatID(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.in)
		t.Run(testname, func(t *testing.T) {
			act := processLine(tt.in)
			if act != tt.want {
				t.Errorf("Expected %d, actual %d", tt.want, act)
			}
		})
	}
}
