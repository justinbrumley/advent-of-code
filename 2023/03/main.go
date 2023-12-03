package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type PartNumber struct {
	Line  int
	Start int
	End   int
	Value int
}

type Grid struct {
	Values          [][]byte
	PartNumbers     []PartNumber
	PartNumberTotal int
	GearRatio       int
}

func (g *Grid) GetPartNumberTotal() int {
	sum := 0

	for _, number := range g.PartNumbers {
		sum += number.Value
	}

	return sum
}

// AddPartNumber will add the given number to the list of part numbers,
// if it's not already in there. Returns true if number was stored.
func (g *Grid) AddPartNumber(line, start, end, value int) bool {
	for _, num := range g.PartNumbers {
		if num.Line == line && num.Start == start && num.End == end {
			return false
		}
	}

	g.PartNumbers = append(g.PartNumbers, PartNumber{
		Line:  line,
		Start: start,
		End:   end,
		Value: value,
	})

	return true
}

// GetNumberAt gets the whole number found at the given position
func (g *Grid) IsSymbolAt(x, y int) bool {
	re := regexp.MustCompile(`^[^a-zA-Z0-9\.\s]$`)
	char := g.Values[y][x : x+1]
	return re.Match(char)
}

// GetNumberAt gets the whole number found at the given position
func (g *Grid) GetNumberAt(x, y int) (int, int, int) {
	line := g.Values[y]

	// First check if target coordinate is a number
	re := regexp.MustCompile(`^\d+$`)
	char := line[x : x+1]

	if !re.Match(char) {
		return 0, -1, -1
	}

	// We have a hit, so expand out to build the whole number
	left := 1
	right := 1

	if x > 0 {
		for {
			char = line[x-left : x]

			if !re.Match(char) {
				// Put it back
				left -= 1
				char = line[x-left : x]
				break
			}

			if x-left == 0 {
				break
			}

			left += 1
		}
	}

	if x < len(line)-1 {
		for {
			char = line[x-left : x+right]

			if !re.Match(char) {
				// Put it back
				right -= 1
				char = line[x-left : x+right]
				break
			}

			if x+right > len(line)-1 {
				break
			}

			right += 1
		}
	}

	val, err := strconv.Atoi(string(char))

	if err != nil {
		fmt.Printf("Error parsing number: %v at %v,%v\n", char, x, y)
		log.Fatal(err)
	}

	return val, x - left, x + right
}

// GetNeighbors returns any numbers neighboring the given point
func (g *Grid) CalculatePartNumbersAround(x, y int, isGear bool) {
	nums := make([]int, 0)

	for _, j := range []int{y - 1, y, y + 1} {
		for _, i := range []int{x - 1, x, x + 1} {
			value, start, end := g.GetNumberAt(i, j)

			if value > 0 {
				if g.AddPartNumber(j, start, end, value) {
					nums = append(nums, value)
				}
			}
		}
	}

	if isGear && len(nums) > 1 {
		// Multiple adjacent numbers together and add to gear ratio total
		gearRatio := nums[0]

		for _, val := range nums[1:] {
			gearRatio *= val
		}

		g.GearRatio += gearRatio
	}
}

// CalculatePartNumbers will find all part numbers and store the total
// Will also check for any gears and set the overall GearRatio
func (g *Grid) CalculatePartNumbers() {
	g.PartNumbers = make([]PartNumber, 0)

	// Loop over each point in the grid, and get any neighboring numbers
	for y, line := range g.Values {
		for x, _ := range line {
			if !g.IsSymbolAt(x, y) {
				continue
			}

			g.CalculatePartNumbersAround(x, y, line[x] == '*')
		}
	}

	g.PartNumberTotal = g.GetPartNumberTotal()
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func main() {
	grid := &Grid{
		Values: getInput(),
	}

	grid.CalculatePartNumbers()

	fmt.Printf("Sum of part numbers:\t%v\n", grid.PartNumberTotal)
	fmt.Printf("Sum of gear ratios:\t%v\n", grid.GearRatio)
}
