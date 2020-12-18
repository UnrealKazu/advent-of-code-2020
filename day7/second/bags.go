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

	num := getNumBags(r, "shiny gold")

	fmt.Printf("Number of individual bags inside shiny gold: %d", num)
}

func getNumBags(r *Rulebook, name string) int64 {
	children := r.Rules[name]

	sum := int64(0)

	for _, c := range children {
		sum += c.NrofInside + (c.NrofInside * getNumBags(r, c.Name))
	}

	return sum
}
