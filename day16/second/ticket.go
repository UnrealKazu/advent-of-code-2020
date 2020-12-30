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
	Name        string
	FirstLower  int
	FirstUpper  int
	SecondLower int
	SecondUpper int
}

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	intervals := make([]Interval, 0)
	valids := make([][]int, 0)

	var ownTicket []int

	fieldSec := true
	ownTicketSec := false
	nearbyTicketsSec := false

	for s.Scan() {
		line := s.Text()

		if fieldSec {
			if line == "" {
				fieldSec = false
				ownTicketSec = true
				continue
			}

			// parse all intervals and store them with their proper name in the intervals slice
			fieldSplit := strings.Split(line, ":")
			intName := fieldSplit[0]
			ranges := strings.Split(fieldSplit[1], " or ")

			var firstLower int
			var firstUpper int
			var secondLower int
			var secondUpper int

			for i, r := range ranges {
				ra := strings.Split(r, "-")

				low, _ := strconv.Atoi(strings.Trim(ra[0], " "))
				upp, _ := strconv.Atoi(strings.Trim(ra[1], " "))

				if i == 0 {
					firstLower = low
					firstUpper = upp
				} else {
					secondLower = low
					secondUpper = upp
				}
			}

			intervals = append(intervals, Interval{
				Name:        intName,
				FirstLower:  firstLower,
				FirstUpper:  firstUpper,
				SecondLower: secondLower,
				SecondUpper: secondUpper,
			})
		}

		if ownTicketSec {
			if line == "your ticket:" {
				continue
			} else if line == "" {
				ownTicketSec = false
				nearbyTicketsSec = true
				continue
			}

			// store our own ticket for now
			ownSplit := strings.Split(line, ",")

			ownTicket = make([]int, len(ownSplit))

			for i, v := range ownSplit {
				ownTicket[i], _ = strconv.Atoi(v)
			}
		}

		if nearbyTicketsSec {
			if line == "nearby tickets:" {
				// skip the first line
				continue
			}

			ints := strings.Split(line, ",")
			valid := true
			vals := make([]int, 0)

			// this little check weeds out all invalid tickets
			for _, i := range ints {
				val, _ := strconv.Atoi(i)

				if !validValue(intervals, val) {
					valid = false
					break
				} else {
					vals = append(vals, val)
				}
			}

			if valid {
				// this way, we have an array of valid tickets
				valids = append(valids, vals)
			}
		}
	}

	// for each ticket field, get all possible fields (i.e. intervals)
	possibles := getPossibleIntervals(intervals, valids)

	// now, reduce all those possible fields to one field
	sorted := determineSequence(possibles)

	mult := 1

	for i, v := range sorted {
		if strings.Contains(v.Name, "departure") {
			mult *= ownTicket[i]
		}
	}

	fmt.Printf("Multiplicity value is %d", mult)
}

func validValue(intervals []Interval, val int) bool {
	for _, i := range intervals {
		if (val >= i.FirstLower && val <= i.FirstUpper) || (val >= i.SecondLower && val <= i.SecondUpper) {
			return true
		}
	}

	return false
}

func getValidIntervals(intervals []Interval, val int) []Interval {
	valids := make([]Interval, 0)

	for _, i := range intervals {
		if (val >= i.FirstLower && val <= i.FirstUpper) || (val >= i.SecondLower && val <= i.SecondUpper) {
			valids = append(valids, i)
		}
	}

	return valids
}

func getPossibleIntervals(intervals []Interval, tickets [][]int) [][]Interval {
	possibles := make([][]Interval, len(tickets[0]))

	for i := 0; i < len(tickets[0]); i++ {
		// for all possible fields

		possInts := make([]Interval, len(intervals))
		copy(possInts, intervals)

		for j := 0; j < len(tickets); j++ {
			// iterate over all tickets

			val := tickets[j][i]

			// filter out all invalid intervals
			possInts = getValidIntervals(possInts, val)
		}

		possibles[i] = possInts
	}

	return possibles
}

func determineSequence(possibles [][]Interval) []Interval {
	visited := make(map[int]bool)

	haveRemoved := true
	for haveRemoved {
		haveRemoved = false

		// loop over all ticket fields
		for i := 0; i < len(possibles); i++ {
			if val, ok := visited[i]; ok && val {
				// we've already checked this field
				continue
			}

			if len(possibles[i]) == 1 {
				// if a field has only one remaining possible field, then we can cross out that field
				// for all other fields, which is what we're doing here
				it := possibles[i][0]

				for j, p := range possibles {
					if val, ok := visited[j]; (ok && val) || i == j {
						// we've already checked this field
						continue
					}

					for k, pp := range p {
						if pp.Name == it.Name {
							// remove this interval from the possibilities
							possibles[j] = remove(p, k)
							haveRemoved = true
						}
					}
				}

				// we're done with this field, so add it to the visited map
				visited[i] = true
			}
		}
	}

	// when we're done we have a slice of 1-size slices, so we can reduce it
	result := make([]Interval, len(possibles))

	for i, p := range possibles {
		result[i] = p[0]
	}

	return result
}

func remove(s []Interval, i int) []Interval {
	// swap the to-be removed to the end
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	// and then return the end minus one
	return s[:len(s)-1]
}

func removeAll(s [][]Interval, i int) [][]Interval {
	// swap the to-be removed to the end
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	// and then return the end minus one
	return s[:len(s)-1]
}
