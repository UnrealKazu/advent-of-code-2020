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

	i := 0
	for s.Scan() {
		line := s.Text()

		inst = append(inst, processLine(line, i))
		i++
	}

	accum := fixLoop(inst)

	fmt.Printf("Loop fixed. Accumulator value: %d", accum)
}

func fixLoop(inst []Instruction) int {
	for i := range inst {
		ex := &inst[i]

		if ex.Type == acc {
			continue
		}

		if ex.Type == nop {
			// change it to an jmp and see if that works
			ex.Type = jmp
		} else if ex.Type == jmp {
			ex.Type = nop
		}

		// we need a deep copy to prevent modifying the original struct
		copy := copyInstructions(inst)

		if ok, val := detectLoop(copy); ok {
			// loop detected, so
			// revert the instruction type and continue
			if ex.Type == nop {
				ex.Type = jmp
			} else if ex.Type == jmp {
				ex.Type = nop
			}
		} else {
			return val
		}
	}

	return -1
}

func detectLoop(inst []Instruction) (bool, int) {
	i := 0
	accum := 0

	for i < len(inst) {
		ex := &inst[i]

		if ex.Visited {
			// loop detected, return the accumulator value
			return true, accum
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

	return false, accum
}

func processLine(line string, i int) Instruction {
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
		ID:      i,
		Type:    it,
		Value:   val,
		Visited: false,
	}
}

func copyInstructions(inst []Instruction) []Instruction {
	copy := make([]Instruction, 0)

	for _, v := range inst {
		copy = append(copy, Instruction{
			ID:      v.ID,
			Type:    v.Type,
			Value:   v.Value,
			Visited: v.Visited,
		})
	}

	return copy
}
