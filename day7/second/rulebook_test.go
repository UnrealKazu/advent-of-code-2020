package main

import (
	"testing"
)

func TestAddRule_WithSingleContentBag_ShouldAddRule(t *testing.T) {
	r := New()

	r.addRule("bright white bags contain 1 shiny gold bag.")

	if len(r.Rules) != 1 {
		t.Errorf("Unexpected number of rules. Expected %d, got %d", 1, len(r.Rules))
	}
}

func TestAddRule_WithNoContentBag_ShouldNotAddRule(t *testing.T) {
	r := New()

	r.addRule("faded blue bags contain no other bags.")

	if len(r.Rules) != 0 {
		t.Errorf("Unexpected number of rules. Expected %d, got %d", 0, len(r.Rules))
	}
}

func TestAddRule_WithMultipleContentsBag_ShouldAddMultipleRules(t *testing.T) {
	r := New()

	r.addRule("dark olive bags contain 3 faded blue bags, 4 dotted black bags.")

	if len(r.Rules) != 2 {
		t.Errorf("Unexpected number of rules. Expected %d, got %d", 2, len(r.Rules))
	}
}

func TestAddRule_WithMultipleParentBags_ShouldHaveMultipleParentRules(t *testing.T) {
	r := New()

	r.addRule("dark olive bags contain 3 faded blue bags.")
	r.addRule("dotted black bags contain 3 faded blue bags.")

	if len(r.Rules) != 1 {
		t.Errorf("Unexpected number of rules. Expected %d, got %d", 1, len(r.Rules))
	}

	if len(r.Rules["faded blue"]) != 2 {
		t.Errorf("Unexpected number of parent rules. Expected %d, got %d", 2, len(r.Rules["faded blue"]))
	}
}
