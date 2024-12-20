package utils

import (
	"bufio"
	"os"
)

// Read input file and return contents.
func GetInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	return lines, err
}

// Read input file and return contents.
func GetInputBytes(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	err = scanner.Err()
	return lines, err
}
