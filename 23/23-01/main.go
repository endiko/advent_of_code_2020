package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cup struct {
	value int
	next *Cup
}

func main() {
	startingCups := readInput()

	currentCup := parseInputPuzzle(startingCups)
	play(currentCup)
	fmt.Println(printAllCups(findCup(currentCup, 1)))
}

func readInput() []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), "")
}

func parseInputPuzzle(data []string) *Cup {
	var currentCup *Cup
	var firstCup *Cup
	for _, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(2)
		}
		tmpCup := &Cup{value: num}
		if firstCup == nil {
			firstCup = tmpCup
		}
		if currentCup != nil {
			currentCup.next = tmpCup
		}
		currentCup = tmpCup
		currentCup.next = firstCup
	}

	return firstCup
}

func printAllCups(c *Cup) (result string) {
	currentCup := c.next
	for c != currentCup {
		result += strconv.Itoa(currentCup.value)
		currentCup = currentCup.next
	}
	return
}

func findCup(c *Cup, num int) *Cup {
	current := c

	for {
		if current.value == num {
			return current
		}
		current = current.next
		if current == c {
			break
		}
	}

	return nil
}

func play(c *Cup) {
	current := c
	for i:=0; i< 100; i++ {
		piece := extractCups(current, 3)
		target := current.value - 1
		dest := findCup(current, target)
		for dest == nil {
			target--
			if target <= 0 {
				target = findMax(current)
			}
			dest = findCup(current, target)
		}
		insert(dest, piece)
		current = current.next
	}
}

func extractCups(start *Cup, count int) (result *Cup) {
	result = start.next
	curr := start
	for i:=0; i<count;i++ {
		curr = curr.next
	}
	start.next = curr.next
	curr.next = nil

	return
}

func insert(dest *Cup, piece *Cup)  {
	curr :=  piece
	for curr.next != nil  {
		curr = curr.next
	}

	curr.next = dest.next
	dest.next = piece
}

func findMax(c *Cup) (result int) {
	result = c.value
	curr := c.next
	for curr != c {
		if result < curr.value {
			result = curr.value
		}
		curr = curr.next
	}
	return
}