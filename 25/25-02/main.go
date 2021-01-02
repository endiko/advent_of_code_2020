package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


func main() {
	addressList := readInput("\n")

	tiles := make(map[[16]byte]*Tile)


}


func readInput(sep string) []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), sep)
}

