package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput("\n")

	cardPK, _ := strconv.Atoi(data[0])
	doorPK, _ := strconv.Atoi(data[1])
	cardLS := findLoopSize(cardPK)
	doorLS := findLoopSize(doorPK)

	cardEK := calcEncryptionKey(cardPK, doorLS)
	doorEK := calcEncryptionKey(doorPK, cardLS)


		fmt.Println(cardEK, doorEK)

}

func calcEncryptionKey(pk int, ls int) (result int) {
	result = 1
	for i:= 0; i< ls; i++ {
		result *= pk
		result %= 20201227
	}
	return
}

func findLoopSize(pk int) int {
	val := 1
	count := 1

	for {
		val *= 7
		val %= 20201227
		if val == pk {
			return count
		}
		count++
	}
}


func readInput(sep string) []string {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), sep)
}
