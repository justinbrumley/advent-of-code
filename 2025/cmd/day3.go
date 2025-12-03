package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/justinbrumley/advent-of-code/2025/utils"
	"github.com/spf13/cobra"
)

// getRatings returns the list of voltage ratings for batteries in the bank.
func getRatings(line []byte) []int {
	ratings := make([]int, len(line))

	for i, b := range line {
		rating, _ := strconv.Atoi(string(b))
		ratings[i] = rating
	}

	return ratings
}

// max returns the highest value within the array, and it's index.
// values are expected to be all positive.
func max(values []int) (int, int) {
	maxValue := -1
	maxIndex := -1

	for idx, val := range values {
		if val > maxValue {
			maxValue = val
			maxIndex = idx
		}
	}

	return maxValue, maxIndex
}

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Advent of Code 2025 - Day 3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInputBytes(args[0])
		if err != nil {
			log.Fatal(err)
		}

		sum := 0
		for _, bank := range lines {
			ratings := getRatings(bank)

			// First find largest number up to length - 1
			first, idx := max(ratings[:len(ratings)-1])

			// Next, find the largest number AFTER the first digit
			second, _ := max(ratings[idx+1:])

			// Combine the two digits, parse as int, and add to total
			num, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
			sum += num
		}

		fmt.Printf("Part 1 Solution: %v\n", sum)

		/* Part Dos */
		sum = 0

		for _, bank := range lines {
			ratings := getRatings(bank)

			digits := make([]string, 12)
			prevIdx := 0

			// This time we need to get TWELVE WHOLE DIGITS from each line
			for i := 0; i < 12; i++ {
				// Get highest digit up to len - i, so we have room left for remaining digits
				val, idx := max(ratings[prevIdx : len(ratings)-(12-i)+1])
				digits[i] = fmt.Sprintf("%d", val)
				prevIdx += idx + 1
			}

			// Concat all the digits into one number
			num, _ := strconv.Atoi(strings.Join(digits, ""))

			// Add to total
			sum += num
		}

		fmt.Printf("Part 2 Solution: %v\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
