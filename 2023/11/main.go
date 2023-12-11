package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

func Abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

const SpaceMultiplier = 1_000_000

func (p *Point) DistanceTo(other *Point, expandedRows, expandedCols []int) int {
	distance := 0

	// Loop over expanded rows to add 1 million to distance
	for _, row := range expandedRows {
		if p.Y > other.Y && row < p.Y && row > other.Y {
			distance += (SpaceMultiplier - 1) // Minus 1 to account for math below
		} else if p.Y < other.Y && row > p.Y && row < other.Y {
			distance += (SpaceMultiplier - 1) // Minus 1 to account for math below
		}
	}

	for _, col := range expandedCols {
		if p.X > other.X && col < p.X && col > other.X {
			distance += (SpaceMultiplier - 1) // Minus 1 to account for math below
		} else if p.X < other.X && col > p.X && col < other.X {
			distance += (SpaceMultiplier - 1) // Minus 1 to account for math below
		}
	}

	// Add rest of distance using units of 1
	return distance + Abs(other.Y-p.Y) + Abs(other.X-p.X)
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	input := bytes.Split(data, []byte{'\n'})
	lines := make([][]byte, 0)

	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}

// For every empty row and column, add additional row or column
func ExpandUniverse(lines [][]byte) ([][]byte, []int, []int) {
	cols := 0
	expandedRows := make([]int, 0)
	expandedCols := make([]int, 0)

	// Start with empty rows
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		isEmpty := true

		for _, b := range line {
			if string(b) == "#" {
				isEmpty = false
			}
		}

		if len(line) > cols {
			cols = len(line)
		}

		if isEmpty {
			expandedRows = append(expandedRows, y)
		}
	}

	// Then empty columns
	for x := 0; x <= cols; x++ {
		isEmpty := true

		for _, line := range lines {
			if x < len(line) && string(line[x]) == "#" {
				isEmpty = false
			}
		}

		if isEmpty {
			expandedCols = append(expandedCols, x)
		}
	}

	return lines, expandedRows, expandedCols
}

func main() {
	lines := getInput()

	// Add a bunch of empty "space"
	lines, expandedRows, expandedCols := ExpandUniverse(lines)

	fmt.Printf("Expanded rows: %v\n", expandedRows)
	fmt.Printf("Expanded columns: %v\n", expandedCols)

	// Build galaxy map
	galaxies := make([]*Point, 0)

	for y, line := range lines {
		for x, b := range line {
			if string(b) == "#" {
				galaxies = append(galaxies, &Point{
					X: x,
					Y: y,
				})
			}
		}
	}

	// For each galaxy, find distance to other galaxies
	// and add them together
	sum := 0
	count := 0

	for i, galaxy := range galaxies {
		// Start AFTER this galaxy
		for j, other := range galaxies {
			if j <= i {
				continue
			}

			distance := galaxy.DistanceTo(other, expandedRows, expandedCols)
			sum += distance
			count++
		}
	}

	fmt.Printf("Sum of shortest paths after comparing %v pairs: %v\n", count, sum)
}
