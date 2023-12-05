package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Conversion struct {
	Destination int
	Source      int
	Range       int
}

type ConversionMap struct {
	Name        []byte // For debugging
	Conversions []Conversion
}

// Convert pipes the number through the current conversions list to try to find a range it fits into.
// If no range is found, the input number is returned.
func (m *ConversionMap) Convert(val int) int {
	for _, c := range m.Conversions {
		if val >= c.Source && val < c.Source+c.Range {
			// Within range! Calculate the diff and add it to destination.
			delta := val - c.Source
			return c.Destination + delta
		}
	}

	// No match found
	return val
}

func (m *ConversionMap) Print() {
	fmt.Printf("%v\n", string(m.Name))

	for _, c := range m.Conversions {
		fmt.Printf("%v %v %v\n", c.Destination, c.Source, c.Range)
	}

	fmt.Println()
}

func NewConversionMap() ConversionMap {
	return ConversionMap{
		Conversions: make([]Conversion, 0),
	}
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func parseConversionMaps(lines [][]byte) []ConversionMap {
	maps := make([]ConversionMap, 0)
	tmp := NewConversionMap()

	for _, line := range lines {
		if line == nil || len(line) == 0 {
			maps = append(maps, tmp)
			tmp = NewConversionMap()
			continue
		}

		if tmp.Name == nil {
			tmp.Name = line
			continue
		}

		re := regexp.MustCompile(`\d+`)
		m := Conversion{}
		for i, match := range re.FindAll(line, -1) {
			val, err := strconv.Atoi(string(match))
			if err != nil {
				log.Fatal(err)
			}

			if i == 0 {
				// Destination
				m.Destination = val
			} else if i == 1 {
				// Source
				m.Source = val
			} else {
				// Range
				m.Range = val
			}
		}

		tmp.Conversions = append(tmp.Conversions, m)
	}

	return maps
}

func Min(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	low := nums[0]
	for _, val := range nums {
		if val < low {
			low = val
		}
	}

	return low
}

func main() {
	lines := getInput()

	ts := time.Now().Unix()

	start := -1
	destination := 0
	maps := parseConversionMaps(lines[2:])

	for _, val := range bytes.Split(lines[0], []byte(" ")) {
		if seed, err := strconv.Atoi(string(val)); err == nil {
			if start == -1 {
				start = seed
			} else {
				for i := start; i < start+seed; i++ {
					result := i

					for _, m := range maps {
						// Check if current value falls with range of any conversion sets
						result = m.Convert(result)
					}

					// Final Destination™
					if destination == 0 || result < destination {
						destination = result
					}
				}

				start = -1
			}
		}
	}

	fmt.Printf("Lowest Destination: %v\n", destination)
	fmt.Printf("Completed after %v seconds\n", time.Now().Unix()-ts)
}
