package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// const vars
const (
	NORTH   = 'N'
	SOUTH   = 'S'
	EAST    = 'E'
	WEST    = 'W'
	LEFT    = 'L'
	RIGHT   = 'R'
	FORWARD = 'F'
)

type coord struct {
	x         int
	y         int
	direction rune
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	currPos := coord{x: 0, y: 0, direction: EAST}
	fmt.Printf("x: %d y: %d direction:%s\n", currPos.x, currPos.y, string(currPos.direction))
	for _, action := range dataSlice {
		num, err := strconv.Atoi(action[1:])
		if err != nil {
			os.Exit(1)
		}

		switch action[0] {
		case RIGHT, LEFT:
			currPos.turn(rune(action[0]), num)
		case FORWARD:
			currPos.move(currPos.direction, num)
		default:
			currPos.move(rune(action[0]), num)
		}
		fmt.Printf("%s%3d : (%3d; %3d) dir:%s\n", string(action[0]), num, currPos.x, currPos.y, string(currPos.direction))
	}

	fmt.Println(calcManhattanDistance(currPos))
}

func calcManhattanDistance(pos coord) int {
	result := 0
	if pos.x < 0 {
		result -= pos.x
	} else {
		result += pos.x
	}
	if pos.y < 0 {
		result -= pos.y
	} else {
		result += pos.y
	}
	return result
}

func (pos *coord) move(direction rune, step int) {
	switch direction {
	case NORTH:
		pos.y -= step
	case SOUTH:
		pos.y += step
	case EAST:
		pos.x += step
	case WEST:
		pos.x -= step
	}
}

func (pos *coord) turn(direction rune, degree int) {
	var rules = map[rune]map[rune]rune{
		NORTH: map[rune]rune{RIGHT: EAST, LEFT: WEST},
		EAST:  map[rune]rune{RIGHT: SOUTH, LEFT: NORTH},
		SOUTH: map[rune]rune{RIGHT: WEST, LEFT: EAST},
		WEST:  map[rune]rune{RIGHT: NORTH, LEFT: SOUTH},
	}
	for click := 0; click < degree; click += 90 {
		pos.direction, _ = rules[pos.direction][direction]
	}
}
