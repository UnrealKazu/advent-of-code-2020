package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	a := make([]int, 0)

	for s.Scan() {
		line := s.Text()

		i, _ := strconv.Atoi(line)
		a = append(a, i)
	}

	num := determineAdapterSeq(a)

	fmt.Printf("Jolt difference multiplication: %d", num)
}

func determineAdapterSeq(a []int) int {
	// the charging outlet has a rating of 0
	a = append(a, 0)

	sort.Ints(a)

	// the device has a rating of the highest + 3
	a = append(a, a[len(a)-1]+3)

	oneD := 0
	twoD := 0
	threeD := 0

	for i, e := range a {
		if i == len(a)-1 {
			// the final entry is the device, so we can skip it
			continue
		}

		diff := a[i+1] - e

		switch diff {
		case 1:
			oneD++
			break
		case 2:
			twoD++
			break
		case 3:
			threeD++
			break
		default:
			fmt.Printf("Unexpected diff detected: %d\n", diff)
		}
	}

	return oneD * threeD
}
