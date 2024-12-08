package aoc

import (
	"math"
	"strconv"
	"strings"
)

func Day2Level1(inputFileName string) int {

	levels := parseInputDay2(inputFileName)

	safeLevel := 0
	for _, level := range levels {
		if checkLevel(level, false, -1) {
			safeLevel++
		}
	}

	return safeLevel
}

func Day2Level2(inputFileName string) int {

	levels := parseInputDay2(inputFileName)

	safeLevel := 0
	for _, level := range levels {
		if checkLevel(level, true, -1) {
			safeLevel++
		}
	}

	return safeLevel
}

func parseInputDay2(inputFileName string) [][]int {
	lines := readFileAsLines(inputFileName)
	levels := make([][]int, 0)

	for _, line := range lines {
		level := make([]int, 0)
		levelSteps := strings.Split(line, " ")
		for _, step := range levelSteps {
			value, err := strconv.Atoi(step)
			check(err)
			level = append(level, value)
		}
		levels = append(levels, level)
	}

	return levels
}

func checkLevel(level []int, extened bool, removeIdx int) bool {
	buffer := make([]int, len(level))
	copy(buffer, level)
	if len(level) == removeIdx {
		return false
	}
	if removeIdx != -1 {
		level = append(level[:removeIdx], level[removeIdx+1:]...)
	}
	lastValue := level[0]
	isDecreasing := lastValue > level[1]
	isSafe := true
	for _, step := range level[1:] {
		diff := int(math.Abs(float64(step) - float64(lastValue)))
		if diff == 0 {
			isSafe = false
			break
		}
		if diff > 3 {
			isSafe = false
			break
		}
		if (isDecreasing && lastValue < step) || (!isDecreasing && lastValue > step) {
			isSafe = false
			break
		}
		lastValue = step
	}

	if !isSafe && extened {
		return checkLevel(buffer, true, removeIdx+1)
	}

	return isSafe
}
