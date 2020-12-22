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
	id     int
	border []string
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n\n")

	tiles := make([]tile, len(dataSlice))
	for i, block := range dataSlice {
		tiles[i] = parseTile(block)
		fmt.Println(tiles[i])
	}

	score := make(map[string]int)
	for _, tile := range tiles {
		for _, b := range tile.border {
			_, ok := score[b]
			if ok {
				score[b]++
			} else {
				score[b] = 1
			}
		}
	}

	product := 1
	for _, tile := range tiles {
		cnt := 0
		for _, b := range tile.border {
			if checkEdge(b, score) {
				cnt++
			}
		}
		if cnt == 2 {
			product *= tile.id
		}
	}

	fmt.Println(product)
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

func parseTile(raw string) tile {
	re := regexp.MustCompile("Tile (\\d+):")
	strs := strings.Split(raw, "\n")
	tileNum := re.FindStringSubmatch(strs[0])
	id, _ := strconv.Atoi(tileNum[1])
	res := tile{id: id}

	borders := make([]string, 4)
	borders[0] = strs[1]
	right := ""
	left := ""
	bottom := ""
	for i := 0; i < 10; i++ {
		right += string(strs[i+1][9])
		left += string(strs[10-i][0])
		bottom += string(strs[10][9-i])
	}
	borders[1] = right
	borders[2] = bottom
	borders[3] = left
	res.border = borders
	return res
}

func reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}