package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Bag represents a single bag with an indication of how much have been contained
type Bag struct {
	Name       string
	NrofInside int64
}

// Rulebook contains all rules about what bag can contain what
type Rulebook struct {
	Rules map[string][]Bag
}

var ruleExp *regexp.Regexp
var bagExp *regexp.Regexp

// New provides a new rulebook, and initializes regex variables
func New() *Rulebook {
	ruleExp, _ = regexp.Compile(`^([a-z\s]+) bags contain (.*)`)
	bagExp, _ = regexp.Compile(`^[ ]*(?P<nrofBags>[0-9|no other]+)(?P<bagName>[a-z\s]*) bag[s.]+$`)

	r := Rulebook{
		Rules: make(map[string][]Bag),
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

		var iBagName string
		nrofInner := int64(0)

		if len(bagMatch) == 2 {
			// we have a bag with no contents
			iBagName = bagMatch[1]
		} else {
			i, _ := strconv.Atoi(strings.Trim(bagMatch[1], " "))
			nrofInner = int64(i)
			iBagName = bagMatch[2]
		}

		innerBag := Bag{
			Name:       iBagName,
			NrofInside: nrofInner,
		}

		if _, ok := r.Rules[bagName]; ok {
			// bag is already known, check if the child bag is also known
			known := false

			for _, k := range r.Rules[bagName] {
				if k.Name == innerBag.Name {
					known = true
					break
				}
			}

			if !known {
				// child bag is not known, so add it
				r.Rules[bagName] = append(r.Rules[bagName], innerBag)
			}
		} else {
			// the entire parent/child bags are not known, so add them
			bags := make([]Bag, 0)
			bags = append(bags, innerBag)
			r.Rules[bagName] = bags
		}
	}
}
