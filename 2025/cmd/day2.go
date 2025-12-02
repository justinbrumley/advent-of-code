package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"github.com/justinbrumley/advent-of-code/2025/utils"
	"github.com/spf13/cobra"
)

// An ID is invalid if there is a repeating sequence
// of numbers anywhere. This can be a single digit (e.g. 66),
// or a repeating pattern of digits (e.g. 3232).
func isInvalidId(value int) bool {
	re := pcre.MustCompile(`^(\d+)\1$`, 0)
	matches := re.MatcherString(strconv.Itoa(value), 0).Matches()
	return matches
}

func isInvalidIdExtended(value int) bool {
	re := pcre.MustCompile(`^(\d+)\1+$`, 0)
	matches := re.MatcherString(strconv.Itoa(value), 0).Matches()
	return matches
}

func getInvalidIds(start, end int) ([]int, []int) {
	invalidIds := make([]int, 0)
	extendedInvalidIds := make([]int, 0)

	for val := start; val <= end; val++ {
		if isInvalidId(val) {
			invalidIds = append(invalidIds, val)
		}

		if isInvalidIdExtended(val) {
			extendedInvalidIds = append(extendedInvalidIds, val)
		}
	}

	return invalidIds, extendedInvalidIds
}

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Advent of Code 2025 - Day 2",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		input := lines[0]

		ranges := strings.Split(input, ",")

		sum := 0
		sumExtended := 0

		for _, r := range ranges {
			parts := strings.Split(r, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			invalidIds, extendedInvalidIds := getInvalidIds(start, end)
			for _, val := range invalidIds {
				sum += val
			}

			for _, val := range extendedInvalidIds {
				sumExtended += val
			}
		}

		fmt.Printf("Part 1 Solution: %v\n", sum)
		fmt.Printf("Part 2 Solution: %v\n", sumExtended)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
