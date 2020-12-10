package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Типы операций
const (
	ACC = "acc"
	JMP = "jmp"
	NOP = "nop"
)

// Action - Сохраняем шаг
type Action struct {
	actionType string
	step       int
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")
	instructionSlice := make([]Action, len(dataSlice))

	rx := regexp.MustCompile("(\\w{3}) \\+?([-\\d]+)")
	for i, str := range dataSlice {
		res := rx.FindStringSubmatch(str)
		instructionSlice[i].actionType = res[1]
		step, err := strconv.Atoi(res[2])
		if err != nil {
			os.Exit(2)
		}
		instructionSlice[i].step = step
	}

	sum := 0
	path := make(map[int]struct{})
	pos := 0
	for {
		if _, ok := path[pos]; ok {
			break
		}
		path[pos] = struct{}{}

		action := instructionSlice[pos]

		fmt.Printf("Line %3d : %s %d === ", pos, action.actionType, action.step)
		switch action.actionType {
		case ACC:
			sum += action.step
			pos++
		case JMP:
			pos += action.step
		default:
			pos++
		}
		fmt.Println(sum)
	}

	fmt.Println(sum)
	fmt.Println(len(path))
}
