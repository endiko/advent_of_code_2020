package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Chinese remainder theorem

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}

	dataSlice := strings.Split(string(bs), "\n")

	for _,val := range dataSlice[1]{
		
	}

	// re := regexp.MustCompile("(\\d+)")
	// res := re.FindAllStringSubmatch(dataSlice[1], -1)
	// busIds := make([]int, len(res))

	// for i, slc := range res {
	// 	busIds[i], err = strconv.Atoi(slc[1])

	// 	if err != nil {
	// 		fmt.Println("Error: ", err)
	// 		os.Exit(2)
	// 	}
	// }

}
