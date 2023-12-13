package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	Rows [][]byte
	Cols [][]byte
}

func NewGrid(rows, cols [][]byte) *Grid {
	return &Grid{
		Rows: rows,
		Cols: cols,
	}
}

func (g *Grid) GetReflectionRow() int {
	// For every row...
	for i, row := range g.Rows {
		if i == len(g.Rows)-1 {
			break
		}

		// Check if next row matches
		if string(row) != string(g.Rows[i+1]) {
			continue
		}

		if i == 0 {
			// Already found the reflection line
			return i
		}

		// Check that all rows back to start (or up to end) match in reflection
		matches := true
		for j := i - 1; j >= 0; j-- {
			reflectionIdx := i + (i - j) + 1
			if reflectionIdx >= len(g.Rows) {
				break
			}

			// Check if they match
			if string(g.Rows[j]) != string(g.Rows[reflectionIdx]) {
				matches = false
				break
			}
		}

		if matches {
			return i
		}
	}

	return -1
}

func (g *Grid) GetReflectionCol() int {
	// For every col...
	for i, col := range g.Cols {
		if i == len(g.Cols)-1 {
			break
		}

		// Check if next col matches
		if string(col) != string(g.Cols[i+1]) {
			continue
		}

		if i == 0 {
			// Already found the reflection line
			return i
		}

		// Check that all cols back to start (or up to end) match in reflection
		matches := true
		for j := i - 1; j >= 0; j-- {
			reflectionIdx := i + (i - j) + 1
			if reflectionIdx >= len(g.Cols) {
				break
			}

			// Check if they match
			if string(g.Cols[j]) != string(g.Cols[reflectionIdx]) {
				matches = false
				break
			}
		}

		if matches {
			return i
		}
	}

	return -1
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
		lines = append(lines, line)
	}

	return lines
}

func main() {
	lines := getInput()

	rows := make([][]byte, 0)
	grids := make([]*Grid, 0)

	for _, line := range lines {
		if len(line) == 0 && len(rows) > 0 {
			// Build column strings out of rows
			cols := make([][]byte, 0)
			for _, row := range rows {
				for i, b := range row {
					if len(cols) <= i {
						cols = append(cols, []byte{})
					}

					cols[i] = append(cols[i], b)
				}
			}

			grids = append(grids, NewGrid(rows, cols))
			rows = make([][]byte, 0)
			continue
		}

		rows = append(rows, line)
	}

	sum := 0

	for _, grid := range grids {
		originalRow := grid.GetReflectionRow()
		originalCol := grid.GetReflectionCol()

		done := false

		for i, _ := range grid.Rows {
			if done {
				break
			}

			for j, _ := range grid.Cols {
				if done {
					break
				}

				// Flip value and i, j and check for reflections
				b := grid.Rows[i][j]

				if b == '#' {
					grid.Rows[i][j] = '.'
					grid.Cols[j][i] = '.'
				} else {
					grid.Rows[i][j] = '#'
					grid.Cols[j][i] = '#'
				}

				reflectionCol := grid.GetReflectionCol()

				if reflectionCol >= 0 && reflectionCol != originalCol {
					sum += reflectionCol + 1
					done = true
					break
				}

				reflectionRow := grid.GetReflectionRow()

				if reflectionRow >= 0 && reflectionRow != originalRow {
					sum += (reflectionRow + 1) * 100
					done = true
					break
				}

				// Put it back when done
				grid.Rows[i][j] = b
				grid.Cols[j][i] = b
			}
		}
	}

	fmt.Printf("Summarized reflection values: %v\n", sum)
}
