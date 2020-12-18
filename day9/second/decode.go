package main

import (
	"sort"
)

// Decoder struct provides a struct with a window for decoding
type Decoder struct {
	Window []int
	Feed   []int
	Attack int
}

// New creates a new Decoder object and initializes the decoder with a preamble with a given size
func New(attack int, feed []int) *Decoder {
	d := Decoder{
		Window: make([]int, 0),
		Feed:   feed,
		Attack: attack,
	}

	return &d
}

// FindWeakness finds the set of contiguous numbers which add up to the attack int. When this window is found,
// it returns the sum of the smallest and largest number
func (d *Decoder) FindWeakness() int {
	sum := 0
	for _, e := range d.Feed {
		// add the new element to it
		d.Window = append(d.Window, e)
		sum += e

		// shrink the window if necessary
		sum = shrinkWindow(d, sum)

		if sum == d.Attack {
			// weakness found, get the smallest and largest number
			sort.Ints(d.Window)
			return d.Window[0] + d.Window[len(d.Window)-1]
		}
	}

	return -1
}

// shrinkWindow checks if the sum of the window is bigger than the attack
// if so, keep on shrinking the window by removing elements from the front
func shrinkWindow(d *Decoder, sum int) int {
	// keep shrinking until the window is below the sum or the window is empty
	for sum > d.Attack && len(d.Window) > 0 {
		sum -= d.Window[0]
		d.Window = d.Window[1:]
	}

	return sum
}
