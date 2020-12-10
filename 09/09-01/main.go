package main

import (
	"fmt"
	"io/ioutil"
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

	for i := 26; i < len(data); i++ {
		if !checkSum(data, i) {
			fmt.Println(data[i], i)
			return
		}
	}
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
	fmt.Println(numsMap, targetNum)
	return false
}
