package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	slope := readInput()

	tt := traverseSlope(slope, 1, 1)
	tt *= traverseSlope(slope, 3, 1)
	tt *= traverseSlope(slope, 5, 1)
	tt *= traverseSlope(slope, 7, 1)
	tt *= traverseSlope(slope, 1, 2)

	fmt.Printf("Number of trees encountered: %d", tt)
}

func traverseSlope(slope map[int]map[int]string, rStep int, dStep int) int {
	i := dStep
	j := rStep

	t := 0

	for i < len(slope) {
		if slope[i][j%len(slope[i])] == "#" {
			t++
		}

		i += dStep
		j += rStep
	}

	return t
}

func readInput() map[int]map[int]string {
	slope := make(map[int]map[int]string)

	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	i := 0
	for s.Scan() {
		line := s.Text()

		// initialize map for this line
		slope[i] = make(map[int]string)

		for j, c := range line {
			slope[i][j] = string(c)
		}

		i++
	}

	return slope
}
