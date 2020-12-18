package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	w := make([]int, 0)
	for s.Scan() {
		line := s.Text()

		i, _ := strconv.Atoi(line)
		w = append(w, i)
	}

	// 85848519 is the result of the attack from the first puzzle
	d := New(85848519, w)

	weak := d.FindWeakness()

	fmt.Printf("Weakness is: %d", weak)
}
