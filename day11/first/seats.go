package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	m := make([][]string, 0)

	i := 0
	for s.Scan() {
		line := s.Text()

		m = append(m, make([]string, len(line)))

		for j, s := range line {
			m[i][j] = string(s)
		}
		i++
	}

	ss := New(m)

	pl := ss.DetermineSeatPlan()

	// determine the number of occupied seats
	occ := 0
	for _, pl1 := range pl {
		for _, pl2 := range pl1 {
			if pl2 == "#" {
				occ++
			}
		}
	}

	fmt.Printf("The number of occupied seats is: %d", occ)
}
