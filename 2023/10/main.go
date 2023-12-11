package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var Pipes = []string{
	"|", // Vertical
	"-", // Horizontal
	"J", // Up/Left
	"7", // Down/Left
	"L", // Up/Right
	"F", // Down/Right
}

type Point struct {
	X     int
	Y     int
	Value string
}

type Bounds struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

var lines [][]byte

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func GetStartPoint() *Point {
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "S" {
				return &Point{
					X:     x,
					Y:     y,
					Value: "S",
				}
			}
		}
	}

	return nil
}

func GetPointAt(x, y int) *Point {
	if y >= len(lines) || y < 0 {
		return nil
	}

	if x >= len(lines[y]) || x < 0 {
		return nil
	}

	return &Point{
		X:     x,
		Y:     y,
		Value: string(lines[y][x]),
	}
}

func GetNextPoint(previous, current *Point) *Point {
	switch current.Value {
	case "|": // Vertical
		// If previous point is above, or nil, then go down
		if previous == nil || previous.Y < current.Y {
			return GetPointAt(current.X, current.Y+1)
		} else {
			return GetPointAt(current.X, current.Y-1)
		}

	case "-": // Horizontal
		// If previous point is left, or nil, then go right
		if previous == nil || previous.X < current.X {
			return GetPointAt(current.X+1, current.Y)
		} else {
			return GetPointAt(current.X-1, current.Y)
		}

	case "J": // Up/Left
		// If previous point is up, or nil, then go left
		if previous == nil || previous.Y < current.Y {
			return GetPointAt(current.X-1, current.Y)
		} else {
			return GetPointAt(current.X, current.Y-1)
		}

	case "7": // Down/Left
		// If previous point is down, or nil, then go left
		if previous == nil || previous.Y > current.Y {
			return GetPointAt(current.X-1, current.Y)
		} else {
			return GetPointAt(current.X, current.Y+1)
		}

	case "L": // Up/Right
		// If previous point is up, or nil, then go right
		if previous == nil || previous.Y < current.Y {
			return GetPointAt(current.X+1, current.Y)
		} else {
			return GetPointAt(current.X, current.Y-1)
		}

	case "F": // Down/Right
		// If previous point is down, or nil, then go right
		if previous == nil || previous.Y > current.Y {
			return GetPointAt(current.X+1, current.Y)
		} else {
			return GetPointAt(current.X, current.Y+1)
		}
	}

	return nil
}

func (p *Point) PointsUp() bool {
	return p != nil && (p.Value == "|" || p.Value == "J" || p.Value == "L")
}

// IsEnclosed checks if the provided point is enclosed in the loop
func (p *Point) IsEnclosed(loop []*Point) bool {
	isInside := false

	for _, point := range loop {
		if point.Y == p.Y && point.X < p.X && point.PointsUp() {
			isInside = !isInside
		}
	}

	return isInside
}

func (p *Point) IsValidPipe() bool {
	above := GetPointAt(p.X, p.Y-1)
	below := GetPointAt(p.X, p.Y+1)
	left := GetPointAt(p.X-1, p.Y)
	right := GetPointAt(p.X+1, p.Y)

	connectsUp := above != nil && (above.Value == "|" || above.Value == "F" || above.Value == "7")
	connectsDown := below != nil && (below.Value == "|" || below.Value == "J" || below.Value == "L")

	connectsLeft := left != nil && (left.Value == "-" || left.Value == "F" || left.Value == "L")
	connectsRight := right != nil && (right.Value == "-" || right.Value == "J" || right.Value == "7")

	switch p.Value {
	case "|":
		return connectsUp && connectsDown

	case "-":
		return connectsLeft && connectsRight

	case "J":
		return connectsLeft && connectsUp

	case "7":
		return connectsLeft && connectsDown

	case "L":
		return connectsRight && connectsUp

	case "F":
		return connectsRight && connectsDown
	}

	return false
}

func main() {
	lines = getInput()
	startPoint := GetStartPoint()

	var loop []*Point

	// Loop over possible pipes to look for largest loop
	for _, pipe := range Pipes {
		seq := make([]*Point, 1)
		seq[0] = &Point{
			X:     startPoint.X,
			Y:     startPoint.Y,
			Value: pipe,
		}

		if !seq[0].IsValidPipe() {
			continue
		}

		seq = append(seq, GetNextPoint(nil, seq[0]))

		for {
			prev := seq[len(seq)-2]
			curr := seq[len(seq)-1]
			next := GetNextPoint(prev, curr)

			if next == nil || next.Value == "S" {
				break
			}

			seq = append(seq, next)
		}

		if len(seq) > len(loop) {
			loop = seq
			startPoint.Value = pipe
		}
	}

	fmt.Printf("Further point on loop: %v\n", len(loop)/2)

	enclosedPoints := 0

	// Look for dots surrounded by unequal amount of pipes
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			point := &Point{
				X:     x,
				Y:     y,
				Value: string(lines[y][x]),
			}

			// Ignore pipes connects to loop.
			// We only care about ground + junk pipes.
			inLoop := false
			for _, p := range loop {
				if p.X == point.X && p.Y == point.Y {
					inLoop = true
				}
			}

			if !inLoop && point.IsEnclosed(loop) {
				enclosedPoints++
			}
		}
	}

	fmt.Printf("Number of enclosed ground tiles: %v\n", enclosedPoints)
}
