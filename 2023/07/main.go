package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var CardList = [][]byte{
	[]byte("J"),
	[]byte("2"),
	[]byte("3"),
	[]byte("4"),
	[]byte("5"),
	[]byte("6"),
	[]byte("7"),
	[]byte("8"),
	[]byte("9"),
	[]byte("T"),
	[]byte("Q"),
	[]byte("K"),
	[]byte("A"),
}

func GetCardIndex(card []byte) int {
	for i, c := range CardList {
		if string(c) == string(card) {
			return i
		}
	}

	return -1
}

type Hand struct {
	Cards   [][]byte
	CardMap map[string]int
	Bet     int
}

func (h *Hand) IsStronger(other Hand) bool {
	str := h.GetStrength()
	otherStr := other.GetStrength()

	if str != otherStr {
		return str > otherStr
	}

	// Compare cards left to right
	for i, card := range h.Cards {
		otherCard := other.Cards[i]

		index := GetCardIndex(card)
		otherIndex := GetCardIndex(otherCard)

		if index != otherIndex {
			return index > otherIndex
		}
	}

	return false
}

func (h *Hand) IsFiveOfAKind() bool {
	for _, val := range h.CardMap {
		if val == 5 {
			return true
		}
	}

	return false
}

func (h *Hand) IsFourOfAKind() bool {
	for _, val := range h.CardMap {
		if val == 4 {
			return true
		}
	}

	return false
}

func (h *Hand) IsFullHouse() bool {
	hasPair := false
	hasThreeOfAKind := false

	for _, val := range h.CardMap {
		if val >= 3 && !hasThreeOfAKind {
			hasThreeOfAKind = true
		} else if val >= 2 {
			hasPair = true
		}
	}

	return hasPair && hasThreeOfAKind
}

func (h *Hand) IsThreeOfAKind() bool {
	for _, val := range h.CardMap {
		if val == 3 {
			return true
		}
	}

	return false
}

func (h *Hand) IsTwoPair() bool {
	if len(h.CardMap) > 3 || len(h.CardMap) < 2 {
		return false
	}

	pairs := 0
	for _, val := range h.CardMap {
		if val == 2 {
			pairs++
		}
	}

	return pairs > 1
}

func (h *Hand) IsOnePair() bool {
	pairs := 0
	for _, val := range h.CardMap {
		if val == 2 {
			pairs++
		}
	}

	return pairs > 0
}

func (h *Hand) GetStrength() int {
	if h.IsFiveOfAKind() {
		return 6
	}

	if h.IsFourOfAKind() {
		return 5
	}

	if h.IsFullHouse() {
		return 4
	}

	if h.IsThreeOfAKind() {
		return 3
	}

	if h.IsTwoPair() {
		return 2
	}

	if h.IsOnePair() {
		return 1
	}

	// High Card
	return 0
}

// getInput parses the input file and returns input split by newlines
func getInput() [][]byte {
	data, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	return bytes.Split(data, []byte{'\n'})
}

func getHands(lines [][]byte) []Hand {
	hands := make([]Hand, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := bytes.Split(line, []byte(" "))
		bet, _ := strconv.Atoi(string(parts[1]))

		hand := Hand{
			Cards: bytes.Split(parts[0], []byte("")),
			Bet:   bet,
		}

		maxCard := ""
		max := 0

		hand.CardMap = make(map[string]int)
		jokers := 0
		for _, card := range hand.Cards {
			if string(card) == "J" {
				jokers++
				continue
			}

			hand.CardMap[string(card)]++

			if hand.CardMap[string(card)] > max && string(card) != "J" {
				max = hand.CardMap[string(card)]
				maxCard = string(card)
			}
		}

		// Add one to each card in the map for each joker
		if max > 0 && jokers > 0 {
			// Duplicate max card with jokers
			hand.CardMap[maxCard] += jokers
		} else if jokers > 0 {
			hand.CardMap["J"] = jokers
		}

		hands = append(hands, hand)
	}

	return hands
}

func main() {
	lines := getInput()
	hands := getHands(lines)

	// Sort by hand strength
	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].IsStronger(hands[j])
	})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.Bet
	}

	fmt.Printf("Winnings: %v\n", winnings)
}
