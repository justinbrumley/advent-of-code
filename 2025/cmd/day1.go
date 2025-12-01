package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2025/utils"
	"github.com/spf13/cobra"
)

func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

// turn takes the original position and amount to adjust.
// It returns the new position, and the amount of times it crossed zero.
func turn(position, amount int) (int, int) {
	numZeroes := 0
	originalPosition := position

	position += amount
	if amount > 0 {
		// Moving right
		for i := originalPosition + 1; i <= position; i++ {
			if i%100 == 0 {
				numZeroes += 1
			}
		}
	}

	if amount < 0 {
		// Moving left
		for i := originalPosition - 1; i >= position; i-- {
			if i%100 == 0 {
				numZeroes += 1
			}
		}
	}

	// Wrap around
	if position < 0 {
		position = (position % 100) + 100
	}

	if position > 99 {
		position = position % 100
	}

	return position, numZeroes
}

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of Code 2025 - Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		position := 50
		numZeroes := 0
		numClicks := 0

		for _, line := range lines {
			direction := []byte(line)[0]
			amount, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal(err)
			}

			if direction == 'L' {
				amount = 0 - amount
			}

			newPosition, clicks := turn(position, amount)

			position = newPosition
			numClicks += clicks

			if position == 0 {
				// Check AFTER calculations when determining how many times we landed on zero
				numZeroes += 1
			}
		}

		fmt.Printf("Final position: %v\n", position)
		fmt.Printf("Part 1 Solution: %v\n", numZeroes)
		fmt.Printf("Part 2 Solution: %v\n", numClicks)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
