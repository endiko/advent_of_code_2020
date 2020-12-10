package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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
	data := make([]int64, len(dataSlice))

	for i, str := range dataSlice {
		num, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			fmt.Println("Не удалось преобразовать в число")
			os.Exit(2)
		}

		data[i] = num
	}

	index, ok := findNumbersIndex(data)

	if !ok {
		os.Exit(3)
	}
	i, j := findRange(data, index)
	fmt.Println(data[i: j+1])
	min, max := findMinMax(data[i:j+1])

	fmt.Println(i, j)
	fmt.Println(min, max)
	fmt.Println(min + max)
}

func checkSum(data []int64, currentIdx int) bool {
	numsMap := make(map[int64]struct{})

	for _, num := range data[currentIdx-25 : currentIdx] {
		numsMap[num] = struct{}{}
	}

	targetNum := data[currentIdx]

	for nm := range numsMap {
		diff := targetNum - nm

		if _, ok := numsMap[diff]; ok {
			return true
		}
	}
	// fmt.Println(numsMap, targetNum)
	return false
}

func findNumbersIndex(data []int64) (int, bool) {
	for i := 26; i < len(data); i++ {
		if !checkSum(data, i) {
			fmt.Println("the number is", data[i])
			return i, true
		}
	}
	return -1, false
}

func findRange(data []int64, idx int) (int, int) {
	subtotal := make(map[int]int64)
	sum := int64(0)
	targetSum := data[idx]
	fmt.Println("targetSum ", targetSum)
	for i, num := range data[0:idx] {
		sum += num

		if sum == targetSum {
			return 0, i
		}

		for j, num := range subtotal {
			if sum-num == targetSum {
				// fmt.Println(subtotal)
				return j + 1, i
			}
		}

		subtotal[i] = sum
	}

	return -1, -1
}

func findMinMax(data []int64) (int64, int64) {
	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)

	for _, num := range data {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return min, max
}
