package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	Time     int
	Distance int
}

// GetWinOptions returns a list of ints representing
// how long the button needs to be held to win the race
func (r *Race) GetWinOptions() []int {
	winOptions := make([]int, 0)

	// Can never win holding it for 0ms, so start at 1
	for i := 1; i < r.Time-1; i++ {
		distance := i * (r.Time - i)
		if distance > r.Distance {
			winOptions = append(winOptions, i)
		}
	}

	return winOptions
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func parseRaces(lines [][]byte) []Race {
	re := regexp.MustCompile(`\d+`)

	times := re.FindAll(lines[0], -1)
	distances := re.FindAll(lines[1], -1)

	races := make([]Race, 0)

	// Part 1
	/*
		for i, time := range times {
			distance := distances[i]

			t, _ := strconv.Atoi(string(time))
			d, _ := strconv.Atoi(string(distance))

			races = append(races, Race{
				Time:     t,
				Distance: d,
			})
		}
	*/

	// Part 2
	t, _ := strconv.Atoi(string(bytes.Join(times, []byte(""))))
	d, _ := strconv.Atoi(string(bytes.Join(distances, []byte(""))))

	races = append(races, Race{
		Time:     t,
		Distance: d,
	})

	return races
}

func main() {
	lines := getInput()
	races := parseRaces(lines)

	winOptions := make([]int, 0)

	for _, race := range races {
		winOptions = append(winOptions, len(race.GetWinOptions()))
	}

	result := 0
	for i, opt := range winOptions {
		if i == 0 {
			result = opt
		} else {
			result *= opt
		}
	}

	fmt.Printf("Multiplied # of win options: %v\n", result)
}
