package main

import (
	"regexp"
	"strings"
)

// Rulebook contains all rules about what bag can contain what
type Rulebook struct {
	Rules map[string][]string
}

var ruleExp *regexp.Regexp
var bagExp *regexp.Regexp

// New provides a new rulebook, and initializes regex variables
func New() *Rulebook {
	ruleExp, _ = regexp.Compile(`^([a-z\s]+) bags contain (.*)`)
	bagExp, _ = regexp.Compile(`^(?P<nrofBags>[0-9|no other]+)(?P<bagName>[a-z\s]*) bag[s.]+$`)

	r := Rulebook{
		Rules: make(map[string][]string),
	}

	return &r
}

func (r *Rulebook) addRule(rule string) {
	match := ruleExp.FindStringSubmatch(rule)

	bagName := match[1]
	bagContents := match[2]

	contents := strings.Split(bagContents, ",")

	for _, bag := range contents {
		bagMatch := bagExp.FindStringSubmatch(bag + ".") // the . is a dirty hack to make the regex more precise. Don't ask

		iBagName := bagMatch[2]

		// if the bagname is empty, we have no contents for the bag
		if iBagName != "" {
			if _, ok := r.Rules[iBagName]; ok {
				// bag is already known, check if the container bag is also known
				known := false

				for _, k := range r.Rules[iBagName] {
					if k == bagName {
						known = true
						break
					}
				}

				if !known {
					r.Rules[iBagName] = append(r.Rules[iBagName], bagName)
				}
			} else {
				newRules := make([]string, 0)
				newRules = append(newRules, bagName)
				r.Rules[iBagName] = newRules
			}
		}
	}
}
