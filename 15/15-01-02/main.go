package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type record struct {
	prev int
	last int
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), ",")
	history := make(map[int]record)
	var prev int

	for i, s := range dataSlice {
		num, err := strconv.Atoi(s)

		if err != nil {
			os.Exit(1)
		}

		history[num] = record{prev: i, last: i}
		prev = num
	}

	for i := len(history); i < 30000000; i++ { // part2 30000000
		if rec, ok := history[prev]; ok {
			prev = rec.last - rec.prev
			updateHistory(history, prev, i)
		} else {
			prev = 0
			updateHistory(history, prev, i)
		}
	}
	fmt.Println(prev)
}

func updateHistory(history map[int]record, prev int, idx int) {
	if rec, ok := history[prev]; ok {
		history[prev] = record{prev: rec.last, last: idx}
	} else {
		history[prev] = record{prev: idx, last: idx}
	}
}
