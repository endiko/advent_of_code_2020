package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	year := 2020
	nums := make(map[int]int)

	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\n")

	for i, num := range s {
		curr, err := strconv.Atoi(num)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		found, result := findNum(year-curr, nums)

		if found {
			fmt.Println(result * curr)
			os.Exit(0)
		}

		nums[curr] = i
	}

}

func findNum(target int, hash map[int]int) (bool, int) {
	for key := range hash {
		temp := target - key

		if _, ok := hash[temp]; ok {
			return true, key * temp
		}
	}

	return false, 0
}
