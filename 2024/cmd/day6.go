package cmd

import (
	"fmt"
	"log"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type Vector struct {
	X int
	Y int
}

type Lab struct {
	Grid          [][]byte
	GuardPosition *Vector
	GuardVelocity *Vector
}

// Moves the guard in the direction they are currently facing.
// If the guard can't move forward, they rotate 90deg and try again.
// If the guard moves off the map entirely, returns false.
// This will also replace the previous space with an X to track where the guard has been.
func (l *Lab) MoveGuard() bool {
	// Get a vector for the next guard position
	next := &Vector{
		X: l.GuardPosition.X + l.GuardVelocity.X,
		Y: l.GuardPosition.Y + l.GuardVelocity.Y,
	}

	// Update current position with X
	l.Grid[l.GuardPosition.Y][l.GuardPosition.X] = 'X'

	// Check if guard moved off the map
	if next.X < 0 || next.Y < 0 || next.Y > len(l.Grid) || next.X > len(l.Grid[next.Y]) {
		// If so, then we are done and guard has left
		return false
	}

	// Check if next move is even viable
	nextChar := l.Grid[next.Y][next.X]
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

	// Successful move, so update position and return true
	l.GuardPosition = next
	return true
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
				fmt.Printf("Found the guard: %v, %v\n", x, y)
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

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Advent of Code 2024 - Day 6",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		// Build the lab
		lab := &Lab{
			Grid: make([][]byte, 0),
		}

		for _, line := range lines {
			lab.Grid = append(lab.Grid, []byte(line))
		}

		// Find the guard starting position, and direction.
		// The guard can be represented by v, >, <, or ^
		lab.FindGuard()

		// Move until the guard is gone
		for lab.MoveGuard() {
		}

		fmt.Printf("Number of distinct positions of the guard: %v\n", lab.GetDistinctGuardSpots())
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
