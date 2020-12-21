package main

import "fmt"

// Seating struct provides a struct for iterating over the seating plan
type Seating struct {
	IsTouched  bool
	Plan       [][]string
	ShadowPlan [][]string
}

// New creates a new seating plan struct
func New(plan [][]string) *Seating {
	// initialize the shadow array with the same dimensions of the plan
	shad := make([][]string, 0)

	for i := 0; i < len(plan); i++ {
		shad = append(shad, make([]string, len(plan[i])))

		for j := 0; j < len(plan[i]); j++ {
			shad[i][j] = plan[i][j]
		}
	}

	return &Seating{
		IsTouched:  false,
		Plan:       plan,
		ShadowPlan: shad,
	}
}

// DetermineSeatPlan iterates over the given seating plan until it is stabilized
func (s *Seating) DetermineSeatPlan() [][]string {
	s.IsTouched = true

	nrofRounds := 0
	// as long as seats are changing, we continue iterating
	for s.IsTouched == true {
		// reset the touched variable
		s.IsTouched = false

		// copy the current seating plan over to the shadow plan
		// so that we have the original to take into account
		for a := 0; a < len(s.Plan); a++ {
			for b := 0; b < len(s.Plan[a]); b++ {
				s.ShadowPlan[a][b] = s.Plan[a][b]
			}
		}

		for i := 0; i < len(s.Plan); i++ {
			for j := 0; j < len(s.Plan[i]); j++ {
				new := ApplyRules(i, j, s.ShadowPlan)

				if new != s.Plan[i][j] {
					s.IsTouched = true
					s.Plan[i][j] = new
				}
			}
		}

		//fmt.Printf("Revised seating for round %d\n", nrofRounds)
		//s.PrintSeating(s.Plan)

		nrofRounds++
	}

	return s.Plan
}

// PrintSeating prints the current seating plan to stdout
func (s *Seating) PrintSeating(plan [][]string) {
	for _, pl1 := range plan {
		for _, pl2 := range pl1 {
			fmt.Printf("%v", pl2)
		}

		fmt.Println()
	}
}

// ApplyRules applies the set of rules to a seat, and returns the new value of the seat
func ApplyRules(i int, j int, shad [][]string) string {
	if shad[i][j] == "." {
		// this is a floor square, which do not change
		return "."
	}

	if shad[i][j] == "L" {
		anyOcc := false

		// empty seat, check if it has no adjacent occupied seats
		for k := -1; k <= 1; k++ {
			ki := i + k
			if ki < 0 || ki > len(shad)-1 {
				// these squares are outside of the matrix, so skip them
				continue
			}

			for l := -1; l <= 1; l++ {
				lj := j + l
				if lj < 0 || lj > len(shad[i])-1 {
					// same here, stop if we're out of bounds
					continue
				}

				if k == 0 && l == 0 {
					// this is identical to the center square, which we do not want to check
					continue
				}

				if shad[ki][lj] == "#" {
					// this neighbour is occupied, so log this
					anyOcc = true
				}
			}
		}

		if !anyOcc {
			// if all neighbours are not occupied, this seat becomes occupied
			return "#"
		}
	}

	if shad[i][j] == "#" {
		nrofOcc := 0

		// occupied seat, check if it has more than 4 adjacent occupied seats
		for k := -1; k <= 1; k++ {
			ki := i + k
			if ki < 0 || ki > len(shad)-1 {
				// these squares are outside of the matrix, so skip them
				continue
			}

			for l := -1; l <= 1; l++ {
				lj := j + l
				if lj < 0 || lj > len(shad[i])-1 {
					// same here, stop if we're out of bounds
					continue
				}

				if k == 0 && l == 0 {
					// this is identical to the center square, which we do not want to check
					continue
				}

				if shad[ki][lj] == "#" {
					// this neighbour is occupied, so log this
					nrofOcc++
				}
			}
		}

		if nrofOcc >= 4 {
			// if 4 or more neighbours are occupied, this seat becomes empty
			return "L"
		}
	}

	return shad[i][j]
}
