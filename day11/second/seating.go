package main

import "fmt"

// Direction shows the direction of the visibility
type Direction int

const (
	UpLeft Direction = iota
	Up
	UpRight
	Left
	Center
	Right
	DownLeft
	Down
	DownRight
)

// Position indicates a position in the matrix
type Position struct {
	I int // I know I and J are not standard, and should be X and Y, but this causes less headaches for this
	J int
}

// Seating struct provides a struct for iterating over the seating plan
type Seating struct {
	IsTouched  bool
	Plan       [][]string
	ShadowPlan [][]string
	MaxI       int
	MaxJ       int
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
		MaxI:       len(plan) - 1,
		MaxJ:       len(plan[0]) - 1,
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
				new := ApplyRules(i, j, s.ShadowPlan, s.MaxI, s.MaxJ)

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
func ApplyRules(i int, j int, shad [][]string, MaxI, MaxJ int) string {
	if shad[i][j] == "." {
		// this is a floor square, which do not change
		return "."
	}

	round := 1
	neighs := GetSeatCheckPositions(round, MaxI, MaxJ, i, j)
	dirBl := make([]Direction, 0)
	occ := 0

	for len(neighs) > 0 {
		for _, n := range neighs {
			if shad[n.I][n.J] == "#" {
				occ++
				// seat detected, blacklist the direction
				dirBl = append(dirBl, getDirection(n, Position{I: i, J: j}))
			} else if shad[n.I][n.J] == "L" {
				// seat detected, blacklist the direction
				dirBl = append(dirBl, getDirection(n, Position{I: i, J: j}))
			}
		}

		round++
		neighs = GetSeatCheckPositions(round, MaxI, MaxJ, i, j)
		neighs = FilterBlacklistedPositions(dirBl, i, j, neighs)
	}

	if shad[i][j] == "L" && occ == 0 {
		return "#"
	} else if shad[i][j] == "#" && occ >= 5 {
		return "L"
	}

	return shad[i][j]
}

// GetSeatCheckPositions returns a list of positions that can be checked given the current coordinates, and the round we're in
func GetSeatCheckPositions(round, iMax, jMax, curI, curJ int) []Position {
	pos := make([]Position, 0)

	// start at the top left, and start scanning per line
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				// this is the current position itself, so skip it
				continue
			}

			// take the current position, and apply the round factor to the relative position
			// this will generate a star effect on the positions
			ai := curI + (round * i)
			aj := curJ + (round * j)

			if (ai < 0 || ai > iMax) || (aj < 0 || aj > jMax) {
				// we're out of bounds of the matrix, so skip it
				continue
			}

			pos = append(pos, Position{I: ai, J: aj})
		}
	}

	return pos
}

// FilterBlacklistedPositions filters out all blacklisted position because the direction is already final
func FilterBlacklistedPositions(bl []Direction, i, j int, pos []Position) []Position {
	filtered := make([]Position, 0)

	for _, p := range pos {
		dir := getDirection(p, Position{I: i, J: j})

		if !containedInBlacklist(bl, dir) {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func getDirection(pos, center Position) Direction {
	if pos.I < center.I {
		// up!
		if pos.J < center.J {
			return UpLeft
		} else if pos.J == center.J {
			return Up
		} else {
			return UpRight
		}
	} else if pos.I == center.I {
		// in the middle!
		if pos.J < center.J {
			return Left
		} else if pos.J == center.J {
			return Center
		} else {
			return Right
		}
	} else {
		// down!
		if pos.J < center.J {
			return DownLeft
		} else if pos.J == center.J {
			return Down
		} else {
			return DownRight
		}
	}
}

func containedInBlacklist(bl []Direction, dir Direction) bool {
	for _, b := range bl {
		if b == dir {
			return true
		}
	}

	return false
}
