package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

const (
	E  = "e"
	SE = "se"
	SW = "sw"
	W  = "w"
	NW = "nw"
	NE = "ne"
)

var moves = map[string]Point{
	E:  {x: 2, y: 0},
	SE: {x: 1, y: -1},
	SW: {x: -1, y: -1},
	W:  {x: -2, y: 0},
	NW: {x: -1, y: 1},
	NE: {x: 1, y: 1},
}

func main() {
	addressList := readInput("\n")

	tiles := make(map[[16]byte]bool)
	for _, s := range addressList {
		route := parseTileRoute(s)
		simplifyRoute(route)
		hash := hash(route)
		if v, ok := tiles[hash]; ok {
			tiles[hash] = !v
		} else {
			tiles[hash] = true
		}
	}

	count := 0
	for _, v := range tiles {
		if v {
			count++
		}
	}

	fmt.Println(count)
}

func readInput(sep string) []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), sep)
}

func prepareRoute() (result map[string]int) {
	result = make(map[string]int)
	for s, _ := range moves {
		result[s] = 0
	}
	return
}

func parseTileRoute(data string) (result map[string]int) {
	result = prepareRoute()

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

func simplifyRoute(route map[string]int) {
	for {
		success := simplify(route, E, NW, NE)
		success = success || simplify(route, NE, W, NW)
		success = success || simplify(route, NW, SW, W)
		success = success || simplify(route, W, SE, SW)
		success = success || simplify(route, SW, E, SE)
		success = success || simplify(route, SE, NE, E)
		success = success || simplify(route, E, W, "")
		success = success || simplify(route, SW, NE, "")
		success = success || simplify(route, SE, NW, "")
		if !success {
			break
		}
	}
}

func simplify(route map[string]int, first, second, result string) bool {
	min := route[first]
	if min > route[second] {
		min = route[second]
	}
	if min == 0 {
		return false
	}
	route[first] -= min
	route[second] -= min
	if _, ok := route[result]; ok {
		route[result] += min
	}
	return true
}

func hash(route map[string]int) [16]byte {
	bytes := []byte(fmt.Sprintf("%v", route))
	return md5.Sum(bytes)
}
