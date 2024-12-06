package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type Vector struct {
	X int
	Y int
}

type Lab struct {
	Grid                    [][]byte
	GuardPosition           *Vector
	GuardVelocity           *Vector
	PositionVelocityHistory []string
}

func (l *Lab) PrintGrid() {
	for _, line := range l.Grid {
		fmt.Println(string(line))
	}
}

// Moves the guard in the direction they are currently facing.
// If the guard can't move forward, they rotate 90deg and try again.
// This will also replace the previous space with an X to track where the guard has been.
// First return value is whether or not the guard has left the map.
// Second return value is whether a loop has been detected.
func (l *Lab) MoveGuard() (bool, bool) {
	// Get a vector for the next guard position
	next := &Vector{
		X: l.GuardPosition.X + l.GuardVelocity.X,
		Y: l.GuardPosition.Y + l.GuardVelocity.Y,
	}

	// Update current position with X
	l.Grid[l.GuardPosition.Y][l.GuardPosition.X] = 'X'

	// If next position is already an 'X', then check if we have looped
	key := fmt.Sprintf("%v,%v:%v,%v", l.GuardPosition.X, l.GuardPosition.Y, l.GuardVelocity.X, l.GuardVelocity.Y)

	// Check if guard moved off the map
	if next.X < 0 || next.Y < 0 || next.Y >= len(l.Grid) || next.X >= len(l.Grid[next.Y]) {
		// If so, then we are done and guard has left
		return true, false
	}

	nextChar := l.Grid[next.Y][next.X]

	if nextChar == 'X' {
		// We'll know we've looped if entry in VelocityMap already exists
		for _, record := range l.PositionVelocityHistory {
			if record == key {
				return true, true
			}
		}
	}

	// Track next position + current velocity for checking for loop later
	l.PositionVelocityHistory = append(l.PositionVelocityHistory, key)

	// Check if next move is even viable
	if nextChar == '#' {
		// Rotate 90deg, and try to move again
		vel := l.GuardVelocity

		if vel.X == 0 && vel.Y == -1 {
			vel.X, vel.Y = 1, 0
		} else if vel.X == 1 && vel.Y == 0 {
			vel.X, vel.Y = 0, 1
		} else if vel.X == 0 && vel.Y == 1 {
			vel.X, vel.Y = -1, 0
		} else if vel.X == -1 && vel.Y == 0 {
			vel.X, vel.Y = 0, -1
		}

		return l.MoveGuard()
	}

	l.GuardPosition = next
	return false, false
}

func (l *Lab) GetDistinctGuardSpots() int {
	count := 0

	for _, line := range l.Grid {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}

	return count
}

// FindGuard updates the current GuardPosition and GuardVelocity,
// after locating the guard in the grid.
func (l *Lab) FindGuard() {
	l.GuardPosition = &Vector{}
	l.GuardVelocity = &Vector{}

	for y, line := range l.Grid {
		for x, char := range line {
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				l.GuardPosition.X = x
				l.GuardPosition.Y = y

				switch char {
				case '^':
					l.GuardVelocity.Y = -1

				case 'v':
					l.GuardVelocity.Y = 1

				case '>':
					l.GuardVelocity.X = 1

				case '<':
					l.GuardVelocity.X = -1
				}

				return
			}
		}
	}
}

// Reset grid back to starting state using input, and reset guard position/velocity.
func (l *Lab) Reset(lines []string) {
	l.Grid = make([][]byte, 0)

	for _, line := range lines {
		l.Grid = append(l.Grid, []byte(line))
	}

	l.FindGuard()
	l.PositionVelocityHistory = make([]string, 0)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Advent of Code 2024 - Day 6",
	Run: func(cmd *cobra.Command, args []string) {
		defer timer("day6")()

		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		// Build the lab
		lab := &Lab{}
		lab.Reset(lines)

		// Find the guard starting position, and direction.
		// The guard can be represented by v, >, <, or ^
		lab.FindGuard()

		// Move until the guard is gone
		for true {
			left, looped := lab.MoveGuard()

			if left || looped {
				break
			}
		}

		fmt.Printf("Number of distinct positions of the guard: %v\n", lab.GetDistinctGuardSpots())

		// Part 2, try to block the guard in and count # of loops
		loopCount := 0

		// Start by resetting state back to beginning
		lab.Reset(lines)

		for y := 0; y < len(lab.Grid); y++ {
			for x := 0; x < len(lab.Grid[y]); x++ {
				char := lab.Grid[y][x]

				// Can only add new obstactles on empty tiles
				if char == '.' {
					lab.Grid[y][x] = '#'

					// Loop and move the guard to check for loop or leaving map
					for true {
						left, looped := lab.MoveGuard()

						if looped {
							loopCount++
						}

						if left || looped {
							// Reset once done
							lab.Reset(lines)
							break
						}
					}
				}
			}
		}

		fmt.Printf("Number of possible loops: %v\n", loopCount)
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
