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

func main() {
	lines = getInput()
	startPoint := GetStartPoint()

	// Loop over possible pipes to look for largest loop
	for _, pipe := range Pipes {
		seq := make([]*Point, 1)
		seq[0] = &Point{
			X:     startPoint.X,
			Y:     startPoint.Y,
			Value: pipe,
		}

		seq = append(seq, GetNextPoint(nil, seq[0]))

		for seq[len(seq)-1].Value != "S" {
			prev := seq[len(seq)-2]
			curr := seq[len(seq)-1]
			seq = append(seq, GetNextPoint(prev, curr))
		}

		fmt.Printf("Further point on loop %v: %v\n", pipe, len(seq)/2)
	}
}
