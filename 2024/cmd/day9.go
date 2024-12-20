package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

// FindValueSpaces returns the index of the start of the range,
// going backwards from the start index.
func FindValueSpaces(fs []int, start int) int {
	val := fs[start]

	for i := start; i >= 0; i-- {
		if fs[i] != val {
			return i + 1
		}
	}

	return 0
}

// FindFreeSpaces will return the index of the first free space range of
// length or greater. If not found, returns -1.
func FindFreeSpaces(fs []int, length, end int) int {
	for i := 0; i < end; i++ {
		if fs[i] != -1 {
			continue
		}

		// Found free space, so get length of range
		for j := i; j < end; j++ {
			if fs[j] == -1 {
				continue
			}

			// End of range
			if j-i >= length {
				return i
			}

			break
		}
	}

	return -1
}

var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Advent of Code 2024 - Day 9",
	Run: func(cmd *cobra.Command, args []string) {
		defer timer("day7")()

		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInputBytes(args[0])
		if err != nil {
			log.Fatal(err)
		}

		// Everything on one line for the lulz
		filesystem := make([]int, 0)

		for i, char := range lines[0] {
			count, _ := strconv.Atoi(string(char))
			id := i / 2

			if i%2 == 0 {
				for j := 0; j < count; j++ {
					filesystem = append(filesystem, id)
				}
			} else {
				for j := 0; j < count; j++ {
					filesystem = append(filesystem, -1)
				}
			}
		}

		// Starting from the back, move characters to first free space
		for i := len(filesystem) - 1; i >= 0; i-- {
			value := filesystem[i]

			if value == -1 {
				continue
			}

			// Get range
			start := FindValueSpaces(filesystem, i)
			length := (i - start) + 1

			// Find free space of span size or greater
			free := FindFreeSpaces(filesystem, length, start-1)

			if free >= 0 {
				// Found some space, so move range there
				for j := 0; j < length; j++ {
					// Swap!
					filesystem[free+j], filesystem[start+j] = filesystem[start+j], filesystem[free+j]
				}
			}

			// Skip past the current range
			i = start
		}

		checksum := 0

		for i, val := range filesystem {
			if val == -1 {
				continue
			}

			checksum += i * val
		}

		fmt.Printf("Checksum: %v\n", checksum)
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
