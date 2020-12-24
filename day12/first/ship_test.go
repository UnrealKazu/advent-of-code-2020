package main

import (
	"fmt"
	"testing"
)

func TestRotate_WithValidDegrees_ShouldResultInCorrectDirection(t *testing.T) {
	var tests = []struct {
		dir Orientation
		deg int
		exp Direction
	}{
		{Right, 90, South},
		{Right, 180, West},
		{Right, 270, North},
		{Right, 360, East},
		{Left, 90, North},
		{Left, 180, West},
		{Left, 270, South},
		{Left, 360, East},
	}

	for _, tt := range tests {
		s := New()

		testname := fmt.Sprintf("%d,%d", tt.dir, tt.deg)
		t.Run(testname, func(t *testing.T) {
			s.Rotate(tt.dir, tt.deg)
			if s.Direction != tt.exp {
				t.Errorf("expected %d, got %d", tt.exp, s.Direction)
			}
		})
	}
}

func TestRotate_WithValidDegreesAndRotatedShip_ShouldResultInCorrectDirection(t *testing.T) {
	var tests = []struct {
		dir Orientation
		deg int
		exp Direction
	}{
		{Right, 90, West},
		{Right, 180, North},
		{Right, 270, East},
		{Right, 360, South},
		{Left, 90, East},
		{Left, 180, North},
		{Left, 270, West},
		{Left, 360, South},
	}

	for _, tt := range tests {
		s := New()

		// pre-rotate the ship once to the right
		s.Rotate(Right, 90)

		testname := fmt.Sprintf("%d,%d", tt.dir, tt.deg)
		t.Run(testname, func(t *testing.T) {
			s.Rotate(tt.dir, tt.deg)
			if s.Direction != tt.exp {
				t.Errorf("expected %d, got %d", tt.exp, s.Direction)
			}
		})
	}
}

func TestMoveInDirection_WithNorthDirection_ShouldIncreaseNorthSouthDistanceSumCorrectly(t *testing.T) {
	s := New()

	s.MoveInDirection(North, 10)

	exp := 10

	if s.NorthSouthDistance != exp {
		t.Errorf("Unexpected distance measure. Expected %d, got %d", exp, s.NorthSouthDistance)
	}
}

func TestMoveInDirection_WithSouthDirection_ShouldIncreaseNorthSouthDistanceSumCorrectly(t *testing.T) {
	s := New()

	s.MoveInDirection(South, 10)

	exp := 10

	if s.NorthSouthDistance != exp {
		t.Errorf("Unexpected distance measure. Expected %d, got %d", exp, s.NorthSouthDistance)
	}
}

func TestMoveInDirection_WithEastDirection_ShouldIncreaseEastWestDistanceSumCorrectly(t *testing.T) {
	s := New()

	s.MoveInDirection(East, 10)

	exp := 10

	if s.EastWestDistance != exp {
		t.Errorf("Unexpected distance measure. Expected %d, got %d", exp, s.EastWestDistance)
	}
}

func TestMoveInDirection_WithWestDirection_ShouldIncreaseEastWestDistanceSumCorrectly(t *testing.T) {
	s := New()

	s.MoveInDirection(West, 10)

	exp := 10

	if s.EastWestDistance != exp {
		t.Errorf("Unexpected distance measure. Expected %d, got %d", exp, s.EastWestDistance)
	}
}
