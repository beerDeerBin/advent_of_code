package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5Level1(inputFileName string) {
	lines := readFileAsLines(inputFileName)

	rules := make(map[string]bool, 0)
	sum := 0

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
				isValid := true
				for i := range numbers {
					for j := i + 1; j < len(numbers); j++ {
						if !verifyRule(rules, numbers[i], numbers[j]) {
							isValid = false
						}
					}
				}
				if isValid {
					sum += numbers[len(numbers)/2]
				}
			}
		}
	}

	fmt.Println(sum)
}

func verifyRule(rules map[string]bool, before int, after int) bool {
	rule := strconv.Itoa(before) + "|" + strconv.Itoa(after)
	if _, ok := (rules)[rule]; ok {
		return true
	}
	return false
}

func Day5Level2(inputFileName string) {
	lines := readFileAsLines(inputFileName)

	rules := make(map[string]bool, 0)
	sum := 0

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
				isValid := true
				for i := range numbers {
					for j := i + 1; j < len(numbers); j++ {
						if !verifyRule(rules, numbers[i], numbers[j]) {
							isValid = false
							break
						}
					}
				}
				if !isValid {
					numbers = orderPage(rules, numbers)
					sum += numbers[len(numbers)/2]
				}
			}
		}
	}

	fmt.Println(sum)
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
