package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	groupUniqueAnswers := make(map[rune]int)
	currentGroupSize := 0
	countUnique := 0

	for _, str := range dataSlice {
		if len(str) < 1 {
			// fmt.Println("Group answer count:", groupUniqueAnswers)
			for _, v := range groupUniqueAnswers {
				if v == currentGroupSize {
					countUnique++
				}
			}
			groupUniqueAnswers = make(map[rune]int)
			currentGroupSize = 0
			continue
		}
		currentGroupSize++
		for _, a := range str {
			if _, ok := groupUniqueAnswers[a]; ok {
				groupUniqueAnswers[a]++
			} else {
				groupUniqueAnswers[a] = 1
			}
		}

		// fmt.Println(str, len(groupUniqueAnswers))
	}
	fmt.Println(countUnique)
}
