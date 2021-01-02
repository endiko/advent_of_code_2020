package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	E  = "e"
	SE = "se"
	SW = "sw"
	W  = "w"
	NW = "nw"
	NE = "ne"
)

var moves = []string{E, SE, SW, W, NW, NE}

type Coord map[string]int

type Tile struct {
	pos  Coord
	flip bool
}

func main() {
	addressList := readInput("\n")

	tiles := make(map[[16]byte]*Tile)
	for _, s := range addressList {
		coord := parseTileCoord(s)
		coord.simplifyRoute()
		hash := coord.hash()

		if v, ok := tiles[hash]; ok {
			v.flip = !v.flip
		} else {
			tiles[hash] = &Tile{pos: coord, flip: true}
		}
	}

	fmt.Printf("Day %d: %d\n", 0, countFlipTiles(tiles))
	for i := 0; i < 100; i++ {
		tiles = liveADay(tiles)
		fmt.Printf("Day %d: %d\n", i+1, countFlipTiles(tiles))
	}

}

func countFlipTiles(tiles map[[16]byte]*Tile) (count int) {
	count = 0
	for _, v := range tiles {
		if v.flip {
			count++
		}
	}
	return
}

func liveADay(tiles map[[16]byte]*Tile) (tomorrow map[[16]byte]*Tile) {
	checkSet := make(map[[16]byte]Coord)
	for _, tile := range tiles {
		if !tile.flip {
			continue
		}
		checkSet[tile.pos.hash()] = tile.pos

		neighbours := listNeighbours(tile.pos)
		for _, pos := range neighbours {
			checkSet[pos.hash()] = pos
		}
	}

	// Переписал аккуратнее проверку полученного количества и всё
	tomorrow = make(map[[16]byte]*Tile)
	for _, pos := range checkSet {
		count := countFlipNeighbours(tiles, pos)
		if count == 0 || count > 2 {
			continue
		}
		if count == 2 {
			tomorrow[pos.hash()] = &Tile{pos: pos, flip: true}
			continue
		}
		if tile, ok := tiles[pos.hash()]; ok && tile.flip {
			tomorrow[pos.hash()] = &Tile{pos: pos, flip: true}
		}
	}
	return
}

func countFlipNeighbours(tiles map[[16]byte]*Tile, pos Coord) (result int) {
	neighbours := listNeighbours(pos)
	for _, pos := range neighbours {
		if v, ok := tiles[pos.hash()]; ok && v.flip {
			result++
		}
	}
	return
}

func listNeighbours(pos Coord) (result []Coord) {
	result = make([]Coord, 6)
	result[0] = incCoord(pos, E)
	result[1] = incCoord(pos, NE)
	result[2] = incCoord(pos, NW)
	result[3] = incCoord(pos, W)
	result[4] = incCoord(pos, SW)
	result[5] = incCoord(pos, SE)

	for _, coord := range result {
		coord.simplifyRoute()
	}
	return
}

func incCoord(pos Coord, direction string) (result Coord) {
	result = make(Coord, len(pos))

	for k, v := range pos {
		result[k] = v
	}

	result[direction]++

	return
}

func readInput(sep string) []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), sep)
}

func prepareCoord() (result Coord) {
	result = make(Coord)
	for _, s := range moves {
		result[s] = 0
	}
	return
}

func parseTileCoord(data string) (result Coord) {
	result = prepareCoord()

	curr := ""
	for _, r := range data {
		curr += string(r)
		switch curr {
		case E:
			result[E]++
		case W:
			result[W]++
		case SE:
			result[SE]++
		case SW:
			result[SW]++
		case NW:
			result[NW]++
		case NE:
			result[NE]++
		default:
			continue
		}
		curr = ""
	}

	return
}

func (c Coord) simplifyRoute() {
	for {
		success := c.simplify(E, NW, NE)
		success = success || c.simplify(NE, W, NW)
		success = success || c.simplify(NW, SW, W)
		success = success || c.simplify(W, SE, SW)
		success = success || c.simplify(SW, E, SE)
		success = success || c.simplify(SE, NE, E)
		success = success || c.simplify(E, W, "")
		success = success || c.simplify(SW, NE, "")
		success = success || c.simplify(SE, NW, "")
		if !success {
			break
		}
	}
}

func (c Coord) simplify(first, second, result string) bool {
	min := c[first]
	if min > c[second] {
		min = c[second]
	}
	if min == 0 {
		return false
	}
	c[first] -= min
	c[second] -= min
	if _, ok := c[result]; ok {
		c[result] += min
	}
	return true
}

func (c Coord) hash() [16]byte {
	bytes := []byte(fmt.Sprintf("%v", c))
	return md5.Sum(bytes)
}
