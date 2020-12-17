package main

// Type shows what the type of the instructionset is
type Type int

const (
	nop Type = iota
	acc
	jmp
)

// Instruction represents an operation that needs to be executed
type Instruction struct {
	Type    Type
	Value   int
	Visited bool
}
