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
	/*
		[1,2,3]
		[4,5,6]
		[7,8,9]

		[7,4,1]
		[8,5,2]
		[9,6,3]
	*/

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

func GetLoopLength(snapshots [][]byte) int {
	for i, snapshot := range snapshots {
		for j := i + 1; j < len(snapshots); j++ {
			if string(snapshot) == string(snapshots[j]) {
				return j - i
			}
		}
	}

	return -1
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
		if ContainsLoop(snapshots) {
			break
		}
	}

	// Rebuild grid using snapshot
	loopLength := GetLoopLength(snapshots)
	offset := len(snapshots) - loopLength
	idx := (1e9 - offset - 1) % loopLength

	fmt.Printf("Offset, Loop length: %v %v\n", offset, loopLength)
	fmt.Printf("Index of final grid: %v\n", idx)

	for i := offset; i < offset+loopLength; i++ {
		g := &Grid{
			Values: bytes.Split(snapshots[i], []byte("\n")),
		}

		fmt.Printf("Snapshot #%v load: %v\n", i, g.GetTotalLoad())
	}
}
