package main

import (
	"testing"
)

func TestConvertMask_WithSimpleStringMask_ShouldReturnCorrectMapMask(t *testing.T) {
	mask := convertMask("XXXXXX1XXX0")

	exp := make(map[int]bool)
	exp[0] = false
	exp[4] = true

	if mask[0] != exp[0] || mask[4] != exp[4] {
		t.Errorf("Unexpected mask returned. Expected %v, got %v", exp, mask)
	}
}

func TestApplyMask_WithSimplePositiveMask_ShouldReturnCorrectValue(t *testing.T) {
	mask := make(map[int]bool, 36)
	mask[1] = true

	// second bit true means that 1 should become 3
	act := applyMask(1, mask)
	exp := 3

	if act != exp {
		t.Errorf("Unexpected mask result. Expected %d, got %d", exp, act)
	}
}

func TestApplyMask_WithSimplePositiveMask_ShouldNotChangeValue(t *testing.T) {
	mask := make(map[int]bool, 36)
	mask[1] = true

	// second bit true means that 3 should stay 3
	act := applyMask(3, mask)
	exp := 3

	if act != exp {
		t.Errorf("Unexpected mask result. Expected %d, got %d", exp, act)
	}
}

func TestApplyMask_WithSimpleNegativeMask_ShouldReturnCorrectValue(t *testing.T) {
	mask := make(map[int]bool, 36)
	mask[1] = false

	// second bit false means that 3 should become 1
	act := applyMask(3, mask)
	exp := 1

	if act != exp {
		t.Errorf("Unexpected mask result. Expected %d, got %d", exp, act)
	}
}

func TestApplyMask_WithSimpleNegativeMask_ShouldNotChangeValue(t *testing.T) {
	mask := make(map[int]bool, 36)
	mask[1] = false

	// second bit false means that 1 should stay 1
	act := applyMask(1, mask)
	exp := 1

	if act != exp {
		t.Errorf("Unexpected mask result. Expected %d, got %d", exp, act)
	}
}

func TestApplyMask_WithComplexMask_ShouldReturnCorrectValue(t *testing.T) {
	mask := make(map[int]bool, 36)
	mask[1] = true
	mask[2] = false

	// mask is 01X, so 10X should become 01X
	// meaning that 5 should become 3
	act := applyMask(5, mask)
	exp := 3

	if act != exp {
		t.Errorf("Unexpected mask result. Expected %d, got %d", exp, act)
	}
}
