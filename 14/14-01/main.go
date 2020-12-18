package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
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

	mem := make(map[int]uint64)
	var strMask string
	mask := make(map[int]uint64)
	rxMask := regexp.MustCompile("mask = ([X01]{36})")      // mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	rxNum := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)") // mem[7] = 101
	for _, s := range dataSlice {
		if strings.HasPrefix(s, "mask") {
			strMask = rxMask.FindStringSubmatch(s)[1]
			parseMask(strMask, mask)
			// fmt.Printf("%64s\n", strMask)
			// fmt.Printf("%064b\n%064b\n", mask[0], mask[1])
		} else {
			nums := rxNum.FindStringSubmatch(s)

			idx, err := strconv.Atoi(nums[1])
			if err != nil {
				os.Exit(2)
			}
			num, err := strconv.ParseUint(nums[2], 10, 64)
			if err != nil {
				os.Exit(3)
			}
			// fmt.Printf("value:   %064b\n", num)
			// fmt.Printf("mask[0]: %064b\n", mask[0])
			// fmt.Printf("mask[1]: %064b\n", mask[1])
			num &= mask[0]
			num |= mask[1]
			// fmt.Printf("result:  %064b\n\n", num)
			mem[idx] = num
		}
	}

	total := uint64(0)
	for _, v := range mem {
		if v == 0 {
			continue
		}
		total += v
	}
	fmt.Println("Total = ", total)
}

func parseMask(s string, mask map[int]uint64) {
	mask[0] = math.MaxUint64
	mask[1] = 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0':
			mask[0] &= ^(1 << (len(s) - i - 1))
		case '1':
			mask[1] |= (1 << (len(s) - i - 1))
		}
	}
}
