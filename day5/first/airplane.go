package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	maxSeatID := 0

	for s.Scan() {
		line := s.Text()

		seatID := processLine(line)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Printf("Highest seat ID is %d", maxSeatID)
}

func processLine(line string) int {
	// first 7 characters are the rows
	row := binarySearch(line[:7], 0.0, 127.0)
	// last 3 characters are the columns
	column := binarySearch(line[7:], 0.0, 7.0)

	// return the seat ID
	return int(row*8 + column)
}

func binarySearch(spec string, lower float64, upper float64) float64 {
	if len(spec) == 0 {
		// search has finished, lower and upper are equal, return either
		return lower
	}

	char := string(spec[0])

	c := (upper - lower) / 2.0

	switch char {
	case "F", "L":
		m := math.Floor(c)
		return binarySearch(spec[1:], lower, lower+m)
	case "B", "R":
		m := math.Ceil(c)
		return binarySearch(spec[1:], lower+m, upper)
	}

	// theoretically, this should never be reached
	return -1
}
