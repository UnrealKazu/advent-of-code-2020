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

	d := New(25, w)

	weak := d.Attack()

	fmt.Printf("Attack entry point is: %d", weak)
}
