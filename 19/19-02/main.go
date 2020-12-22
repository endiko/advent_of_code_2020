package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type rule struct {
	value    string
	variants map[int][]string
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n\n")

	rulesData := strings.Split(dataSlice[0], "\n")
	msgs := strings.Split(dataSlice[1], "\n")

	rules := make(map[string]rule, len(rulesData))

	reStr := regexp.MustCompile("\"(\\w)\"")
	reNum := regexp.MustCompile("(\\d+)")

	for _, str := range rulesData {
		ruleSlice := strings.Split(str, ": ")

		tmp := rule{variants: make(map[int][]string)}
		if reStr.MatchString(ruleSlice[1]) {
			tmp.value = reStr.FindStringSubmatch(ruleSlice[1])[1]
		} else {
			groups := strings.Split(ruleSlice[1], "|")

			for i, str := range groups {
				tmp.variants[i] = reNum.FindAllString(str, -1)
			}
		}
		rules[ruleSlice[0]] = tmp
	}

	count := 0

	for _, msg := range msgs {
		total, ok := rules["0"].check(msg, rules)

		fmt.Println("total: ", total)
		fmt.Println("msg: ", msg, len(msg))
		if ok && total == len(msg) {
			// fmt.Println(msg)
			count++
		}
	}

	// fmt.Println(count)
}

func (r rule) check(str string, rules map[string]rule) (int, bool) {
	if len(str) == 0 {
		return 0, false
	}
	if r.value != "" {
		return 1, strings.HasPrefix(str, r.value)
	}

	for _, v := range r.variants {
		offset := 0
		success := true
		for _, rl := range v {
			idx, ok := rules[rl].check(str[offset:], rules)
			success = ok
			if !ok {
				break
			}
			offset += idx
		}
		if success {
			return offset, true
		}
	}
 	return 0, false
}
