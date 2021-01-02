package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	startingCups := readInput()

	cups, first := parseInputPuzzle(startingCups)
	play(cups, first)
	theOne := cups[1]
	next := cups[theOne]
	//fmt.Println(cups)
	fmt.Println(next * theOne)
}

func readInput() []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), "")
}

func parseInputPuzzle(data []string) (map[int]int, int) {
	result := make(map[int]int, 1_000_000)
	currentCup := -1
	firstCup := -1
	for _, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(2)
		}
		if firstCup == -1 {
			firstCup = num
		}
		if currentCup != -1 {
			result[currentCup] = num
		}
		currentCup = num
		result[currentCup] = firstCup
	}

	for i := 10; i <= 1_000_000; i++ {
		if currentCup != -1 {
			result[currentCup] = i
		}
		currentCup = i
		result[currentCup] = firstCup
	}

	return result, firstCup
}

func play(cups map[int]int, first int) {
	current := first
	for i := 0; i < 10_000_000; i++ {
		pieceStart := cups[current]
		piece := extractCups(cups, current, 3)
		dest := current - 1
		for {
			_, ok := piece[dest]
			if !ok && dest != 0 {
				break
			}
			dest--
			if dest <= 0 {
				dest = findMax(cups, piece)
			}
		}
		insert(cups, dest, piece, pieceStart)

		current = cups[current]
		//if i%10_000 == 0 {
		//	fmt.Println(i)
		//}
	}
}

func extractCups(cups map[int]int, start, count int) (result map[int]struct{}) {
	result = make(map[int]struct{}, count)
	curr := start
	for i := 0; i < count; i++ {
		curr = cups[curr]
		result[curr] = struct{}{}
	}
	cups[start] = cups[curr]
	cups[curr] = -1
	return
}

func insert(cups map[int]int, dest int, piece map[int]struct{}, first int) {
	for key := range piece {
		if cups[key] == -1 {
			cups[key] = cups[dest]
		}
	}
	cups[dest] = first
}

func findMax(cups map[int]int, exclude map[int]struct{}) (result int) {
	result = -1
	for key := range cups {
		if _, ok := exclude[key]; ok {
			continue
		}
		if result < key {
			result = key
		}
	}
	return
}
