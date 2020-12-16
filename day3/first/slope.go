package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	slope := readInput()
	t := traverseSlope(slope)
	fmt.Printf("Number of trees encountered: %d", t)
}

func traverseSlope(slope map[int]map[int]string) int {
	i := 1
	j := 3

	t := 0

	for i < len(slope) {
		if slope[i][j%len(slope[i])] == "#" {
			t++
		}

		i++
		j += 3
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
