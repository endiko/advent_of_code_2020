package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\n")
	x := 0
	bump := 0 // trees counter
	for _, str := range s {
		if str[x] == '#' {
			bump++
		}

		x += 3

		if x >= len(str) {
			x -= len(str)
		}
	}
	fmt.Println(bump)
}
