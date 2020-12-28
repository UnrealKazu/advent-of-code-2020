package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(59 - (939 % 59))

	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	l := make([]string, 2)
	i := 0
	for s.Scan() {
		line := s.Text()

		l[i] = line
		i++
	}

	wait, _ := strconv.Atoi(l[0])

	bussesStr := strings.Split(l[1], ",")

	// initialize at max int
	lowestWait := int(^uint(0) >> 1)
	var theBus int

	for _, b := range bussesStr {
		// skip busses that are out of commission
		if b == "x" {
			continue
		}

		bus, _ := strconv.Atoi(b)

		busWait := bus - (wait % bus)

		if busWait < lowestWait {
			lowestWait = busWait
			theBus = bus
		}
	}

	fmt.Printf("You should take bus: %d. Multiplication answer: %d", theBus, theBus*lowestWait)
}
