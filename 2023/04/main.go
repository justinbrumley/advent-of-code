package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Card struct {
	Index          int
	WinningNumbers []int
	Numbers        []int
	Score          int
	Matches        int
}

type Cards []Card

func (cards Cards) GetByIndex(index int) *Card {
	for _, card := range cards {
		if card.Index == index {
			return &card
		}
	}

	return nil
}

func (cards *Cards) Add(card Card) {
	*cards = append(*cards, card)
}

// AddCopy adds a copy of a scratchcard by index
func (cards *Cards) AddCopy(index int) {
	card := cards.GetByIndex(index)

	if card != nil {
		cards.Add(card.Copy())
	}
}

func (c *Card) Copy() Card {
	return Card{
		Index:          c.Index,
		WinningNumbers: c.WinningNumbers,
		Numbers:        c.Numbers,
		Score:          c.Score,
		Matches:        c.Matches,
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

// parseInputLine returns the list of winning numbers and the list of user numbers on the line
func parseInputLine(index int, line []byte) Card {
	parts := bytes.Split(
		line[bytes.Index(line, []byte(":")):],
		[]byte{'|'},
	)

	re := regexp.MustCompile(`\d+`)

	winningNumbers := make([]int, 0)

	for _, val := range re.FindAll(parts[0], -1) {
		num, err := strconv.Atoi(string(val))
		if err != nil {
			log.Fatal(err)
		}

		winningNumbers = append(winningNumbers, num)
	}

	numbers := make([]int, 0)

	for _, val := range re.FindAll(parts[1], -1) {
		num, err := strconv.Atoi(string(val))
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
	}

	card := Card{
		Index:          index + 1, // Starts at 1 in input file
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
	}

	// Calculate score and number of matches for the card
	score := 0
	matches := 0

	for _, winningNum := range card.WinningNumbers {
		for _, num := range card.Numbers {
			if winningNum == num {
				matches += 1

				// Score!
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}

	card.Score = score
	card.Matches = matches

	return card
}

func main() {
	lines := getInput()

	cards := make(Cards, 0)

	// Parse and store initial list of cards
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		cards.Add(parseInputLine(i, line))
	}

	// Part 1: Get total score of cards
	points := 0
	for _, card := range cards {
		points += card.Score
	}

	fmt.Printf("Total score of all winning numbers: %v\n", points)

	// Part 2: Loop over and reward card copies based on matches
	for i := 0; i < len(cards); i++ {
		card := cards[i]

		if card.Matches > 0 {
			for i := card.Index + 1; i < card.Index+1+card.Matches; i++ {
				cards.AddCopy(i)
			}
		}
	}

	fmt.Printf("Total number of cards after rewards: %v\n", len(cards))
}
