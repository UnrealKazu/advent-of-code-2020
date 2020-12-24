package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	ship := New()
	for s.Scan() {
		line := s.Text()

		ProcessInstruction(ship, line)
	}

	fmt.Printf("Manhattan distance travelled is %d", int(ship.GetManhattanDistance()))
}

// ProcessInstruction parses the raw instruction string, and feeds the command to the ship
func ProcessInstruction(s *Ship, dir string) {
	d := dir[:1]
	units, _ := strconv.ParseFloat(dir[1:], 64)

	switch d {
	case "N":
		s.Waypoint.Move(North, units)
		break
	case "E":
		s.Waypoint.Move(East, units)
		break
	case "S":
		s.Waypoint.Move(South, units)
		break
	case "W":
		s.Waypoint.Move(West, units)
	case "F":
		s.MoveForward(units)
	case "R":
		s.Waypoint.Rotate(units * -1)
	case "L":
		s.Waypoint.Rotate(units)
	}
}
