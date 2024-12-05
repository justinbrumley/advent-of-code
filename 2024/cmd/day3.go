package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Advent of Code 2024 - Day 3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInput(args[0])
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`(mul|don't|do)\(((\d{1,3}),(\d{1,3}))?\)`)

		enabled := true
		total := 0

		for _, line := range lines {
			instructions := re.FindAllStringSubmatch(line, -1)

			for _, instruction := range instructions {
				cmd := instruction[1]

				if cmd == "mul" && enabled {
					x, _ := strconv.Atoi(instruction[3])
					y, _ := strconv.Atoi(instruction[4])
					total += (x * y)
				} else if cmd == "don't" {
					enabled = false
				} else if cmd == "do" {
					enabled = true
				}
			}
		}

		fmt.Printf("Sum of enabled mul() instructions: %v\n", total)
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
