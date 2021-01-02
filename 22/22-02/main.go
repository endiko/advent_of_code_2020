package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type gameSetup struct {
	deck1 []int
	deck2 []int
}

func main() {
	startingDecks := readInput()

	game := gameSetup{
		deck1: parsePlayersDeck(startingDecks[0]),
		deck2: parsePlayersDeck(startingDecks[1]),
	}

	if game.play() {
		fmt.Println("Player 1 won", game.deck1)
		fmt.Println(calcScore(game.deck1))
	} else {
		fmt.Println("Player 2 won", game.deck2)
		fmt.Println(calcScore(game.deck2))
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
	result := make([]int, len(lines)-1)
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

func (game *gameSetup) play() bool {
	history := make(map[[16]byte]struct{})
	for len(game.deck1) > 0 && len(game.deck2) > 0 {
		hash := game.hash()
		if _, ok := history[hash]; ok {
			return true
		}
		history[hash] = struct{}{}

		card1 := game.deck1[0]
		card2 := game.deck2[0]
		game.deck1 = game.deck1[1:]
		game.deck2 = game.deck2[1:]
		var res bool
		if card1 <= len(game.deck1) && card2 <= len(game.deck2) {
			d1 := make([]int, card1)
			d2 := make([]int, card2)
			copy(d1, game.deck1[:card1])
			copy(d2, game.deck2[:card2])
			recursiveGame := gameSetup{deck1: d1, deck2: d2}
			res = recursiveGame.play()
		} else {
			res = card1 > card2
		}
		if res {
			game.deck1 = append(game.deck1, card1, card2)
		} else {
			game.deck2 = append(game.deck2, card2, card1)
		}
	}
	return len(game.deck1) > 0
}

func (game gameSetup) hash() [16]byte {
	bytes := []byte(fmt.Sprintf("%v", game))
	return md5.Sum(bytes)
}
