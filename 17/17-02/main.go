package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// OFFSET is offset
const OFFSET = 6

type tile struct {
	active bool
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n")
	grid := generateTiles(20, 20, 13, 20)

	for i, str := range dataSlice {
		for j, r := range str {
			grid[i+OFFSET][j+OFFSET][OFFSET][OFFSET].active = r == '#'
		}
	}

	for c := 0; c < OFFSET; c++ {
		tempGrid := generateTiles(20, 20, 13, 20)

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				for k := 0; k < len(grid[i][j]); k++ {
					for m := 0; m < len(grid[i][j][k]); m++ {
						tempGrid[i][j][k][m].active = checkIsActive(grid, i, j, k, m)
					}
				}
			}
		}

		grid = tempGrid
	}
	fmt.Println(countActive(grid))
}

func countActive(grid [][][][]tile) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for k := 0; k < len(grid[i][j]); k++ {
				for m := 0; m < len(grid[i][j][k]); m++ {
					if grid[i][j][k][m].active {
						counter++
					}
				}
			}
		}
	}

	return counter
}

func generateTiles(x int, y int, z int, w int) [][][][]tile {
	var tiles [][][][]tile

	tiles = make([][][][]tile, x)

	for i := range tiles {
		tiles[i] = make([][][]tile, y)
		for j := range tiles[i] {
			tiles[i][j] = make([][]tile, z)
			for k := range tiles[i][j] {
				tiles[i][j][k] = make([]tile, w)
			}
		}
	}

	return tiles
}

func checkIsActive(data [][][][]tile, x int, y int, z int, w int) bool {
	var count = 0

	startX := x - 1
	if x == 0 {
		startX = 0
	}
	startY := y - 1
	if y == 0 {
		startY = 0
	}
	startZ := z - 1
	if z == 0 {
		startZ = 0
	}
	startW := w - 1
	if w == 0 {
		startW = 0
	}

	endX := x + 1
	if x == len(data)-1 {
		endX = x
	}
	endY := y + 1
	if y == len(data[x])-1 {
		endY = y
	}
	endZ := z + 1
	if z == len(data[x][y])-1 {
		endZ = z
	}

	endW := w + 1
	if w == len(data[x][y][z])-1 {
		endW = w
	}

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			for k := startZ; k <= endZ; k++ {
				for m := startW; m <= endW; m++ {
					if i == x && j == y && k == z && m == w {
						continue
					}
					if data[i][j][k][m].active {
						count++
					}
				}
			}
		}
	}

	if data[x][y][z][w].active {
		return count == 2 || count == 3
	}
	return count == 3
}
