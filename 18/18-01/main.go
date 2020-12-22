package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type expression struct {
	numbers map[int]int
	op      operation
}

type operation func(l int, r int) int

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n")

	total := 0
	re := regexp.MustCompile("[\\d*+\\(\\)]")
	for _, s := range dataSlice {
		res := re.FindAllStringSubmatch(s, -1)
		e := res[0]
		for i := 1; i < len(res); i++ {
			e = append(e, res[i]...)
		}
		// fmt.Println(s)
		total += process(e)
	}
	fmt.Println(total)
}

func process(e []string) int {
	// fmt.Println(e)
	result, _ := strconv.Atoi(e[0])
	increment := func(l int, r int) int { return l + r }
	production := func(l int, r int) int { return l * r }

	ex := expression{numbers: make(map[int]int), op: nil}
	for i := 0; i < len(e); i++ {
		switch e[i] {
		case "*":
			ex.op = production
		case "+":
			ex.op = increment
		case "(":
			idx := findPairParenthesisIdx(e, i)
			num := process(e[i+1 : idx])
			ex.addNumber(num)
			i = idx
		default:
			num, _ := strconv.Atoi(e[i])
			ex.addNumber(num)
		}
		if ex.ready() {
			result = ex.run()
			ex = expression{numbers: make(map[int]int), op: nil}
			ex.addNumber(result)
		}
	}

	return result
}

func (e expression) run() int {
	// fmt.Println(e.numbers, e.op(e.numbers[0], e.numbers[1]))
	return e.op(e.numbers[0], e.numbers[1])
}

func (e expression) addNumber(num int) {
	e.numbers[len(e.numbers)] = num
}

func (e expression) ready() bool {
	if len(e.numbers) < 2 {
		return false
	}
	if e.op == nil {
		return false
	}
	return true
}

func findPairParenthesisIdx(e []string, idx int) int {
	level := 0
	for i := idx; i < len(e); i++ {
		switch e[i] {
		case "(":
			level++
		case ")":
			level--
		default:
			continue
		}
		if level == 0 {
			return i
		}
	}
	return -1
}
