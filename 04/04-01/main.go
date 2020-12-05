package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	/*
	   byr (Birth Year)
	   iyr (Issue Year)
	   eyr (Expiration Year)
	   hgt (Height)
	   hcl (Hair Color)
	   ecl (Eye Color)
	   pid (Passport ID)
	   cid (Country ID)
	*/

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	rx := regexp.MustCompile("(\\w{3}):")

	dataSlice := strings.Split(string(bs), "\n")
	currCredentials := make(map[string]struct{})
	validCount := 0
	for _, currLine := range dataSlice {
		if len(currLine) <= 1 {
			if checkCredentials(currCredentials, requiredFields) {
				validCount++
			}
			currCredentials = make(map[string]struct{})
			continue
		}
		tmp := rx.FindAllStringSubmatch(currLine, -1)
		for _, keySlice := range tmp {
			currCredentials[keySlice[1]] = struct{}{}
		}
	}
	fmt.Println(validCount)
}

func checkCredentials(data map[string]struct{}, requiredFields []string) bool {
	for _, field := range requiredFields {
		if _, ok := data[field]; !ok {
			// fmt.Println("field not found: ", field)
			return false
		}
	}
	// fmt.Println("All fields found")
	return true
}
