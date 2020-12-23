// This solution has an exponential running time because it has to recursivily traverse all possible paths
// Which is fine if the number of adapters are low, like in the given examples (see the unit tests)
// However, the actual puzzle has almost 100 adapters. The puzzle description already hinted at potentially
// trillions possible combinations. So having an exponential solution is impossible for this.
// See the DAG solution for the actual best solution which runs in lineair time.
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

	num := DetermineAdapterSeq(a)

	fmt.Printf("Jolt difference multiplication: %d", num)
}

// DetermineAdapterSeq determines the amount of possible permutations of the adapters
func DetermineAdapterSeq(a []int) int64 {
	// the charging outlet has a rating of 0
	a = append(a, 0)

	sort.Ints(a)

	// the device has a rating of the highest + 3
	a = append(a, a[len(a)-1]+3)

	return processSeq(a)
}

func processSeq(a []int) int64 {
	sum := int64(0)

	fmt.Printf("Doing array: %v\n", a)

	if len(a) == 1 {
		// we have made it to the end, so this is a valid sequence. Count it as such
		fmt.Println("Full sequence")
		return 1
	}

	e := a[0]

	for i := 1; i < len(a); i++ {
		n := a[i]
		diff := n - e

		if diff > 3 {
			// the array is sorted, so once we encounter a difference more than 3,
			// we will not get another element that is lower than 3
			break
		}

		sum += processSeq(a[i:])
	}

	//fmt.Printf("Done with array starting with %d. Returning sum: %d\n", e, sum)

	// if we've arrived here, then this direct sequence is not valid
	return sum
}
