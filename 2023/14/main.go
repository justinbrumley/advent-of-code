package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	SquareRock = '#'
	EmptySpace = '.'
	Rock       = 'O'
)

type Grid struct {
	Values [][]byte
}

func (g *Grid) GetRowCount() int {
	return len(g.Values)
}

func (g *Grid) GetColumnCount() int {
	if len(g.Values) > 0 {
		return len(g.Values[0])
	}

	return 0
}

// RollNorth will attempt to roll the rock at position x,y north
func (g *Grid) RollNorth(x, y int) {
	for j := y - 1; j >= 0; j-- {
		val := g.Values[j][x]

		if val == Rock || val == SquareRock {
			// Set original space to empty
			g.Values[y][x] = EmptySpace

			// Move to just below other rock
			g.Values[j+1][x] = Rock

			break
		} else if j == 0 && val == EmptySpace {
			// Set original space to empty
			g.Values[y][x] = EmptySpace

			// Move to empty space at top
			g.Values[0][x] = Rock

			break
		}
	}
}

// TiltNorth rolls all round rocks up the grid until reaching the top or square rocks
func (g *Grid) TiltNorth() {
	// Do one column at a time
	for x := 0; x < g.GetColumnCount(); x++ {
		for y := 1; y < g.GetRowCount(); y++ {
			tile := g.Values[y][x]

			if tile == Rock {
				g.RollNorth(x, y)
			}
		}
	}
}

// GetTotalLoad returns the total load of rocks, based on position in the grid
func (g *Grid) GetTotalLoad() int {
	// Starting weight is the # of rows
	weight := g.GetRowCount()
	sum := 0

	for _, row := range g.Values {
		for _, val := range row {
			if val == Rock {
				sum += weight
			}
		}

		weight--
	}

	return sum
}

// Rotate the grid clockwise (north, west, south, east)
func (g *Grid) Rotate() {
	newValues := make([][]byte, 0)

	// Loop over columns, reversed to get new rows
	for i := 0; i < g.GetColumnCount(); i++ {
		row := make([]byte, g.GetColumnCount())

		for j := 0; j < g.GetRowCount(); j++ {
			val := g.Values[g.GetRowCount()-j-1][i]
			row[j] = val
		}

		newValues = append(newValues, row)
	}

	g.Values = newValues
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

func ContainsLoop(snapshots [][]byte) bool {
	for i, snapshot := range snapshots {
		for j := i + 1; j < len(snapshots); j++ {
			if string(snapshot) == string(snapshots[j]) {
				return true
			}
		}
	}

	return false
}

func GetLoop(snapshots [][]byte) ([][]byte, int, int) {
	start := -1
	end := -1

	for i, snapshot := range snapshots {
		if start > -1 {
			break
		}

		for j := i + 1; j < len(snapshots); j++ {
			if string(snapshot) == string(snapshots[j]) {
				start = i
				end = j
				break
			}
		}
	}

	return snapshots[start:end], start, end
}

func main() {
	grid := &Grid{
		Values: getInput(),
	}

	snapshots := make([][]byte, 0)

	// Cycle 1 billion times, rotating 4 times per cycle
	for i := 0; i < 1e9; i++ {
		// North
		grid.TiltNorth()
		grid.Rotate()

		// West
		grid.TiltNorth()
		grid.Rotate()

		// South
		grid.TiltNorth()
		grid.Rotate()

		// East
		grid.TiltNorth()
		grid.Rotate()

		// Store snapshot to track loops
		snapshots = append(snapshots, bytes.Join(grid.Values, []byte("\n")))
		if ContainsLoop(snapshots) && i > 20 {
			break
		}
	}

	loop, start, _ := GetLoop(snapshots)
	idx := (1e9 - start - 1) % len(loop)

	finalGrid := &Grid{
		Values: bytes.Split(loop[idx], []byte("\n")),
	}

	fmt.Printf("1,000,000,000th load: %v\n", finalGrid.GetTotalLoad())
}
