package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

const (
	MinDelta = 1
	MaxDelta = 3
)

// Remove a value from slice by index, and return the resulting slice.
func splice(values []int, idx int) []int {
	ret := make([]int, 0)

	for i, val := range values {
		if i != idx {
			ret = append(ret, val)
		}
	}

	return ret
}

// isSafe determines if a row of numbers is considered "safe".
// This means they only increase XOR decrease by 1-3 at a time.
func isSafe(values []int, dampen bool) bool {
	dir := 1
	if values[0] > values[1] {
		dir *= -1
	}

	for i, val := range values {
		if i == len(values)-1 {
			break
		}

		if (dir == 1 && val > values[i+1]) || (dir == -1 && val < values[i+1]) {
			// Wrong direction

			if dampen {
				// Try again with different variations (brute force)
				for j, _ := range values {
					if isSafe(splice(values, j), false) {
						return true
					}
				}
			}

			return false
		}

		delta := abs(values[i+1] - val)

		if delta < MinDelta || delta > MaxDelta {
			// Outside viable range

			if dampen {
				// Try again with the problem value(s) removed
				for j, _ := range values {
					if isSafe(splice(values, j), false) {
						return true
					}
				}
			}

			return false
		}
	}

	return true
}

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Advent of Code 2024 - Day 2",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.GetInput("inputs/day2")
		if err != nil {
			log.Fatal(err)
		}

		safeReports := 0
		safeReportsWithDampening := 0

		re := regexp.MustCompile(`\s+`)

		for _, line := range lines {
			parts := re.Split(line, -1)
			values := make([]int, len(parts))

			for i, val := range parts {
				num, _ := strconv.Atoi(val)
				values[i] = num
			}

			if isSafe(values, false) {
				safeReports += 1
			}

			if isSafe(values, true) {
				safeReportsWithDampening += 1
			}
		}

		fmt.Printf("Number of safe reports: %v\n", safeReports)
		fmt.Printf("Number of safe reports, with dampening: %v\n", safeReportsWithDampening)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
