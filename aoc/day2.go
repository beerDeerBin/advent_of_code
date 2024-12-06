package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day2Level1(inputFileName string) {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	safeLevel := 0
	for _, line := range lines {
		if checkLevel(line) {
			safeLevel++
		}
	}
	fmt.Println(safeLevel)
}

func checkLevel(level string) (isSafe bool) {
	levelSteps := strings.Split(level, " ")

	lastValue := 0
	isDecreasing := false
	for i, step := range levelSteps {
		value, err := strconv.Atoi(step)
		check(err)
		if i == 0 {
			lastValue = value
		} else {
			diff := int(math.Abs(float64(value) - float64(lastValue)))
			if diff == 0 {
				return false
			}
			if diff > 3 {
				return false
			}
			if i == 1 {
				isDecreasing = lastValue > value
			} else {
				if (isDecreasing && lastValue < value) || (!isDecreasing && lastValue > value) {
					return false
				}
			}
			lastValue = value
		}
	}

	return true
}

func Day2Level2(inputFileName string) {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	safeLevel := 0
	for _, line := range lines {
		level := make([]int, 0)
		levelSteps := strings.Split(line, " ")
		for _, step := range levelSteps {
			value, err := strconv.Atoi(step)
			check(err)
			level = append(level, value)
		}
		if checkLevel2(level) {
			safeLevel++
		} else {
			buffer := make([]int, len(level))
			for i := range level {
				copy(buffer, level)
				changedLevel := append(buffer[0:i], buffer[i+1:]...)
				if checkLevel2(changedLevel) {
					safeLevel++
					break
				}
			}
		}
	}
	fmt.Println(safeLevel)
}

func checkLevel2(level []int) (isSafe bool) {
	lastValue := 0
	isDecreasing := false
	for i, step := range level {
		if i == 0 {
			lastValue = step
		} else {
			diff := int(math.Abs(float64(step) - float64(lastValue)))
			if diff == 0 {
				return false
			}
			if diff > 3 {
				return false
			}
			if i == 1 {
				isDecreasing = lastValue > step
			} else {
				if (isDecreasing && lastValue < step) || (!isDecreasing && lastValue > step) {
					return false
				}
			}
			lastValue = step
		}
	}

	return true
}
