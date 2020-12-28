package main

import (
	"testing"
)

func TestApplyMask_WithMask_ShouldReturnCorrectAddress(t *testing.T) {
	mask := convertMask("000000000000000000000000000000X1001X")

	// address number 42 should return 26, 27, 58, and 59
	addrs := applyMask(42, mask, len(mask)-1)
	exp := 4

	if len(addrs) != exp {
		t.Errorf("Unexpected number of addresses returned. Expected %d, got %d", exp, len(addrs))
	}
}

func TestApplyMask_WithBiggerMask_ShouldReturnCorrectAddress(t *testing.T) {
	mask := convertMask("00000000000000000000000000000000X0XX")

	// address number 26 should return 16, 17, 18, 19, 24, 25, 26, 27
	addrs := applyMask(26, mask, len(mask)-1)
	exp := 8

	if len(addrs) != exp {
		t.Errorf("Unexpected number of addresses returned. Expected %d, got %d", exp, len(addrs))
	}
}
