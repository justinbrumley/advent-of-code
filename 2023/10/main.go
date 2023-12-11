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

func (p *Point) IsCorner() bool {
	for _, pipe := range Pipes[2:] {
		if pipe == p.Value {
			return true
		}
	}

	return false
}

type Bounds struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

type Side struct {
	P1 *Point
	P2 *Point
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

func GetBounds(points []*Point) *Bounds {
	left := -1
	right := -1
	top := -1
	bottom := -1

	for _, point := range points {
		if left == -1 || point.X < left {
			left = point.X
		}

		if right == -1 || point.X > right {
			right = point.X
		}

		if top == -1 || point.Y < top {
			top = point.Y
		}

		if bottom == -1 || point.Y > bottom {
			bottom = point.Y
		}
	}

	return &Bounds{
		Top:    top,
		Bottom: bottom,
		Left:   left,
		Right:  right,
	}
}

// IsEnclosed checks if the provided point is enclosed in the loop
func (p *Point) IsEnclosed(sides []*Side) bool {
	// D = (x2 - x1) * (yp - y1) - (xp - x1) * (y2 - y1)
	prev := 0

	// 2, 6

	for _, s := range sides {
		d := (s.P2.X-s.P1.X)*(p.Y-s.P1.Y) - (p.X-s.P1.X)*(s.P2.Y-s.P1.Y)

		if p.X == 2 && p.Y == 6 {
			fmt.Printf("D Value (%v) for (%v, %v) -> (%v, %v)\n", d, s.P1.X, s.P1.Y, s.P2.X, s.P2.Y)
		}

		if prev == 0 {
			prev = d
		} else if d < 0 && prev > 0 {
			// fmt.Printf("(%v, %v) Values on wrong side: %v and %v\n", p.X, p.Y, d, prev)
			return false
		} else if d > 0 && prev < 0 {
			// fmt.Printf("(%v, %v), Values on wrong side: %v and %v\n", p.X, p.Y, d, prev)
			return false
		}
	}

	return true
}

func main() {
	lines = getInput()
	startPoint := GetStartPoint()

	var loop []*Point
	var sides []*Side

	// Loop over possible pipes to look for largest loop
	for _, pipe := range Pipes {
		seq := make([]*Point, 1)
		seq[0] = &Point{
			X:     startPoint.X,
			Y:     startPoint.Y,
			Value: pipe,
		}

		seq = append(seq, GetNextPoint(nil, seq[0]))

		s := make([]*Side, 0)
		sideStart := seq[0]

		if seq[len(seq)-1].IsCorner() {
			s = append(s, &Side{
				P1: sideStart,
				P2: seq[len(seq)-1],
			})

			sideStart = seq[len(seq)-1]
		}

		for seq[len(seq)-1].Value != "S" {
			prev := seq[len(seq)-2]
			curr := seq[len(seq)-1]
			next := GetNextPoint(prev, curr)

			if next == nil {
				break
			}

			seq = append(seq, next)

			if next.IsCorner() {
				s = append(s, &Side{
					P1: sideStart,
					P2: seq[len(seq)-1],
				})

				sideStart = seq[len(seq)-1]
			}
		}

		if len(seq) > len(loop) {
			loop = seq
			sides = s
		}
	}

	fmt.Printf("Further point on loop: %v\n", len(loop)/2)

	fmt.Printf("Sides (%v):\n", len(sides))
	for _, s := range sides {
		fmt.Printf("(%v, %v) -> (%v, %v)\n", s.P1.X, s.P1.Y, s.P2.X, s.P2.Y)
	}

	// Calculate all points inclosed in the loop.
	// Start at top/left corner and go to bottom/right
	bounds := GetBounds(loop)
	enclosedPoints := 0

	// Look for dots surrounded by unequal amount of pipes
	for x := bounds.Left; x <= bounds.Right; x++ {
		for y := bounds.Top; y <= bounds.Bottom; y++ {
			point := &Point{
				X:     x,
				Y:     y,
				Value: string(lines[y][x]),
			}

			if point.Value == "." && point.IsEnclosed(sides) {
				// fmt.Printf("Enclosed point at %v, %v\n", point.X, point.Y)
				enclosedPoints++
			}
		}
	}

	fmt.Printf("Number of enclosed ground tiles: %v\n", enclosedPoints)
}
