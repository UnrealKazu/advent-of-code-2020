// This solution uses a lineair solution for getting all possible (valid) adapter permutations.
// First we create a DAG (directed acyclic graph) using each adapter as a node, and edges for all possible other
// adapters. Next, we just have to recursively visit all nodes and their possible edges to count all possible paths.
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

	a := make([]int, 1)
	a[0] = 0

	for s.Scan() {
		line := s.Text()

		i, _ := strconv.Atoi(line)
		a = append(a, i)
	}

	sort.Ints(a)

	a = append(a, a[len(a)-1]+3)

	g := New(a)
	num := g.GetNrofAllPaths()

	fmt.Printf("Jolt difference multiplication: %d", num)
}
