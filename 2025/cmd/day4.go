package cmd

import (
	"fmt"
	"log"

	"github.com/justinbrumley/advent-of-code/2025/utils"
	"github.com/spf13/cobra"
)

// getNeighbors returns a list of all neighbors that contain a roll of paper (@ symbol).
func getNeighborCount(x, y int, grid [][]byte) int {
	neighbors := 0

	for j := y - 1; j <= y+1; j++ {
		if j < 0 || j >= len(grid) {
			continue
		}

		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(grid[j]) || (i == x && j == y) {
				continue
			}

			if grid[j][i] == '@' {
				neighbors++
			}
		}
	}

	return neighbors
}

// isAccessible determines if the coordinate on the grid is accessible.
// A tile is accessible if < 4 neighbors contain a roll of paper (@ symbol).
func isAccessible(x, y int, grid [][]byte) bool {
	return getNeighborCount(x, y, grid) < 4
}

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Advent of Code 2025 - Day 4",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInputBytes(args[0])
		if err != nil {
			log.Fatal(err)
		}

		grid := make([][]byte, 0)
		for _, line := range lines {
			grid = append(grid, line)
		}

		accessible := 0
		for x, line := range grid {
			for y, tile := range line {
				if tile == '@' && isAccessible(x, y, grid) {
					accessible++
				}
			}
		}

		fmt.Printf("Part 1 Solution: %v\n", accessible)

		// Do it again, but remove tiles that are now accessible
		accessible = 0
		restart := true
		for restart {
			restart = false

			for y, line := range grid {
				for x, tile := range line {
					if tile == '@' && isAccessible(x, y, grid) {
						accessible++
						grid[y][x] = '.' // Move the paper roll out of the grid
						restart = true   // Restart check to make sure we didn't miss any rolls
					}
				}
			}
		}

		fmt.Printf("Part 2 Solution: %v\n", accessible)
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
