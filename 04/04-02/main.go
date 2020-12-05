package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	rx := regexp.MustCompile("(\\w+):(\\S+)")

	dataSlice := strings.Split(string(bs), "\n")
	currCredentials := make(map[string]string)
	validCount := 0
	for _, currLine := range dataSlice {
		if len(currLine) <= 1 {
			if validate(currCredentials) {
				validCount++
			} else {
				// fmt.Println("Checking failed", currCredentials)
			}
			currCredentials = make(map[string]string)
			continue
		}
		tmp := rx.FindAllStringSubmatch(currLine, -1)
		for _, keySlice := range tmp {
			currCredentials[keySlice[1]] = keySlice[2]
		}
	}
	fmt.Println(validCount)
}

func validate(data map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	if !checkRequiredFieldsPresence(data, requiredFields) {
		return false
	}
	if !checkNumRange(data["byr"], 1920, 2002) {
		// fmt.Println("Field 'byr' out of range: ", data["byr"], 1920, 2002)
		return false
	}
	if !checkNumRange(data["iyr"], 2010, 2020) {
		// fmt.Println("Field 'iyr' out of range: ", data["iyr"], 2010, 2020)
		return false
	}
	if !checkNumRange(data["eyr"], 2020, 2030) {
		// fmt.Println("Field 'eyr' out of range: ", data["eyr"], 2020, 2030)
		return false
	}
	if !checkHeight(data["hgt"]) {
		return false
	}
	if !checkStringValue(data["hcl"], "#[\\da-f]{6}") {
		// fmt.Println("Field 'hcl' is not valid: ", data["hcl"])
		return false
	}
	if !checkStringValue(data["ecl"], "amb|blu|brn|gry|grn|hzl|oth") {
		// fmt.Println("Field 'ecl' is not valid: ", data["ecl"])
		return false
	}
	if !checkStringValue(data["pid"], "^\\d{9}\\z") {
		// fmt.Println("Field 'pid' is not valid: ", data["pid"])
		return false
	}
	return true
}

func checkNumRange(str string, minValue int, maxValue int) bool {
	value, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return value >= minValue && value <= maxValue
}

func checkStringValue(str string, expression string) bool {
	rx := regexp.MustCompile(expression)
	return rx.MatchString(str)
}

func checkHeight(str string) bool {
	rx := regexp.MustCompile("(\\d+)(cm|in)")
	res := rx.FindAllStringSubmatch(str, -1)
	// fmt.Println(str, res)
	if len(res) == 0 {
		// fmt.Println("Field 'hgt' is not valid: ", str)
		return false
	}
	height, err := strconv.Atoi(res[0][1])
	if err != nil {
		// fmt.Println("Field 'hgt' is not valid number: ", str)
		return false
	}
	if res[0][2] == "cm" {
		// if height < 150 {
		// 	fmt.Println("Field 'hgt' is too low: ", str)
		// 	return false
		// }
		// if height > 193 {
		// 	fmt.Println("Field 'hgt' is too high: ", str)
		// 	return false
		// }
		// return true
		return height >= 150 && height <= 193
	}
	// if height < 59 {
	// 	fmt.Println("Field 'hgt' is too low valid: ", str)
	// 	return false
	// }
	// if height > 76 {
	// 	fmt.Println("Field 'hgt' is too high: ", str)
	// 	return false
	// }
	// return true
	return height >= 59 && height <= 76
}

func checkRequiredFieldsPresence(data map[string]string, requiredFields []string) bool {
	for _, field := range requiredFields {
		if _, ok := data[field]; !ok {
			// fmt.Println("field not found: ", field)
			return false
		}
	}
	return true
}
