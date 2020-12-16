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

func TestRegexGroup(t *testing.T) {
	rexp, _ := regexp.Compile("byr:([0-9]{4})")

	fmt.Println(rexp.FindString("eyr:2029 pid:157374862 byr:1991 ecl:amb hcl:#a97842 hgt:178cm"))
}

func TestValidity(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		{"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f", true},
		{"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm", true},
		{"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022", true},
		{"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719", true},
		{"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926", false},
		{"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946", false},
		{"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277", false},
		{"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.in)
		t.Run(testname, func(t *testing.T) {
			act := checkValueValidity(tt.in)
			if act != tt.want {
				t.Errorf("Expected %t, actual %t", tt.want, act)
			}
		})
	}
}
