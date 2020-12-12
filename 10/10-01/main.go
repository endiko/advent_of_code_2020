package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")
	data := make([]int, len(dataSlice))

	for i, str := range dataSlice {
		num, err := strconv.Atoi(str)

		if err != nil {
			fmt.Println("Не удалось преобразовать в число")
			os.Exit(2)
		}

		data[i] = num
	}

	sort.Ints(data)

	fmt.Println(data)

	differences := make(map[int]int)

	prevJoltage := 0
	for _, i := range data {
		diff := i - prevJoltage
		differences[diff]++
		prevJoltage = i
		fmt.Printf("%d +%d : %d\n", i, diff, differences[diff])
	}
	differences[3]++
	fmt.Println(differences)
	fmt.Println(differences[3] * differences[1])
}
