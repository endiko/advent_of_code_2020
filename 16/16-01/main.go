package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	low  validRange
	high validRange
}

type validRange struct {
	start int
	end   int
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dataSlice := strings.Split(string(bs), "\n\n")

	rules := make(map[string]rule)
	parseRules(dataSlice[0], rules)
	fmt.Println(rules)

	res := 0

	strs := strings.Split(dataSlice[2], "\n")
	for _, s := range strs[1:] {
		t := parseTicket(s)
		for _, num := range t {
			if !checkNumber(num, rules) {
				res += num
			}
		}
	}
	fmt.Println(res)
}

func checkNumber(num int, rules map[string]rule) bool {
	res := false
	for _, r := range rules {
		res = res || r.check(num)
	}
	return res
}

func (r rule) check(v int) bool {
	return r.low.check(v) || r.high.check(v)
}

func (r validRange) check(v int) bool {
	return v >= r.start && v <= r.end
}

func parseRules(input string, rules map[string]rule) {
	data := strings.Split(input, "\n")
	re := regexp.MustCompile("([\\w ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	for _, s := range data {
		res := re.FindStringSubmatch(s)
		curr := rule{}
		curr.low = makeRange(res[2], res[3])
		curr.high = makeRange(res[4], res[5])
		rules[res[1]] = curr
	}
}

func makeRange(s1 string, s2 string) validRange {
	res := validRange{}
	res.start, _ = strconv.Atoi(s1)
	res.end, _ = strconv.Atoi(s2)
	return res
}

func parseTicket(s string) []int {
	strs := strings.Split(s, ",")
	res := make([]int, len(strs))
	for i, val := range strs {
		res[i], _ = strconv.Atoi(val)
	}
	return res
}
