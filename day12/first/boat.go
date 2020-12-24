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

	fmt.Printf("Manhattan distance travelled is %d", ship.GetManhattanDistance())
}

// ProcessInstruction parses the raw instruction string, and feeds the command to the ship
func ProcessInstruction(s *Ship, dir string) {
	d := dir[:1]
	units, _ := strconv.Atoi(dir[1:])

	switch d {
	case "N":
		s.MoveInDirection(North, units)
		break
	case "E":
		s.MoveInDirection(East, units)
		break
	case "S":
		s.MoveInDirection(South, units)
		break
	case "W":
		s.MoveInDirection(West, units)
	case "F":
		s.MoveForward(units)
	case "R":
		s.Rotate(Right, units)
	case "L":
		s.Rotate(Left, units)
	}
}
