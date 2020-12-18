package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Rule contains ticket field rules
type Rule struct {
	name   string
	column int
	low    Range
	high   Range
}

// Range for ticket fields
type Range struct {
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

	rules := make(map[string]Rule)
	parseRules(dataSlice[0], rules)

	ourTicketBlock := strings.Split(dataSlice[1], "\n")
	ourTicket := parseTicket(ourTicketBlock[1])
	fmt.Println("our ticket:", ourTicket)

	validTickets := make(map[int][]int)
	strs := strings.Split(dataSlice[2], "\n")
	for idx, s := range strs[1:] {
		t := parseTicket(s)
		if checkTicket(t, rules) {
			validTickets[idx] = t
		}
	}

	fmt.Println("valid tickets count:", len(validTickets))
	fieldCount := len(ourTicket)

	scores := make(map[Rule]map[int]int)
	for _, r := range rules {
		scores[r] = make(map[int]int, fieldCount)
	}
	for _, t := range validTickets {
		checkRules(t, scores)
	}

	for calcNotFoundRules(rules) > 0 {
		for _, r := range rules {
			if idx, ok := findSingleCandidate(scores[r], fieldCount); ok {
				r.column = idx
				resetColumn(scores, idx)
			}
		}
		for i := 0; i < fieldCount; i++ {
			scanRules(scores, i, fieldCount)
		}
	}
	fmt.Println(rules)
}

func calcNotFoundRules(rules map[string]Rule) int {
	count := 0
	for _, r := range rules {
		if !r.found() {
			count++
		}
	}
	return count
}

func scanRules(scores map[Rule]map[int]int, idx int, targetNum int) {
	var rule Rule
	for r, cols := range scores {
		if cols[idx] == targetNum {
			if rule.name != "" {
				return
			}
			rule = r
		}
	}
	rule.column = idx
	resetColumn(scores, idx)
}

func findSingleCandidate(cols map[int]int, targetNum int) (int, bool) {
	res := -1
	for idx, val := range cols {
		if val == targetNum {
			if res != -1 {
				return -1, false
			}
			res = idx
		}
	}
	return res, true
}

func resetColumn(scores map[Rule]map[int]int, idx int) {
	for _, cols := range scores {
		cols[idx] = 0
	}
}

func (r Rule) found() bool {
	return r.column != -1
}

func checkRules(t []int, scores map[Rule]map[int]int) {
	for col, num := range t {
		for r, colScoreMap := range scores {
			if r.check(num) {
				colScoreMap[col]++
			}
		}
	}
}

func checkTicket(t []int, rules map[string]Rule) bool {
	for _, num := range t {
		if !checkNumber(num, rules) {
			return false
		}
	}
	return true
}

func checkNumber(num int, rules map[string]Rule) bool {
	for _, r := range rules {
		if r.check(num) {
			return true
		}
	}
	return false
}

func (r Rule) check(v int) bool {
	return r.low.check(v) || r.high.check(v)
}

func (r Range) check(v int) bool {
	return v >= r.start && v <= r.end
}

func parseRules(input string, rules map[string]Rule) {
	data := strings.Split(input, "\n")
	re := regexp.MustCompile("([\\w ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	for _, s := range data {
		res := re.FindStringSubmatch(s)
		curr := Rule{}
		curr.low = makeRange(res[2], res[3])
		curr.high = makeRange(res[4], res[5])
		curr.name = res[1]
		curr.column = -1
		rules[res[1]] = curr
	}
}

func makeRange(s1 string, s2 string) Range {
	res := Range{}
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
