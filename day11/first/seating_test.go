package main

import (
	"fmt"
	"testing"
)

func TestApplyRules_WithEmptySeats_ShouldReturnOccupiedSeat(t *testing.T) {
	m := [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}

	seat := ApplyRules(1, 1, m)
	exp := "#"

	if seat != exp {
		t.Errorf("Unexpected seat. Expected %s, got %s", exp, seat)
	}
}

func TestApplyRules_WithOneOccupiedSeat_SeatShouldStayEmpty(t *testing.T) {
	var tests = []struct {
		in   [][]string
		want string
	}{
		{[][]string{
			{"#", "L", "L"},
			{"L", "L", "L"},
			{"L", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "#", "L"},
			{"L", "L", "L"},
			{"L", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "#"},
			{"L", "L", "L"},
			{"L", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "L"},
			{"#", "L", "L"},
			{"L", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "L"},
			{"L", "L", "#"},
			{"L", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "L"},
			{"L", "L", "L"},
			{"#", "L", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "L"},
			{"L", "L", "L"},
			{"L", "#", "L"},
		}, "L"},
		{[][]string{
			{"L", "L", "L"},
			{"L", "L", "L"},
			{"L", "L", "#"},
		}, "L"},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			act := ApplyRules(1, 1, tt.in)
			if act != tt.want {
				t.Errorf("Expected %s, actual %s", tt.want, act)
			}
		})
	}
}

func TestApplyRules_WithOccupiedSeatAndFourOccupied_ShouldReturnEmptySeat(t *testing.T) {
	m := [][]string{
		{"#", "L", "L"},
		{"#", "#", "#"},
		{"L", "L", "#"},
	}

	seat := ApplyRules(1, 1, m)
	exp := "L"

	if seat != exp {
		t.Errorf("Unexpected seat. Expected %s, got %s", exp, seat)
	}
}

func TestApplyRules_WithOccupiedSeatAndThreeOccupied_SeatShouldStayOccupied(t *testing.T) {
	m := [][]string{
		{"#", "L", "L"},
		{"#", "#", "#"},
		{"L", "L", "L"},
	}

	seat := ApplyRules(1, 1, m)
	exp := "#"

	if seat != exp {
		t.Errorf("Unexpected seat. Expected %s, got %s", exp, seat)
	}
}

func TestApplyRules_WithEmptySeatAndEmptyAndFloorSurrounding_ShouldOccupySeat(t *testing.T) {
	m := [][]string{
		{"L", ".", "L"},
		{".", "L", "."},
		{"L", ".", "L"},
	}

	seat := ApplyRules(1, 1, m)
	exp := "#"

	if seat != exp {
		t.Errorf("Unexpected seat. Expected %s, got %s", exp, seat)
	}
}

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

	exp := 37

	if act != exp {
		t.Errorf("Unexpected number of occupied seats. Expected %d, got %d", exp, act)
	}
}
