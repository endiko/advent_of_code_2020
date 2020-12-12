package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// const vars
const (
	FLOOR    = '.'
	FREE     = 'L'
	OCCUPIED = '#'
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")
	grid := make([][]rune, len(dataSlice))

	for i, str := range dataSlice {
		grid[i] = make([]rune, len(str))
		for j, s := range str {
			grid[i][j] = s
		}
	}

	for {
		copyGrid := make([][]rune, len(dataSlice))
		for i := 0; i < len(grid); i++ {
			copyGrid[i] = make([]rune, len(grid[i]))
			for j := 0; j < len(grid[i]); j++ {
				copyGrid[i][j] = grid[i][j]

				switch grid[i][j] {
				case OCCUPIED:
					if countSeats(grid, i, j, OCCUPIED) >= 5 {
						copyGrid[i][j] = FREE
					}
				case FREE:
					if countSeats(grid, i, j, OCCUPIED) == 0 {
						copyGrid[i][j] = OCCUPIED
					}
				}
				// fmt.Printf("%s ", string(copyGrid[i][j]))
			}
			// fmt.Println()
		}
		// fmt.Println("===================")
		if compareGrids(grid, copyGrid) {
			break
		}
		grid = copyGrid
	}
	fmt.Println(countOccupiedSeats(grid))
}

func countSeats(data [][]rune, posX int, posY int, sym rune) int {
	increment := func(v int, r int) int { return v + r }
	decrement := func(v int, r int) int { return v - r }
	same := func(v int, r int) int { return v }

	count := 0
	count += calcLine(data, posX, posY, sym, decrement, same)
	count += calcLine(data, posX, posY, sym, decrement, decrement)
	count += calcLine(data, posX, posY, sym, same, decrement)
	count += calcLine(data, posX, posY, sym, increment, decrement)
	count += calcLine(data, posX, posY, sym, increment, same)
	count += calcLine(data, posX, posY, sym, increment, increment)
	count += calcLine(data, posX, posY, sym, same, increment)
	count += calcLine(data, posX, posY, sym, decrement, increment)
	return count
}

func calcLine(data [][]rune, startX int, startY int, sym rune, getX func(int, int) int, getY func(int, int) int) int {
	count := 0
	for r := 1; ; r++ {
		x := getX(startX, r)
		y := getY(startY, r)
		// fmt.Println(x, y)
		if x < 0 {
			break
		}
		if y < 0 {
			break
		}
		if x >= len(data) {
			break
		}
		if y >= len(data[x]) {
			break
		}
		if data[x][y] == sym {
			count++
			break
		} else if data[x][y] != FLOOR {
			break
		}
	}
	return count
}

func compareGrids(gridA [][]rune, gridB [][]rune) bool {
	for i := range gridA {
		for j := range gridA[i] {
			if gridA[i][j] != gridB[i][j] {
				return false
			}
		}
	}
	return true
}

func countOccupiedSeats(grid [][]rune) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == OCCUPIED {
				count++
			}
		}
	}
	return count
}
