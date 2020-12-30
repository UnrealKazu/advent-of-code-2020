// This solution is an incredibly straightforward, non-intelligent brute-force solution.
// But hey, I knew that part two would be more difficult, and that this first part would not be
// worth the time to completely flesh out something like an interval tree or stuff like that.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interval defines a simple struct to keep track of the defined intervals
type Interval struct {
	Lower int
	Upper int
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	intervals := make([]Interval, 0)
	invalids := make([]int, 0)

	fieldSec := true
	ownTicket := false
	nearbyTickets := false

	for s.Scan() {
		line := s.Text()

		if fieldSec {
			if line == "" {
				fieldSec = false
				ownTicket = true
				continue
			}

			fieldSplit := strings.Split(line, ":")
			ranges := strings.Split(fieldSplit[1], " or ")

			for _, r := range ranges {
				ra := strings.Split(r, "-")

				low, _ := strconv.Atoi(strings.Trim(ra[0], " "))
				upp, _ := strconv.Atoi(strings.Trim(ra[1], " "))

				intervals = append(intervals, Interval{
					Lower: low,
					Upper: upp,
				})
			}
		}

		if ownTicket {
			// we don't need our own ticket, so pass over it
			if line == "" {
				ownTicket = false
				nearbyTickets = true
				continue
			}
		}

		if nearbyTickets {
			if line == "nearby tickets:" {
				// skip the first line
				continue
			}

			ints := strings.Split(line, ",")

			for _, i := range ints {
				val, _ := strconv.Atoi(i)

				if !validValue(intervals, val) {
					invalids = append(invalids, val)
				}
			}
		}
	}

	sum := 0
	for _, in := range invalids {
		sum += in
	}

	fmt.Printf("Ticket scanning error rate is %d", sum)
}

func validValue(intervals []Interval, val int) bool {
	for _, i := range intervals {
		if val >= i.Lower && val <= i.Upper {
			return true
		}
	}

	return false
}
