package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)

	inst := make([]Instruction, 0)

	for s.Scan() {
		line := s.Text()

		inst = append(inst, processLine(line))
	}

	accum := detectLoop(inst)

	fmt.Printf("Loop detected. Accumulator value: %d", accum)
}

func detectLoop(inst []Instruction) int {
	i := 0
	accum := 0

	for i < len(inst) {
		ex := &inst[i]

		if ex.Visited {
			// loop detected, return the accumulator value
			return accum
		}

		switch ex.Type {
		case nop:
			// nop results in just a continuation
			i++
			break
		case acc:
			// modify the accumulator and continue with the next instruction
			accum += ex.Value
			i++
			break
		case jmp:
			// jump to the given row, don't touch the accumulator
			i += ex.Value
			break
		}

		ex.Visited = true
	}

	return -1
}

func processLine(line string) Instruction {
	expl := strings.Split(line, " ")

	var it Type
	val, _ := strconv.Atoi(expl[1])

	switch expl[0] {
	case "nop":
		it = nop
		break
	case "acc":
		it = acc
		break
	case "jmp":
		it = jmp
	}

	return Instruction{
		Type:    it,
		Value:   val,
		Visited: false,
	}
}
