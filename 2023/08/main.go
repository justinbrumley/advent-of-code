package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Location struct {
	Value string

	Left         string
	LeftLocation *Location

	Right         string
	RightLocation *Location
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func getLocations(lines [][]byte) []*Location {
	locations := make([]*Location, 0)

	re := regexp.MustCompile(`[A-Z1-9]{3}`)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := re.FindAll(line, -1)

		locations = append(locations, &Location{
			Value: string(parts[0]),
			Left:  string(parts[1]),
			Right: string(parts[2]),
		})
	}

	// Loop and attach references to prevent excess looping later
	for i, loc := range locations {
		locations[i].LeftLocation = findByValue(locations, loc.Left)
		locations[i].RightLocation = findByValue(locations, loc.Right)
	}

	return locations
}

func findByValue(locations []*Location, value string) *Location {
	for _, loc := range locations {
		if loc.Value == value {
			return loc
		}
	}

	return nil
}

func isFilled(values []int) bool {
	for _, val := range values {
		if val == 0 {
			return false
		}
	}

	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(values []int) int {
	result := 1

	for _, val := range values {
		result = result * val / GCD(result, val)
	}

	return result
}

func main() {
	lines := getInput()
	seq := bytes.Split(lines[0], []byte(""))
	locations := getLocations(lines[2:])

	steps := 0

	// Get list of start nodes
	nodes := make([]*Location, 0)

	for _, loc := range locations {
		if strings.HasSuffix(loc.Value, "A") {
			nodes = append(nodes, loc)
		}
	}

	paths := make([]int, len(nodes))

	fmt.Printf("Processing from %v starting nodes...\n", len(nodes))

	for !isFilled(paths) {
		for _, dir := range seq {
			steps++

			for i, node := range nodes {
				if string(dir) == "L" {
					nodes[i] = node.LeftLocation
				} else {
					nodes[i] = node.RightLocation
				}

				if strings.HasSuffix(nodes[i].Value, "Z") && paths[i] == 0 {
					paths[i] = steps
				}
			}
		}
	}

	fmt.Printf("Paths: %v\n", paths)

	// Get LCM of all paths
	lcm := LCM(paths)

	fmt.Printf("LCM: %v\n", lcm)
}
