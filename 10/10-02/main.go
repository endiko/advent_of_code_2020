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
	data = append(data, 0, data[len(data)-1]+3)
	sort.Ints(data)

	chunks := make(map[int][]int)
	currChunk := 0
	startPos := 0
	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]
		if diff == 3 {
			chunks[currChunk] = data[startPos:i]
			currChunk++
			startPos = i
		}
		if i == len(data)-1 {
			chunks[currChunk] = data[startPos:]
		}
	}

	product := 1
	for _, val := range chunks {
		num := calcVariants(val)
		product *= num
	}
	fmt.Println(product)
}

func calcVariants(data []int) int {
	if len(data) < 3 {
		return 1
	}
	count := calcVariants(data[1:])
	if data[2]-data[0] <= 3 {
		count += calcVariants(data[2:])
	}
	if len(data) > 3 && data[3]-data[0] <= 3 {
		count += calcVariants(data[3:])
	}
	return count
}
