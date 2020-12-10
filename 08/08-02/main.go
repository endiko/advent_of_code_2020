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

// Parents - просто Parents
type Parents struct {
	value map[int]struct{}
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")
	instructionSlice := make([]Action, len(dataSlice))
	relations := make(map[int]Parents)

	rx := regexp.MustCompile("(\\w{3}) \\+?([-\\d]+)")
	for i, str := range dataSlice {
		res := rx.FindStringSubmatch(str)
		instructionSlice[i].actionType = res[1]
		step, err := strconv.Atoi(res[2])
		if err != nil {
			os.Exit(2)
		}
		instructionSlice[i].step = step

		target := 0
		if res[1] == JMP {
			target = i + step
		} else {
			target = i + 1
		}

		rec, ok := relations[target]
		if !ok {
			rec = Parents{value: make(map[int]struct{})}
			relations[target] = rec
		}
		rec.addParent(i)
	}

	ancestors := make(map[int]struct{})
	fillAncestors(relations, len(instructionSlice), ancestors)

	sum := 0
	path := make(map[int]struct{})
	pos := 0
	fixed := false
	for pos != len(instructionSlice) {
		if _, ok := path[pos]; ok {
			fmt.Println("still infinite loop")
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
			_, ok := ancestors[pos+1]
			if ok && !fixed {
				fmt.Print("we step")
				fixed = true
				pos++
			} else {
				pos += action.step
			}
		case NOP:
			_, ok := ancestors[pos+action.step]
			if ok && !fixed {
				fixed = true
				fmt.Print("we jump")
				pos += action.step
			} else {
				pos++
			}
		}
		fmt.Println()
		//  fmt.Println(sum)
	}

	fmt.Println(sum)
	// fmt.Println(ancestors)
	fmt.Println(path)
}

func (p Parents) addParent(idx int) {
	p.value[idx] = struct{}{}
}

func fillAncestors(relations map[int]Parents, idx int, ancestors map[int]struct{}) {
	if _, ok := ancestors[idx]; ok {
		return
	}

	ancestors[idx] = struct{}{}

	for id := range relations[idx].value {
		fillAncestors(relations, id, ancestors)
	}
}

func trySwap(idx int, ancestors map[int]struct{}) {

}
