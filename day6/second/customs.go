package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	newGroup := true
	ans := ""

	sum := 0

	for s.Scan() {
		line := s.Text()

		if line != "" {
			// same group, check for differences
			if newGroup {
				// first person of a group, so everything is distinct
				ans = line
				newGroup = false
			} else {
				ans = intersectAnswers(ans, line)
			}
		} else {
			// end of group, sort the answers
			sum += len(ans)

			// reset the answer and continue with a new passport
			ans = ""
			newGroup = true
		}
	}

	fmt.Printf("Total sum of matching 'yes' answers: %d", sum)
}

func intersectAnswers(s1, s2 string) string {
	isect := ""

	for _, ss2 := range s2 {
		if strings.ContainsRune(s1, ss2) {
			isect += string(ss2)
		}
	}

	return isect
}
