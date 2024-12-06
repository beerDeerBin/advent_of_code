package aoc

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1Level1(inputFileName string) int {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	leftValues := make([]int, 0)
	rightValues := make([]int, 0)
	for _, line := range lines {
		values := strings.Split(line, "   ")
		val, err := strconv.Atoi(values[0])
		check(err)
		leftValues = append(leftValues, val)
		val, err = strconv.Atoi(values[1])
		check(err)
		rightValues = append(rightValues, val)
	}

	// Sort the values
	sort.Ints(leftValues)
	sort.Ints(rightValues)

	sum := 0
	for i, val := range leftValues {
		sum += int(math.Abs(float64(val) - float64(rightValues[i])))
	}

	return sum
}

func Day1Level2(inputFileName string) int {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	leftValues := make([]int, 0)
	rightValues := make(map[int]int)
	for _, line := range lines {
		values := strings.Split(line, "   ")

		val, err := strconv.Atoi(values[0])
		check(err)
		leftValues = append(leftValues, val)

		val, err = strconv.Atoi(values[1])
		check(err)
		rightValues[val] = rightValues[val] + 1
	}

	sum := 0

	for _, key := range leftValues {
		similarity_score := key * rightValues[key]
		sum += similarity_score
	}

	return sum
}
