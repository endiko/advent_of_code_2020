package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	startingDecks := readInput()

	deckPlayer1 := parsePlayersDeck(startingDecks[0])
	deckPlayer2 := parsePlayersDeck(startingDecks[1])

	for len(deckPlayer1) > 0 && len(deckPlayer2) > 0 {
		card1 := deckPlayer1[0]
		card2 := deckPlayer2[0]
		deckPlayer1 = deckPlayer1[1:]
		deckPlayer2 = deckPlayer2[1:]
		if card1 > card2 {
			deckPlayer1 = append(deckPlayer1, card1, card2)
		} else {
			deckPlayer2 = append(deckPlayer2, card2, card1)
		}
	}
	if len(deckPlayer1) > 0 {
		fmt.Println(calcScore(deckPlayer1))
	} else {
		fmt.Println(calcScore(deckPlayer2))
	}
}

func readInput() []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), "\n\n")
}

func parsePlayersDeck(s string) []int {
	lines := strings.Split(s, "\n")
	result := make([]int, len(lines) - 1)
	for i, line := range lines[1:] {
		card, err := strconv.Atoi(line)
		if err != nil {
			os.Exit(2)
		}
		result[i] = card
	}
	return result
}

func calcScore(deck []int) (result int) {
	for idx, card := range deck {
		result += card * (len(deck) - idx)
	}
	return
}
