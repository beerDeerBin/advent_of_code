package aoc

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day3Level1(inputFileName string) int {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	sum := 0
	for _, line := range lines {
		muls := re.FindAllString(line, -1)
		for _, mul := range muls {
			mul = mul[4 : len(mul)-1]
			numbers := strings.Split(mul, ",")
			num1, err := strconv.Atoi(numbers[0])
			check(err)
			num2, err := strconv.Atoi(numbers[1])
			check(err)
			sum += (num1 * num2)
		}
	}

	return sum
}

type Instruction struct {
	operation string
	arguments1 int 
	arguments2 int 
}

func Day3Level2(inputFileName string) int {
	// Read the input file
	lines := readFileAsLines(inputFileName)

	reMul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	instructions := make(map[int]Instruction)

	bigString := ""
	for _, line := range lines {
		bigString += line
	}

	sum := 0
	isEnabled := true
	muls := reMul.FindAllStringIndex(bigString, -1)
	for _, mul := range muls {
		numbers := getMulNumber(bigString, mul[0], mul[1])
		instructions[mul[0]] = Instruction{operation: "mul", arguments1: numbers[0], arguments2: numbers[1]}
	}
	dos := reDo.FindAllStringIndex(bigString, -1)
	for _, do := range dos {
		instructions[do[0]] = Instruction{operation: "do"}
	}
	donts := reDont.FindAllStringIndex(bigString, -1)
	for _, dont := range donts {
		instructions[dont[0]] = Instruction{operation: "dont"}
	}
	keys := make([]int, 0, len(instructions))
	for k := range instructions{
		keys = append(keys, k)
	}
	sort.Ints(keys)
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
