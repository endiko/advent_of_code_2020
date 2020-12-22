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
	left  func() int
	right func() int
	op    string
}

type operation func(l int, r int) int

var operations = map[string]func(l int, r int) int{
	"+": func(l int, r int) int {
		fmt.Println(l, "+", r)
		return l + r
	},
	"*": func(l int, r int) int {
		fmt.Println(l, "*", r)
		return l * r
	},
}

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
		fmt.Println(s, "=", process(e))
		total += process(e)
	}
	fmt.Println(total)
}

func process(e []string) int {
	root := expression{}
	curr := root
	for i := 0; i < len(e); i++ {
		switch e[i] {
		case "*":
			if curr.op == "" {
				curr.op = "*"
			} else {
				tmp := expression{}
				tmp.op = "*"
				tmp.left = curr.get
				curr = tmp
			}
		case "+":
			if curr.op == "" {
				curr.op = "+"
			} else {
				tmp := expression{}
				tmp.op = "+"
				tmp.left = curr.right
				curr.right = tmp.get
				curr = tmp
			}
		case "(":
			// idx := findPairParenthesisIdx(e, i)
			// num := process(e[i+1 : idx])
			// ex.addNumber(num)
			// i = idx
		default:
			num, _ := strconv.Atoi(e[i])
			if curr.op == "" {
				curr.left = func() int { return num }
			} else {
				curr.right = func() int { return num }
			}
		}
	}

	return root.get()
}

func (e expression) get() int {
	if e.right == nil {
		return e.left()
	}
	return operations[e.op](e.left(), e.right())
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
