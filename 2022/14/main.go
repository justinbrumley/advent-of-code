package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Point struct {
	X    int
	Y    int
	Type string
}

var StartPoint = &Point{
	X: 500,
	Y: 0,
}

func ClearConsole() {
	fmt.Println("\033[2J")
}

type Grid struct {
	Points map[string]*Point
	Floor  int

	// Track some more bounds just in case we want to do some fancy printing of the grid
	Ceiling int
	Left    int
	Right   int
}

func NewGrid() *Grid {
	return &Grid{
		Points:  make(map[string]*Point),
		Ceiling: 0,
		Floor:   0,
	}
}

func (g *Grid) AddPoint(point *Point) {
	g.Points[fmt.Sprintf("%d %d", point.X, point.Y)] = point

	if point.Type != "sand" && point.Y+2 > g.Floor {
		g.Floor = point.Y + 2
	}

	if point.X < g.Left {
		g.Left = point.X
	}

	if point.X > g.Right {
		g.Right = point.X
	}
}

func (g *Grid) GetPoint(x, y int) *Point {
	return g.Points[fmt.Sprintf("%d %d", x, y)]
}

func (g *Grid) DrawGridAroundPoint(p *Point) {
	time.Sleep(100 * time.Millisecond)
	ClearConsole()

	size := 40

	for y := p.Y - (size / 2); y < p.Y+(size/2); y++ {
		line := ""

		for x := p.X - size; x < p.X+size; x++ {
			if x == p.X && y == p.Y {
				line += "+"
				continue
			}

			if x == 500 && y == 0 {
				line += "S"
				continue
			}

			p := g.GetPoint(x, y)
			if p == nil {
				line += "."
			} else if p.Type == "sand" {
				line += "o"
			} else {
				line += "#"
			}
		}

		fmt.Println(line)
	}
}

func (g *Grid) DrawGrid() {
	for y := 0; y < 150; y++ {
		line := ""

		for x := 450; x < 550; x++ {
			if x == 500 && y == 0 {
				line += "S"
				continue
			}

			p := g.GetPoint(x, y)
			if p == nil {
				line += "."
			} else if p.Type == "sand" {
				line += "o"
			} else {
				line += "#"
			}
		}

		fmt.Println(line)
	}
}

// DropSand will drop a single sand tile down until it can't move anymore
// Returns true if sand landed somewhere, false if it fell forever
func (g *Grid) DropSand() bool {
	// Just treat sand like a point that starts at 500, 0
	sand := &Point{
		X:    500,
		Y:    0,
		Type: "sand",
	}

	for {
		// intoTheAbyss := true

		// Part Two
		// Check if next tile down is the floor and stop if so
		if sand.Y+1 >= g.Floor {
			g.AddPoint(sand)
			g.DrawGridAroundPoint(sand)
			return true
		}

		pointDown := g.GetPoint(sand.X, sand.Y+1)
		pointDownLeft := g.GetPoint(sand.X-1, sand.Y+1)
		pointDownRight := g.GetPoint(sand.X+1, sand.Y+1)

		// Point always moves down
		if pointDown == nil {
			// Move down
			sand.Y++
		} else if pointDownLeft == nil {
			// Move down and left
			sand.X--
			sand.Y++
		} else if pointDownRight == nil {
			// Move down and right
			sand.X++
			sand.Y++
		} else {
			// Nowhere to move, so add point to grid, and return true
			g.AddPoint(sand)
			g.DrawGridAroundPoint(sand)

			// If current point is the starting point, return false to signify we can't drop more sand
			return sand.X != StartPoint.X || sand.Y != StartPoint.Y
		}
	}
}

// GetPointsInPath takes two points and returns all points in a straight line between them
func GetPointsInPath(from, to *Point) []Point {
	points := make([]Point, 0)

	if from.X == to.X {
		// Move vertically
		start := from
		end := to

		if start.Y > end.Y {
			start = to
			end = from
		}

		for i := start.Y; i <= end.Y; i++ {
			points = append(points, Point{
				X: start.X, // Constant
				Y: i,
			})
		}
	} else if from.Y == to.Y {
		// Move horizontally
		start := from
		end := to

		if start.X > end.X {
			start = to
			end = from
		}

		for i := start.X; i <= end.X; i++ {
			points = append(points, Point{
				X: i,
				Y: start.Y, // Constant
			})
		}
	}

	return points
}

// getInput parses the input file in the current directory, and returns the contents split by newlines
func getInput() ([][]byte, error) {
	file, err := os.Open("./input")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0)

	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	return lines, nil
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	grid := NewGrid()

	// Build our grid
	for _, line := range input {
		positions := bytes.Split(line, []byte(" -> "))

		for i := 0; i < len(positions)-1; i++ {
			coordinates := bytes.Split(positions[i], []byte(","))

			x, _ := strconv.Atoi(string(coordinates[0]))
			y, _ := strconv.Atoi(string(coordinates[1]))

			from := &Point{
				X: x,
				Y: y,
			}

			coordinates = bytes.Split(positions[i+1], []byte(","))

			x, _ = strconv.Atoi(string(coordinates[0]))
			y, _ = strconv.Atoi(string(coordinates[1]))

			to := &Point{
				X: x,
				Y: y,
			}

			// This will fetch a list of all points in the path, including the from and to values
			points := GetPointsInPath(from, to)

			for _, point := range points {
				grid.AddPoint(&point)
			}
		}
	}

	// Let's start dropping some sand
	// Count how many grains of sand actually come to rest
	count := 0
	for {
		count++

		if !grid.DropSand() {
			break
		}
	}

	// grid.DrawGrid()
	// fmt.Println("=============================")
	fmt.Printf("%d pieces of sand came to rest\n", count)
}
