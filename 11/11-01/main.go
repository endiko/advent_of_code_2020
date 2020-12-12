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
					if countSeats(grid, i, j, OCCUPIED) >= 4 {
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
		// fmt.Println("========================")
		if compareGrids(grid, copyGrid) {
			break
		}
		grid = copyGrid
	}
	fmt.Println(countOccupiedSeats(grid))
}

func countSeats(data [][]rune, posX int, posY int, sym rune) int {
	startX := posX - 1
	endX := posX + 1
	if posX == 0 {
		startX = 0
	}
	if posX == len(data)-1 {
		endX = posX
	}

	startY := posY - 1
	endY := posY + 1
	if posY == 0 {
		startY = 0
	}
	if posY == len(data[posX])-1 {
		endY = posY
	}

	count := 0
	for k := startX; k <= endX; k++ {
		for m := startY; m <= endY; m++ {
			if k == posX && m == posY {
				continue
			}
			if data[k][m] == sym {
				count++
			}
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
