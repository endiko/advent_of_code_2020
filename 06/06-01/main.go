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

	var Empty struct{}
	groupUniqueAnswers := make(map[rune]struct{})
	countUnique := 0

	for _, str := range dataSlice {
		if len(str) < 1 {
			fmt.Println("Group unique answer count:", len(groupUniqueAnswers))
			countUnique += len(groupUniqueAnswers)
			groupUniqueAnswers = make(map[rune]struct{})
			continue
		}

		for _, a := range str {
			groupUniqueAnswers[a] = Empty
		}

		fmt.Println(str, len(groupUniqueAnswers))
	}
	fmt.Println(countUnique)
}
