package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tile struct {
	id          int
	border      []string
	flipped     bool
	orientation int
	data        []string
	fixed       bool
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n\n")

	tiles := make(map[int]tile, len(dataSlice))
	for _, block := range dataSlice {
		tile := parseTile(block)
		tiles[tile.id] = tile
	}

	allBorders := make(map[string]int)
	for _, tile := range tiles {
		for _, b := range tile.border {
			_, ok := allBorders[b]
			if ok {
				allBorders[b]++
			} else {
				allBorders[b] = 1
			}
		}
	}

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}
	currTile := findFirstCorner(tiles, allBorders)
	currTile.fixed = true
	matrix[0][0] = currTile.id

	for i := 1; i < 3; i++ {
		tile := findNeighbour(currTile.border[2], tiles, 0)
		tile.fixed = true
		matrix[i][0] = tile.id
		currTile = tile
	}
	fmt.Println(matrix)

	for i := 0; i < 3; i++ {
		currTile = tiles[matrix[i][0]]
		for j := 1; j < 3; j++ {
			tile := findNeighbour(currTile.border[1], tiles, 3)
			tile.fixed = true
			matrix[i][j] = tile.id
			currTile = tile
		}
	}

	fmt.Println(matrix)
}

func findNeighbour(border string, tiles map[int]tile, borderIdx int) (result tile) {
	freeBorders := collectFreeBorders(tiles)
	rev := reverse(border)
	if idx, ok := freeBorders[rev]; ok {
		result = tiles[idx]
	} else if idx, ok := freeBorders[border]; ok {
		result = tiles[idx]
		result.flip()
	}
	for rev != result.border[borderIdx] {
		result.rotate()
	}
	return
}

func collectFreeBorders(tiles map[int]tile) map[string]int {
	result := make(map[string]int)
	for _, tile := range tiles {
		if tile.fixed {
			continue
		}
		for _, b := range tile.border {
			result[b] = tile.id
		}
	}
	return result
}

func findFirstCorner(tiles map[int]tile, allBorders map[string]int) (result tile) {
	for _, t := range tiles {
		if t.countEdgeBorders(allBorders) == 2 {
			result = t
			break
		}
	}
	for {
		if checkEdge(result.border[0], allBorders) &&
			checkEdge(result.border[3], allBorders) {
			break
		}
		result.rotate()
	}
	return
}

func (t *tile) rotate() {
	t.orientation = (t.orientation + 1) % 4
	tmp := []string{t.border[3]}
	tmp = append(tmp, t.border[0:3]...)
	t.border = tmp
}

func (t *tile) flip() {
	t.flipped = !t.flipped
	tmp := make([]string, len(t.border))
	tmp[0] = reverse(t.border[2])
	tmp[1] = reverse(t.border[1])
	tmp[2] = reverse(t.border[0])
	tmp[3] = reverse(t.border[3])
}

func (t tile) countEdgeBorders(allBorders map[string]int) (cnt int) {
	for _, b := range t.border {
		if checkEdge(b, allBorders) {
			cnt++
		}
	}
	return
}

func checkEdge(b string, allBorders map[string]int) bool {
	if _, ok := allBorders[reverse(b)]; ok {
		return false
	}
	if cnt, ok := allBorders[b]; ok {
		return cnt == 1
	}
	return false
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func parseTile(raw string) tile {
	lines := strings.Split(raw, "\n")

	re := regexp.MustCompile("Tile (\\d+):")
	tileNum := re.FindStringSubmatch(lines[0])
	id, _ := strconv.Atoi(tileNum[1])

	res := tile{id: id, flipped: false, orientation: 0}

	right := ""
	left := ""
	bottom := ""
	for i := 0; i < 10; i++ {
		right += string(lines[i+1][9])
		left += string(lines[10-i][0])
		bottom += string(lines[10][9-i])
	}
	borders := []string{lines[1], right, bottom, left}
	res.border = borders

	data := make([]string, 8)
	for i := 0; i < len(data); i++ {
		data[i] = lines[i+2][1:9]
	}
	res.data = data

	return res
}
