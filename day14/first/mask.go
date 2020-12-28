package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	mem := make(map[int]int)
	var mask map[int]bool

	r := regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)

	for s.Scan() {
		line := s.Text()

		if strings.Contains(line, "mask") {
			// (re)initialize the mask
			mask = convertMask(line)
		} else {
			res := r.FindStringSubmatch(line)
			pos, _ := strconv.Atoi(res[1])
			val, _ := strconv.Atoi(res[2])

			mem[pos] = applyMask(val, mask)
		}
	}

	sum := 0
	for _, e := range mem {
		sum += e
	}

	fmt.Printf("Sum of all values in memory is %d", sum)
}

func applyMask(val int, mask map[int]bool) int {
	for i, m := range mask {
		if m == true {
			// shift to 1
			val |= 1 << i
		} else {
			// shift to 0
			val &= ^(1 << i)
		}
	}

	return val
}

func convertMask(rawMask string) map[int]bool {
	mask := make(map[int]bool)

	for i, m := range rawMask {
		switch string(m) {
		case "X":
			continue
		case "1":
			mask[len(rawMask)-1-i] = true
		case "0":
			mask[len(rawMask)-1-i] = false
		}
	}

	return mask
}
