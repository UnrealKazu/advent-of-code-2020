// The solution for this puzzle (day 7, first) is a complete mess, I must admit...
// It is a brute-force solution with a bottom-up solution. At first I thought this would be the simplest one,
// as thought that implementing this as a tree structure would be too much overkill.
// However, in the end I'm convinced that using a tree would've been way easier and more straight forward...
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	r := New()

	for s.Scan() {
		line := s.Text()

		r.addRule(line)
	}

	num := getShinyContainers(r)

	fmt.Printf("Number of possible bag containers for shiny gold bag: %d", num)
}

func getShinyContainers(r *Rulebook) int {
	parents := r.Rules["shiny gold"]
	// we need a blacklist to keep track of bags we've already visited
	blacklist := r.Rules["shiny gold"]
	final := 0

	for len(parents) != 0 {
		// each bag that can contain the previous one counts as a bag that can hold the shiny gold one _eventually_
		final += len(parents)

		// we need a new array to prevent modifying the one we're looping over
		newParents := make([]string, 0)

		for _, parent := range parents {
			if val, ok := r.Rules[parent]; ok {
				// parent(s) found, add them all to the possible parents array
				for _, c := range val {
					if !checkBlacklist(blacklist, c) {
						// we haven't done this bagtype yet, so we can add it
						newParents = append(newParents, c)
						blacklist = append(blacklist, c)
					}
				}
			}
		}

		parents = newParents
	}

	return final
}

func checkBlacklist(list []string, val string) bool {
	for _, v := range list {
		if val == v {
			return true
		}
	}

	return false
}
