package main

import "fmt"

// Number is a number that has been spoken. Either for the first time or not, and with the round it has been spoken in
type Number struct {
	LastRound int
	PrevRound int
	First     bool
}

func main() {
	starting := []int{0, 14, 6, 20, 1, 4}

	final := play(starting)

	fmt.Printf("The 30000000th spoken number is %d", final)
}

func play(starting []int) int {
	// we play until the 2020th round

	spoken := make(map[int]*Number)

	for i, s := range starting {
		spoken[s] = &Number{
			LastRound: i,
			PrevRound: i,
			First:     true,
		}
	}

	last := starting[len(spoken)-1]

	// because 0 is a special case, we set the First property to false
	spoken[0].First = false

	for i := len(spoken); i < 30000000; i++ {
		checkNumber(spoken, last, i-1)

		num := spoken[last]

		if num.First {
			// this is second time the number has been spoken, so speak 0 and toggle this number
			last = 0
		} else {
			// this number has been spoken before, so generate the new number based on the diff in rounds and continue
			last = num.LastRound - num.PrevRound
		}
	}

	return last
}

// checkNumber checks if the given number exists in our history of spoken numbers
// if not, it initializes it and adds it
// if it is, it updates the rounds
func checkNumber(spoken map[int]*Number, number, round int) {
	if num, ok := spoken[number]; ok {
		num.First = false
		num.PrevRound = num.LastRound
		num.LastRound = round
		return
	}

	spoken[number] = &Number{
		LastRound: round,
		PrevRound: round,
		First:     true,
	}
}
