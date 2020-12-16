package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
			cat += line + " "
		} else {
			// end of passport, check for fields
			m := rexp.FindAllString(cat, -1)

			if len(m) == 7 {
				// all required fields are present. Now check the validity of the values
				if checkValueValidity(strings.Trim(cat, " ")) {
					valid++
				}
			}

			// reset the cat and continue with a new passport
			cat = ""
		}
	}

	fmt.Printf("Number of valid passports: %d", valid)
}

func checkValueValidity(line string) bool {
	lspl := strings.Split(line, " ")

	for _, pair := range lspl {
		keyspl := strings.Split(pair, ":")

		if len(keyspl) != 2 {
			// we have an invalid key-value combo, so by definition this is an invalid passport
			return false
		}

		t := keyspl[0]
		val := keyspl[1]

		var valid bool

		switch t {
		case "byr":
			valid = checkNumber(val, 1920, 2002)
			break
		case "iyr":
			valid = checkNumber(val, 2010, 2020)
			break
		case "eyr":
			valid = checkNumber(val, 2020, 2030)
			break
		case "hgt":
			valid = checkHeight(val)
			break
		case "hcl":
			valid = checkHairColor(val)
			break
		case "ecl":
			valid = checkEyeColor(val)
			break
		case "pid":
			valid = checkPid(val)
			break
		case "cid":
			// these are valid by default, because they are optional
			valid = true
			break
		}

		if !valid {
			return false
		}
	}

	return true
}

func checkNumber(value string, lower int, upper int) bool {
	num, err := strconv.Atoi(value)

	if err != nil {
		return false
	}

	if num >= lower && num <= upper {
		return true
	}

	return false
}

func checkHeight(value string) bool {
	t := value[len(value)-2:]
	n := value[:len(value)-2]

	switch t {
	case "cm":
		return checkNumber(n, 150, 193)
	case "in":
		return checkNumber(n, 59, 76)
	default:
		return false
	}
}

func checkHairColor(value string) bool {
	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", value)

	return match
}

func checkEyeColor(value string) bool {
	match, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", value)

	return match
}

func checkPid(value string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", value)

	return match
}
