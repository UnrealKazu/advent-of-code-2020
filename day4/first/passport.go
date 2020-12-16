package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	valid := 0

	// we treat cid as optional, so it is not present in the regex
	rexp, _ := regexp.Compile("byr|iyr|eyr|hgt|hcl|ecl|pid")

	cat := ""

	for s.Scan() {
		line := s.Text()

		if line != "" {
			// same passport, concat
			cat += line
		} else {
			// end of passport, check for fields
			m := rexp.FindAllString(cat, -1)

			if len(m) == 7 {
				// full match
				valid++
			}

			// reset the cat and continue with a new passport
			cat = ""
		}
	}

	fmt.Printf("Number of valid passports: %d", valid)
}
