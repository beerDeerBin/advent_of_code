package aoc

import (
	"strconv"
	"strings"
)

type Equation struct {
	answer   int
	operands []int
}

func Day7Level1(inputFileName string) int {
	
	equations := parseInputDay7(inputFileName)
	sum := 0
	for _, equation := range equations {
		if testAllPossibleEquations(equation.operands, equation.answer, []string{"add", "mul"}) {
			sum += equation.answer
		}
	}

	return sum
}

func Day7Level2(inputFileName string) int {

	equations := parseInputDay7(inputFileName)
	sum := 0
	for _, equation := range equations {
		if testAllPossibleEquations(equation.operands, equation.answer, []string{"add", "mul", "concat"}) {
			sum += equation.answer
		}
	}

	return sum
}

func parseInputDay7(inputFileName string) []Equation {
	lines := readFileAsLines(inputFileName)

	equations := make([]Equation, len(lines))

	for i, line := range lines {
		line = strings.Replace(line, ":", "", -1)
		split := strings.Split(line, " ")
		equation := Equation{}
		num, err := strconv.Atoi(split[0])
		check(err)
		equation.answer = num
		split = split[1:] // Remove the total number from the split
		equation.operands = make([]int, len(split))
		for j, s := range split {
			num, err := strconv.Atoi(s)
			check(err)
			equation.operands[j] = num
		}
		equations[i] = equation
	}

	return equations
}

func testAllPossibleEquations(numbers []int, answer int, ops []string) bool {
	if len(numbers) == 1 {
		return numbers[0] == answer
	}

	firstTwo := numbers[:2]
	rest := numbers[2:]
	for _, op := range ops {
		newNumber := firstTwo[0] + firstTwo[1]
		if op == "mul" {
			newNumber = firstTwo[0] * firstTwo[1]
		} else if op == "concat" {
			newNumber, _ = strconv.Atoi(strconv.Itoa(firstTwo[0]) + strconv.Itoa(firstTwo[1]))
		}
		if testAllPossibleEquations(append([]int{newNumber}, rest...), answer, ops) {
			return true
		}
	}
	return false
}
