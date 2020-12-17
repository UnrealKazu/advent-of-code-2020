package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	cat := ""

	sum := 0

	for s.Scan() {
		line := s.Text()

		if line != "" {
			// same group, concat
			cat += line
		} else {
			// end of group, sort the answers
			sum += getNrofAnswers(cat)

			// reset the cat and continue with a new passport
			cat = ""
		}
	}

	fmt.Printf("Total sum of 'yes' answers: %d", sum)
}

func getNrofAnswers(answers string) int {
	i := make([]string, 0)

	// convert the string to a slice of strings, so that we can sort the characters
	for _, rr := range answers {
		i = append(i, string(rr))
	}

	// sort all answers
	sort.Strings(i)

	count := 0
	prev := ""
	for _, a := range i {
		if a != prev {
			count++
		}

		prev = a
	}

	return count
}
