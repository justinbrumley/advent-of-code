package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

// IsEmpty returns true if all values are set to 0
func IsEmpty(values []int) bool {
	for _, val := range values {
		if val != 0 {
			return false
		}
	}

	return true
}

func prepend(slice []int, value int) []int {
	s := make([]int, 1)
	s[0] = value
	return append(s, slice...)
}

// GetPrevValue determines the prev value that would come before the sequence (slice)
func GetPrevValue(values []int) int {
	sequences := make([][]int, 1)
	sequences[0] = values

	// Start by building up sequences
	for {
		prevSeq := sequences[len(sequences)-1]
		if len(prevSeq) == 0 {
			break
		}

		seq := make([]int, len(prevSeq)-1)

		for i, val := range prevSeq {
			// Skip the last one
			if i == len(prevSeq)-1 {
				break
			}

			diff := prevSeq[i+1] - val
			seq[i] = diff
		}

		sequences = append(sequences, seq)

		if IsEmpty(seq) {
			break
		}
	}

	// Now reverse back up to append one extra value to each sequence
	for i := len(sequences) - 1; i >= 0; i-- {
		if i == len(sequences)-1 {
			// Last row always gets a zero prepended
			sequences[i] = prepend(sequences[i], 0)
			continue
		}

		// All other rows get shiny new value
		currSeq := sequences[i]
		nextSeq := sequences[i+1]

		firstVal := currSeq[0]
		diff := nextSeq[0]
		sequences[i] = prepend(sequences[i], firstVal-diff)
	}

	// Now just return the first value of the first sequence
	return sequences[0][0]
}

// GetNextValue determines the next value that would come at the end of the sequence (slice)
func GetNextValue(values []int) int {
	sequences := make([][]int, 1)
	sequences[0] = values

	// Start by building up sequences
	for {
		prevSeq := sequences[len(sequences)-1]
		if len(prevSeq) == 0 {
			break
		}

		seq := make([]int, len(prevSeq)-1)

		for i, val := range prevSeq {
			// Skip the last one
			if i == len(prevSeq)-1 {
				break
			}

			diff := prevSeq[i+1] - val
			seq[i] = diff
		}

		sequences = append(sequences, seq)

		if IsEmpty(seq) {
			break
		}
	}

	// Now reverse back up to append one extra value to each sequence
	for i := len(sequences) - 1; i >= 0; i-- {
		if i == len(sequences)-1 {
			// Last row always gets a zero appended
			sequences[i] = append(sequences[i], 0)
			continue
		}

		// All other rows get shiny new value
		currSeq := sequences[i]
		nextSeq := sequences[i+1]

		lastVal := currSeq[len(currSeq)-1]
		diff := nextSeq[len(nextSeq)-1]
		sequences[i] = append(sequences[i], lastVal+diff)
	}

	// Now just return the last value of the first sequence
	return sequences[0][len(sequences[0])-1]
}

func main() {
	lines := getInput()

	matrix := make([][]int, 0)
	re := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nums := re.FindAll(line, -1)
		row := make([]int, len(nums))

		for i, num := range nums {
			val, err := strconv.Atoi(string(num))
			if err != nil {
				log.Fatal(err)
			}

			row[i] = val
		}

		matrix = append(matrix, row)
	}

	firstSum := 0
	lastSum := 0

	// GetPrevValue and GetNextValue could easily be squashed into one,
	// but easier to just copy and change what I need
	for _, row := range matrix {
		val := GetPrevValue(row)
		firstSum += val

		val = GetNextValue(row)
		lastSum += val
	}

	fmt.Printf("Sum of new first values in sequences: %v\n", firstSum)
	fmt.Printf("Sum of new final values in sequences: %v\n", lastSum)
}
