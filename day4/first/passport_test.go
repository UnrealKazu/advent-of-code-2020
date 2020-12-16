package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	rexp, _ := regexp.Compile("byr|iyr|eyr|hgt|hcl|ecl|pid")

	fmt.Println(rexp.FindAllString("eyr:2029 pid:157374862byr:1991 ecl:amb hcl:#a97842 hgt:178cm", -1))
}
