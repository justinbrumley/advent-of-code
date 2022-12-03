package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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
	lines, err := getInput()
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	prio := 0
	i := 0

	/* Part 2 solution */
	for i+3 <= len(lines) {
		group := lines[i : i+3]

		// Find the common character in all three lines
		var cVal byte

		for _, b := range group[0] {
			// For every byte in the first group,
			// check if that byte is in second and third
			if bytes.Contains(group[1], []byte{b}) && bytes.Contains(group[2], []byte{b}) {
				cVal = b
				break
			}
		}

		if cVal >= 'A' && cVal <= 'Z' {
			// 64 - 90
			prio += int(cVal) - 38 // Magic number
		} else if cVal >= 'a' && cVal <= 'z' {
			// 97 - 122
			prio += int(cVal) - 96 // Magic number
		}

		i += 3
	}

	/* Part 1 Solution */
	/*
		for _, line := range lines {
			perGroup := len(line) / 2

			c1 := bytes.Trim(line[0:perGroup], " \n")
			c2 := bytes.Trim(line[perGroup:], " \n")

			var cVal byte

			for _, b := range c1 {
				// Look for common letters
				// Assume first common letter found is the duplicate item
				if bytes.Contains(c2, []byte{b}) {
					cVal = b
					break
				}
			}

			if cVal >= 'A' && cVal <= 'Z' {
				// 64 - 90
				prio += int(cVal) - 38 // Magic number
			} else if cVal >= 'a' && cVal <= 'z' {
				// 97 - 122
				prio += int(cVal) - 96 // Magic number
			}
		}
	*/

	fmt.Printf("Sum of prio: %v\n", prio)
}
