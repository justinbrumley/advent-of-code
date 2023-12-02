package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var maxValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Pull represents the colors pulled in a game, separated by `;` in the input
type Pull struct {
	Colors map[string]int
}

// Game represents one line in the input
type Game struct {
	Number    int
	Power     int
	MaxColors map[string]int
	Pulls     []Pull
}

// IsValid checks if any pulls in the game exceed the max allowed values per color
func (g *Game) IsValid() bool {
	for _, pull := range g.Pulls {
		for color, value := range pull.Colors {
			if value > maxValues[color] {
				return false
			}
		}
	}

	return true
}

// getGameList parses games out of input line by line,
// and calculates power by checking max values of each color.
func getGameList(lines [][]byte) ([]Game, error) {
	games := make([]Game, 0)

	for _, line := range lines {
		split := bytes.Split(line, []byte{':'})

		if len(split) < 2 {
			continue
		}

		info, pullList := split[0], split[1]

		// Parse out the game info
		number, err := strconv.Atoi(string(info[5:]))

		if err != nil {
			return nil, err
		}

		game := Game{
			Number:    number,
			Pulls:     make([]Pull, 0),
			MaxColors: make(map[string]int),
		}

		pulls := bytes.Split(pullList, []byte{';'})

		for _, pull := range pulls {
			p := Pull{
				Colors: make(map[string]int),
			}

			// Use regexp to parse out color values with their numbers
			re := regexp.MustCompile(`(\d+) ([a-z]+),?`)
			matches := re.FindAllSubmatch(pull, -1)

			for _, match := range matches {
				val, err := strconv.Atoi(string(match[1]))
				if err != nil {
					return nil, err
				}

				color := string(match[2])
				p.Colors[color] += val

				// Calculate Power using highest color values for entire game
				if val > game.MaxColors[color] {
					game.MaxColors[color] = val
				}
			}

			game.Pulls = append(game.Pulls, p)
			game.Power = game.MaxColors["red"] * game.MaxColors["blue"] * game.MaxColors["green"]
		}

		games = append(games, game)
	}

	return games, nil
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
	games, err := getGameList(getInput())

	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	power := 0

	// Loop over all the games
	for _, game := range games {
		// Regardless of the rules, get the sum of Power scores
		power += game.Power

		// Check that color values never exceed the max...
		// If check succeeds, then add up the game numbers
		if game.IsValid() {
			sum += game.Number
		}
	}

	fmt.Printf("Sum of valid game numbers: %v\n", sum)
	fmt.Printf("Sum of powers: %v\n", power)
}
