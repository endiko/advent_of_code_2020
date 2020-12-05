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
		min, _ := strconv.Atoi(curr[1])
		max, _ := strconv.Atoi(curr[2])
		letter := curr[3]
		pass := curr[4]

		check := regexp.MustCompile(fmt.Sprintf("(%s)", letter))
		num := check.FindAllStringSubmatch(pass, -1)
		if min <= len(num) && len(num) <= max {
			total++
		}
		fmt.Printf("searching for %s in '%s' from %s to %s times. Found %d\n", curr[3], curr[4], curr[1], curr[2], len(num))
	}
	fmt.Println(total)
}
