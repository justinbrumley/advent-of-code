package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/justinbrumley/advent-of-code/2025/utils"
	"github.com/spf13/cobra"
)

func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func mergeOverlappingRanges(values [][]int) [][]int {
	validRanges := make([][]int, 0)
	processed := make([]bool, len(values))

	for i, r := range values {
		if processed[i] == true {
			continue
		}

		start := r[0]
		end := r[1]

		restart := true
		for restart {
			restart = false

			for j, r2 := range values {
				if i == j || processed[j] == true {
					continue
				}

				// If not overlapping, then continue
				s := r2[0]
				e := r2[1]

				if end < s || start > e {
					continue
				}

				// Otherwise we got something overlapping, so extend start/end
				if s < start {
					start = s
				}

				if e > end {
					end = e
				}

				// Mark range as processed, and ensure loop is re-ran to catch
				// any previous ranges now that we've expanded.
				restart = true
				processed[j] = true
			}
		}

		// Flag current range as processed so it's skipped in future iterations
		processed[i] = true
		validRanges = append(validRanges, []int{start, end})
	}

	return validRanges
}

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Advent of Code 2025 - Day 5",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		validRanges := make([][]int, 0)
		checking := false

		totalFresh := 0
		validIngredients := 0

		for _, line := range lines {
			if string(line) == "" {
				checking = true
				validRanges = mergeOverlappingRanges(validRanges)

				for _, r := range validRanges {
					totalFresh += (r[1] - r[0]) + 1
				}

				continue
			}

			if !checking {
				parts := strings.Split(line, "-")
				start, _ := strconv.Atoi(parts[0])
				end, _ := strconv.Atoi(parts[1])

				r := []int{start, end}
				validRanges = append(validRanges, r)
			} else {
				id, _ := strconv.Atoi(line)

				// Check if ID falls within any of the valid ranges
				for _, r := range validRanges {
					start := r[0]
					end := r[1]

					if id >= start && id <= end {
						validIngredients++
						break
					}
				}
			}
		}

		fmt.Printf("Part 1 Solution: %v\n", validIngredients)
		fmt.Printf("Part 2 Solution: %v\n", totalFresh)
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
