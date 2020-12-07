package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Bag struct {
	parents map[string]struct{}
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	bags := make(map[string]Bag)

	ancestors := make(map[string]struct{})

	re := regexp.MustCompile("([\\w ]+) bags contain")
	subRe := regexp.MustCompile("\\d+ ([\\w ]+) bag")
	for _, str := range dataSlice {
		res := re.FindStringSubmatch(str)
		subRes := subRe.FindAllStringSubmatch(str, -1)

		if len(subRes) == 0 {
			continue
		}

		parentName := res[1]
		for _, sl := range subRes {
			currChild := sl[1]

			if bag, ok := bags[currChild]; ok {
				bag.addParentBag(parentName)
			} else {
				newBag := Bag{parents: make(map[string]struct{})}
				newBag.addParentBag(parentName)
				bags[currChild] = newBag
			}
		}
	}

	fillAncestors(bags, "shiny gold", ancestors)

	fmt.Println(len(ancestors) - 1)
}

func (b Bag) addParentBag(parentName string) {
	if _, ok := b.parents[parentName]; !ok {
		b.parents[parentName] = struct{}{}
	}
}

func fillAncestors(bags map[string]Bag, name string, ancestors map[string]struct{}) {
	if _, ok := ancestors[name]; ok {
		return
	}

	ancestors[name] = struct{}{}

	for bagName := range bags[name].parents {
		fillAncestors(bags, bagName, ancestors)
	}
}
