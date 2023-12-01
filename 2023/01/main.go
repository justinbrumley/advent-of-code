package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var digitTextValues = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

// getDigits returns a list of all digits found in the string
func getDigits(line []byte) []int {
	digits := make([]int, 0)

	for i, b := range line {
		num, err := strconv.Atoi(string(b))
		if err == nil {
			// Value is a number already
			digits = append(digits, num)
			continue
		}

		// Now check if this is a text sequence matching a digit
		substr := line[i:]

		for k, prefix := range digitTextValues {
			if bytes.HasPrefix(substr, prefix) {
				digits = append(digits, k+1)
			}
		}
	}

	return digits
}

func main() {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(data, []byte{'\n'})

	sum := 0

	for _, line := range lines {
		digits := getDigits(line)

		if len(digits) > 0 {
			value, err := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))

			if err != nil {
				log.Fatal(err)
			}

			sum += value
		}
	}

	fmt.Printf("Sum of calibration values: %v\n", sum)
}
