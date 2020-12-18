package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}

	dataSlice := strings.Split(string(bs), "\n")

	timestamp, err := strconv.Atoi(dataSlice[0])

	if err != nil {
		os.Exit(2)
	}

	re := regexp.MustCompile("(\\d+)")
	res := re.FindAllStringSubmatch(dataSlice[1], -1)
	busIds := make([]int, len(res))

	for i, slc := range res {
		busIds[i], err = strconv.Atoi(slc[1])

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(2)
		}
	}

	minTimestamp := math.MaxInt64

	schedule := make(map[int]int)

	for _, id := range busIds {
		schedule[id] = (timestamp/id)*id + id - timestamp
	}
	minID := -1
	for id, ts := range schedule {
		if ts < minTimestamp {
			minTimestamp = ts
			minID = id
		}
	}

	fmt.Println(minTimestamp * minID)
}
