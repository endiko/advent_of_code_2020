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

type waypoint struct {
	x int
	y int
}

type coord struct {
	x  int
	y  int
	wp waypoint
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	currPos := coord{x: 0, y: 0, wp: waypoint{10, 1}}
	fmt.Printf("x: %d y: %d waypoint:%v\n", currPos.x, currPos.y, currPos.wp)
	for _, action := range dataSlice {
		num, err := strconv.Atoi(action[1:])
		if err != nil {
			os.Exit(1)
		}

		switch action[0] {
		case RIGHT, LEFT:
			currPos.turn(rune(action[0]), num)
		case FORWARD:
			currPos.move(num)
		default:
			currPos.moveWp(rune(action[0]), num)
		}
		fmt.Printf("%s%3d : (%3d; %3d) waypoint:%v\n", string(action[0]), num, currPos.x, currPos.y, currPos.wp)
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

func (pos *coord) move(step int) {
	pos.x += pos.wp.x * step
	pos.y += pos.wp.y * step
}

func (pos *coord) moveWp(direction rune, step int) {
	switch direction {
	case NORTH:
		pos.wp.y += step
	case SOUTH:
		pos.wp.y -= step
	case EAST:
		pos.wp.x += step
	case WEST:
		pos.wp.x -= step
	}
}

func (pos *coord) turn(direction rune, degree int) {
	var tbl = map[int][]int{
		90:  {0, 1},
		180: {-1, 0},
		270: {0, -1},
	}
	dir := 1
	if direction == RIGHT {
		dir = -1
	}
	x := pos.wp.x
	y := pos.wp.y
	cos := tbl[degree][0]
	sin := tbl[degree][1]
	// fmt.Println(dir, sin, cos)
	pos.wp.x = x*cos - dir*y*sin
	pos.wp.y = dir*x*sin + y*cos
}
