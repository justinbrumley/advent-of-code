package cmd

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

func sum(values []int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func countOccurences(values []int, value int) int {
	occurences := 0

	for _, val := range values {
		if val == value {
			occurences += 1
		}
	}

	return occurences
}

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of Code 2024 - Day 1",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.GetInput("inputs/day1")
		if err != nil {
			log.Fatal(err)
		}

		list1 := make([]int, len(lines))
		list2 := make([]int, len(lines))
		re := regexp.MustCompile(`\s+`)

		for idx, line := range lines {
			parts := re.Split(line, -1)

			i, _ := strconv.Atoi(parts[0])
			j, _ := strconv.Atoi(parts[1])

			list1[idx] = i
			list2[idx] = j
		}

		sort.Slice(list1, func(i, j int) bool {
			return list1[i] < list1[j]
		})

		sort.Slice(list2, func(i, j int) bool {
			return list2[i] < list2[j]
		})

		distances := make([]int, len(list1))

		for idx, _ := range list1 {
			distances[idx] = abs(list1[idx] - list2[idx])
		}

		// Sum the distances for part 1 solution
		fmt.Printf("Part 1 Solution: %v\n", sum(distances))

		score := 0

		// Loop over first list, and generate similarity scores
		for _, val := range list1 {
			// Count how many times the number appears in list2
			count := countOccurences(list2, val)

			// Multiple it by the value and add to overall score
			score += (val * count)
		}

		fmt.Printf("Part 2 Solution: %v\n", score)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
