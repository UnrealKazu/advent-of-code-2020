package main

// Decoder struct provides a struct with a window for decoding
type Decoder struct {
	Window       []int
	Feed         []int
	PreambleSize int
}

// New creates a new Decoder object and initializes the decoder with a preamble with a given size
func New(pSize int, feed []int) *Decoder {
	d := Decoder{
		Window:       make([]int, 0),
		Feed:         feed,
		PreambleSize: pSize,
	}

	for _, e := range feed[:pSize] {
		d.Window = append(d.Window, e)
	}

	return &d
}

// Attack seeks in the given feed for an integer that provides us with a possible attack entrypoint
func (d *Decoder) Attack() int {
	for _, e := range d.Feed[d.PreambleSize:] {
		if !checkSum(e, d.Window) {
			return e
		}

		// move the window
		d.Window = append(d.Window, e)
		d.Window = d.Window[1:]
	}

	return -1
}

func checkSum(i int, window []int) bool {
	for _, i1 := range window {
		for _, i2 := range window {
			if i1 != i2 && i1+i2 == i {
				return true
			}
		}
	}

	return false
}
