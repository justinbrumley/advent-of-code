package cmd

import (
	"fmt"
	"log"

	"github.com/justinbrumley/advent-of-code/2024/utils"
	"github.com/spf13/cobra"
)

type AntennaMap struct {
	Grid         [][]byte
	AntiNodeGrid [][]byte
}

func (m *AntennaMap) PrintGrid() {
	for _, line := range m.Grid {
		fmt.Println(string(line))
	}
}

func (m *AntennaMap) PrintAntiNodeGrid() {
	for _, line := range m.AntiNodeGrid {
		fmt.Println(string(line))
	}
}

func NewAntennaMap(grid [][]byte) *AntennaMap {
	antiNodeGrid := make([][]byte, 0)

	for _, line := range grid {
		l := make([]byte, 0)

		for range line {
			l = append(l, '.')
		}

		antiNodeGrid = append(antiNodeGrid, l)
	}

	return &AntennaMap{
		Grid:         grid,
		AntiNodeGrid: antiNodeGrid,
	}
}

func (m *AntennaMap) FindAll(ch byte, exclude Vector) []Vector {
	vectors := make([]Vector, 0)

	for y, line := range m.Grid {
		for x, char := range line {
			if char != '.' && char == ch && !(y == exclude.Y && x == exclude.X) {
				vectors = append(vectors, Vector{
					X: x,
					Y: y,
				})
			}
		}
	}

	return vectors
}

func GetDistance(p1, p2 Vector) Vector {
	return Vector{
		X: p2.X - p1.X,
		Y: p2.Y - p1.Y,
	}
}

// Place antinodes on grid around p1, relative to p2.
func (m *AntennaMap) PlaceAntiNodes(p1, p2 Vector) {
	d := GetDistance(p1, p2)

	x := p1.X
	y := p1.Y

	for true {
		vector := Vector{
			X: x,
			Y: y,
		}

		if vector.X >= 0 && vector.Y >= 0 && vector.Y < len(m.AntiNodeGrid) && vector.X < len(m.AntiNodeGrid[vector.Y]) {
			m.AntiNodeGrid[vector.Y][vector.X] = '#'
			x -= d.X
			y -= d.Y
			continue
		}

		break
	}
}

func (m *AntennaMap) BuildAntiNodes() {
	for y, line := range m.Grid {
		for x, char := range line {
			vector := Vector{X: x, Y: y}
			vectors := m.FindAll(char, vector)

			for _, otherVector := range vectors {
				m.PlaceAntiNodes(vector, otherVector)
			}
		}
	}
}

func (m *AntennaMap) CountDistinctAntiNodes() int {
	count := 0

	for _, line := range m.AntiNodeGrid {
		for _, char := range line {
			if char == '#' {
				count++
			}
		}
	}

	return count
}

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Advent of Code 2024 - Day 8",
	Run: func(cmd *cobra.Command, args []string) {
		defer timer("day7")()

		if len(args) < 1 {
			log.Fatal("Missing required argument: <input_file>")
		}

		lines, err := utils.GetInputBytes(args[0])
		if err != nil {
			log.Fatal(err)
		}

		m := NewAntennaMap(lines)

		m.BuildAntiNodes()

		m.PrintGrid()
		fmt.Println("------------------------")
		m.PrintAntiNodeGrid()

		fmt.Printf("Count of distinct anti-nodes: %v\n", m.CountDistinctAntiNodes())
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
