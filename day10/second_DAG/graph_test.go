package main

import "testing"

func TestNew_WithSetOfAdapters_ShouldGenerateCorrectGraph(t *testing.T) {
	values := []int{0, 1, 2, 3}

	g := New(values)

	if g.Root == nil || g.Sink == nil {
		t.Errorf("Incorrect graph generated")
	}
}
