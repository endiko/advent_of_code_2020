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
	jump := 1

	bump := 0 // trees counter
	for i, str := range s {
		if i%2 == 1 {
			continue
		}
		if str[x] == '#' {
			bump++
		}

		x += jump

		if x >= len(str) {
			x -= len(str)
		}
	}
	fmt.Println(bump)
}
