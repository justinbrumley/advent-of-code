package cmd

import (
	"fmt"
	"log"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type Grid struct {
	Values [][]byte
}

type Point struct {
	X int
	Y int
}

func (g *Grid) Read(from, to Point) string {
	xDir := 1
	yDir := 1

	if from.X > to.X {
		xDir = -1
	} else if from.X == to.X {
		xDir = 0
	}

	if from.Y > to.Y {
		yDir = -1
	} else if from.Y == to.Y {
		yDir = 0
	}

	out := make([]byte, 0)

	for x, y := from.X, from.Y; x != to.X || y != to.Y; x, y = x+xDir, y+yDir {
		if x < 0 || y < 0 {
			break
		}

		if y >= len(g.Values) || x >= len(g.Values[y]) {
			break
		}

		out = append(out, g.Values[y][x])
	}

	return string(out)
}

func (g *Grid) CountFromPosition(x, y int) int {
	from := Point{X: x, Y: y}
	length := 4
	strings := []string{
		g.Read(from, Point{X: x, Y: y - length}),          // Up
		g.Read(from, Point{X: x, Y: y + length}),          // Down
		g.Read(from, Point{X: x - length, Y: y}),          // Left
		g.Read(from, Point{X: x + length, Y: y}),          // Right
		g.Read(from, Point{X: x + length, Y: y - length}), // Up-Right
		g.Read(from, Point{X: x - length, Y: y - length}), // Up-Left
		g.Read(from, Point{X: x - length, Y: y + length}), // Down-Left
		g.Read(from, Point{X: x + length, Y: y + length}), // Down-Right
	}

	count := 0
	for _, val := range strings {
		if val == "XMAS" {
			count++
		}
	}

	return count
}

// funny name
func (g *Grid) MasFromPos(x, y int) bool {
	// Checking characters in X formation around the current position
	topLeftToBottomRight := g.Read(
		Point{X: x - 1, Y: y - 1},
		Point{X: x + 2, Y: y + 2}, // Off by one tomfoolery, but too lazy to fix in Read()
	)

	topRightToBottomLeft := g.Read(
		Point{X: x + 1, Y: y - 1},
		Point{X: x - 2, Y: y + 2}, // Off by one tomfoolery, but too lazy to fix in Read()
	)

	if (topLeftToBottomRight == "MAS" || topLeftToBottomRight == "SAM") && (topRightToBottomLeft == "MAS" || topRightToBottomLeft == "SAM") {
		return true
	}

	return false
}

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Advent of Code 2024 - Day 4",
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := utils.GetInput("inputs/day4")
		if err != nil {
			log.Fatal(err)
		}

		// Build a grid of letters
		values := make([][]byte, 0)
		for _, line := range lines {
			row := []byte(line)
			values = append(values, row)
		}

		grid := &Grid{Values: values}

		count := 0
		countMas := 0

		for y, row := range grid.Values {
			for x, char := range row {
				if char == 'X' {
					count += grid.CountFromPosition(x, y)
				}

				if char == 'A' && grid.MasFromPos(x, y) {
					countMas += 1
				}
			}
		}

		fmt.Printf("Total # of XMAS values: %v\n", count)
		fmt.Printf("Total # of X-MAS values: %v\n", countMas)
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
