package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mask struct {
	main     uint64
	floating []uint64
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")

	mem := make(map[uint64]uint64)
	var strMask string
	m := mask{}
	rxMask := regexp.MustCompile("mask = ([X01]{36})")      // mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	rxNum := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)") // mem[7] = 101
	for _, s := range dataSlice {
		if strings.HasPrefix(s, "mask") {
			strMask = rxMask.FindStringSubmatch(s)[1]
			parseMask(strMask, &m)
		} else {
			nums := rxNum.FindStringSubmatch(s)

			idx, err := strconv.ParseUint(nums[1], 10, 64)
			if err != nil {
				os.Exit(2)
			}
			num, err := strconv.ParseUint(nums[2], 10, 64)
			if err != nil {
				os.Exit(3)
			}
			// fmt.Printf("%064b - address\n", idx)
			// fmt.Printf("%064b - main mask\n", m.main)
			idx |= m.main
			for _, f := range m.floating {
				// fmt.Printf("%064b - floating\n", f)
				// fmt.Printf("%064b - xor\n\n", idx^f)
				mem[idx^f] = num
			}
			// fmt.Println(mem)
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

func parseMask(s string, mask *mask) {
	mask.main = 0
	idxList := make([]int, strings.Count(s, "X"))
	pos := 0
	varCount := 1

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'X':
			idxList[pos] = len(s) - i - 1
			pos++
			varCount <<= 1
		case '1':
			mask.main |= (1 << (len(s) - i - 1))
		}
	}

	mask.floating = make([]uint64, varCount)
	for i := 0; i < varCount; i++ {
		var currMask uint64 = 0
		for j, idx := range idxList {
			if i&(1<<j) > 0 {
				currMask |= 1 << idx
			}
		}
		mask.floating[i] = currMask
	}
	// fmt.Printf("%64s\n", s)
	// fmt.Printf("%064b\n", mask.main)
	// for _, m := range mask.floating {
	// 	fmt.Printf("%064b\n", m)
	// }
	// fmt.Println()
}
