package aoc

import (
	"strconv"
	"strings"
)

func Day5Level1(inputFileName string) int {

	rules, pages := parseInputDay5(inputFileName)
	sum := 0

	for _, page := range pages {
		isValid := true
		for i := range page {
			for j := i + 1; j < len(page); j++ {
				if !verifyRule(rules, page[i], page[j]) {
					isValid = false
				}
			}
		}
		if isValid {
			sum += page[len(page)/2]
		}
	}

	return sum
}

func Day5Level2(inputFileName string) int {
	
	rules, pages := parseInputDay5(inputFileName)
	sum := 0

	for _, page := range pages {
		isValid := true
		for i := range page {
			for j := i + 1; j < len(page); j++ {
				if !verifyRule(rules, page[i], page[j]) {
					isValid = false
				}
			}
		}
		if !isValid {
			orderPage := orderPage(rules, page)
			sum += orderPage[len(orderPage)/2]
		}
	}

	return sum
}

func parseInputDay5(inputFileName string) (map[string]bool, [][]int) {
	lines := readFileAsLines(inputFileName)

	rules := make(map[string]bool, 0)
	pages := make([][]int, 0)

	for _, line := range lines {
		if line != "" {
			if strings.Contains(line, "|") {
				rules[line] = true
			} else {
				splits := strings.Split(line, ",")
				numbers := make([]int, 0)
				for _, split := range splits {
					number, err := strconv.Atoi(split)
					check(err)
					numbers = append(numbers, number)
				}
				pages = append(pages, numbers)
			}
		}
	}

	return rules, pages
}

func verifyRule(rules map[string]bool, before int, after int) bool {
	rule := strconv.Itoa(before) + "|" + strconv.Itoa(after)
	if _, ok := (rules)[rule]; ok {
		return true
	}
	return false
}

func orderPage(rules map[string]bool, page []int) []int {
	for i := range page {
		for j := i + 1; j < len(page); j++ {
			rule := strconv.Itoa(page[i]) + "|" + strconv.Itoa(page[j])
			if _, ok := rules[rule]; !ok {
				buffer := page[i]
				page[i] = page[j]
				page[j] = buffer
				return orderPage(rules, page)
			}
		}
	}
	return page
}
