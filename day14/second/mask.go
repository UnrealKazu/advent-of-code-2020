package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Operation indicates the bitoperation that needs to be done on an address
type Operation int

const (
	// Noop indicates no operation needs to be done
	Noop Operation = iota
	// Overwrite indicates that the bit is overwritten with 1
	Overwrite
	// Floating indicates that both possibilities of the bit are valid (so both 0 and 1)
	Floating
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	mem := make(map[int]int)
	var mask map[int]Operation

	r := regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)

	for s.Scan() {
		line := s.Text()

		if strings.Contains(line, "mask") {
			// (re)initialize the mask
			rawMask := strings.Split(line, " = ")[1]
			mask = convertMask(rawMask)
		} else {
			res := r.FindStringSubmatch(line)
			addr, _ := strconv.Atoi(res[1])
			val, _ := strconv.Atoi(res[2])

			addrs := applyMask(addr, mask, len(mask)-1)

			for _, ad := range addrs {
				mem[ad] = val
			}
		}
	}

	sum := 0
	for _, e := range mem {
		sum += e
	}

	fmt.Printf("Sum of all values in memory is %d", sum)
}

func applyMask(addr int, mask map[int]Operation, pos int) []int {
	addrs := make([]int, 0)

	for i := pos; i >= 0; i-- {
		m := mask[i]

		if m == Overwrite {
			// shift to 1
			addr |= 1 << i
		} else if m == Floating {
			// go recursive in both possibilities, and break off the current loop
			addr1 := addr
			addr2 := addr

			addr1 |= 1 << i
			addr2 &= ^(1 << i)

			addrs = append(addrs, applyMask(addr1, mask, i-1)...)
			addrs = append(addrs, applyMask(addr2, mask, i-1)...)

			return addrs
		}
	}

	// if we've reached this point, it means we're at the end of one address, so return it
	addrs = append(addrs, addr)

	return addrs
}

func convertMask(rawMask string) map[int]Operation {
	mask := make(map[int]Operation)

	for i, m := range rawMask {
		var op Operation
		switch string(m) {
		case "X":
			op = Floating
		case "1":
			op = Overwrite
		case "0":
			op = Noop
		}

		mask[len(rawMask)-1-i] = op
	}

	return mask
}
