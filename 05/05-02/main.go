package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(bs), "\n")
	var maxSeatID int64 = -1
	bordingPassMap := make(map[int]struct{})

	for _, str := range dataSlice {
		str = strings.ReplaceAll(str, "F", "0")
		str = strings.ReplaceAll(str, "B", "1")
		str = strings.ReplaceAll(str, "L", "0")
		str = strings.ReplaceAll(str, "R", "1")

		currSeatID, err := strconv.ParseInt(str, 2, 32)
		if err != nil {
			panic(err)
		}

		if currSeatID > maxSeatID {
			maxSeatID = currSeatID
		}

		bordingPassMap[int(currSeatID)] = struct{}{}
	}

	for seatID := 1; seatID < int(maxSeatID); seatID++ {
		if _, ok := bordingPassMap[seatID]; ok {
			continue
		}
		_, prevExists := bordingPassMap[seatID-1]
		_, nextExists := bordingPassMap[seatID+1]

		if prevExists && nextExists {
			fmt.Println(seatID)
			break
		}
	}
}
