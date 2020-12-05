package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	total := 0

	re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
	match := re.FindAllStringSubmatch(string(bs), -1)
	for _, curr := range match {
		first, _ := strconv.Atoi(curr[1])
		second, _ := strconv.Atoi(curr[2])
		letter := curr[3]
		pass := curr[4]

		check := regexp.MustCompile(fmt.Sprintf("%s{1}", letter))
		indexes := check.FindAllStringIndex(pass, -1)
		startIndexes := make(map[int]int)
		for _, segment := range indexes {
			startIndexes[segment[0]] = segment[1]
		}
		_, okFirst := startIndexes[first-1]
		_, okSecond := startIndexes[second-1]
		if okFirst != okSecond {
			total++
		}
		// fmt.Printf("searching for %s in '%s' It must be on %d or %d position. %v - %v\n", letter, pass, first, second, okFirst, okSecond)
	}
	fmt.Println(total)
}
