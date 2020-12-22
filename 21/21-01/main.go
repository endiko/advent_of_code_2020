package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n")
	foods := make([]food, len(dataSlice))

	for i, line := range dataSlice {
		foodList := strings.Split(line, " (")

		foods[i].ingredients = strings.Split(foodList[0], " ")
		foods[i].allergens = strings.Split(foodList[1][9:len(foodList[1])-1], ", ")
	}

	allergensMap := make(map[string]map[string]struct{})

	for _, f := range foods {
		for _, a := range f.allergens {
			if _, ok := allergensMap[a]; !ok {
				allergensMap[a] = make(map[string]struct{})
				for _, ing := range f.ingredients {
					allergensMap[a][ing] = struct{}{}
				}
			} else {
				tmp := make(map[string]struct{})
				//fmt.Println(f.ingredients)
				for _, ing := range f.ingredients {
					if _, ok := allergensMap[a][ing]; ok {
						tmp[ing] = struct{}{}
					}
				}
				allergensMap[a] = tmp
			}
			//fmt.Println(a, allergensMap[a])
		}
	}

	//fmt.Println(allergensMap)

	potentialAllergen := make(map[string]struct{})

	for _, a := range allergensMap {
		for ing := range a {
			//fmt.Println(ing)
			if _, ok := potentialAllergen[ing]; !ok {
				potentialAllergen[ing] = struct{}{}
			}
		}
	}

	count := 0
	for _, f := range foods {
		for _, ing := range f.ingredients {
			if _,ok := potentialAllergen[ing]; !ok {
				count++
			}
		}

	}



	fmt.Println(count)
}
