package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")

	s := bufio.NewScanner(f)

	nrofValid := 0

	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, " ")

		char := split[1][0]
		pass := split[2]

		policy := strings.Split(split[0], "-")
		lower, _ := strconv.Atoi(policy[0])
		upper, _ := strconv.Atoi(policy[1])

		nrof := strings.Count(pass, string(char))

		if nrof >= lower && nrof <= upper {
			nrofValid++
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", nrofValid)
}
