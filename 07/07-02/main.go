package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	children map[string]int
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	bags := make(map[string]Bag)

	re := regexp.MustCompile("([\\w ]+) bags contain")
	subRe := regexp.MustCompile("(\\d+) ([\\w ]+) bag")
	for _, str := range dataSlice {
		res := re.FindStringSubmatch(str)
		subRes := subRe.FindAllStringSubmatch(str, -1)

		if len(subRes) == 0 {
			continue
		}

		parentName := res[1]
		currentBag := Bag{children: make(map[string]int)}
		for _, sl := range subRes {
			currChildAmount, err := strconv.Atoi(sl[1])

			if err != nil {
				fmt.Println("Не удалось преобразовать в число")
				os.Exit(1)
			}

			currChild := sl[2]

			currentBag.children[currChild] = currChildAmount
		}
		bags[parentName] = currentBag
	}

	total := calcOffspring(bags, "shiny gold")

	fmt.Println(total - 1)
}

func calcOffspring(bags map[string]Bag, name string) int {
	sum := 1
	for childName, amount := range bags[name].children {
		sum += amount * calcOffspring(bags, childName)
	}

	return sum
}
