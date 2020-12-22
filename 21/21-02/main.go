package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	for calcPotentialAllergens(allergensMap) > 0 {
		for a := range allergensMap {
			if len(allergensMap[a]) == 1 {
				for ia := range allergensMap {
					if ia != a {
						delete(allergensMap[ia], getFirstKey(allergensMap[a]))
					}
				}
			}
		}
	}

	sortedAllergens := sortByCanonicalAllergens(allergensMap)
	fmt.Println(getAllergensList(sortedAllergens, allergensMap))
}

func getFirstKey(data map[string]struct{}) (result string) {
	for key := range data {
		result = key
		break
	}
	return
}

func calcPotentialAllergens(allergensMap map[string]map[string]struct{}) (result int) {
	for a := range allergensMap {
		if len(allergensMap[a]) > 1 {
			result++
		}
	}
	return
}

func sortByCanonicalAllergens(data map[string]map[string]struct{}) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getAllergensList(data []string, allergensMap map[string]map[string]struct{}) string {
	var temp []string
	for _, k := range data {
		firstKey := getFirstKey(allergensMap[k])
		temp = append(temp, firstKey)
	}
	return strings.Join(temp, ",")
}
