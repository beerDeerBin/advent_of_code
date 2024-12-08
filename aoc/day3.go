package aoc

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	arguments1 int 
	arguments2 int 
}

func Day3Level1(inputFileName string) int {

	instructions, keys := parseInputDay3(inputFileName)

	sum := 0
	for _, key := range keys {
		instruction := instructions[key]
		if instruction.operation == "mul" {
			sum += instruction.arguments1 * instruction.arguments2
		}
	}

	return sum
}

func Day3Level2(inputFileName string) int {

	instructions, keys := parseInputDay3(inputFileName)

	isEnabled := true
	sum := 0
	for _, key := range keys {
		instruction := instructions[key]
		if instruction.operation == "mul" {
			if isEnabled {
				sum += instruction.arguments1 * instruction.arguments2
			}
		} else if instruction.operation == "do" {
			isEnabled = true
		} else if instruction.operation == "dont" {
			isEnabled = false
		}
	}

	return sum
}

func parseInputDay3(inputFileName string) (map[int]Instruction, []int) {
	lines := readFileAsLines(inputFileName)

	instructions := make(map[int]Instruction)

	reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	inputString := ""
	for _, line := range lines {
		inputString += line
	}

	muls := reMul.FindAllStringIndex(inputString, -1)
	for _, mul := range muls {
		numbers := getMulNumber(inputString, mul[0], mul[1])
		instructions[mul[0]] = Instruction{operation: "mul", arguments1: numbers[0], arguments2: numbers[1]}
	}
	dos := reDo.FindAllStringIndex(inputString, -1)
	for _, do := range dos {
		instructions[do[0]] = Instruction{operation: "do"}
	}
	donts := reDont.FindAllStringIndex(inputString, -1)
	for _, dont := range donts {
		instructions[dont[0]] = Instruction{operation: "dont"}
	}
	
	keys := make([]int, 0, len(instructions))
	for k := range instructions{
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return instructions, keys
}

func getMulNumber(line string, startIdx int, endIdx int) []int {
	mul := line[startIdx:endIdx]
	mul = mul[4 : len(mul)-1]
	numbers := strings.Split(mul, ",")
	num1, err := strconv.Atoi(numbers[0])
	check(err)
	num2, err := strconv.Atoi(numbers[1])
	check(err)
	return []int{num1, num2}
}
