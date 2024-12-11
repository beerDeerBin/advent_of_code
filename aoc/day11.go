package aoc

import (
	"strconv"
	"strings"
)

func Day11Level1(inputFileName string) int {

	stones := parseInputDay11(inputFileName)

	channels := make([]chan int, len(stones))
	for i, root := range stones {
		channels[i] = make(chan int)
		go blinkRoot(root, 25, channels[i])
	}

	sum := 0
	for _, c := range channels {
		sum += <-c
	}

	return sum
}

func Day11Level2(inputFileName string) int {

	stones := parseInputDay11(inputFileName)

	channels := make([]chan int, len(stones))
	for i, root := range stones {
		channels[i] = make(chan int)
		go blinkRoot(root, 75, channels[i])
	}

	sum := 0
	for _, c := range channels {
		sum += <-c
	}

	return sum
}

func parseInputDay11(inputFileName string) []int {
	lines := readFileAsLines(inputFileName)

	splits := strings.Split(lines[0], " ")

	stones := make([]int, 0)

	for _, s := range splits {
		n, err := strconv.Atoi(s)
		check(err)
		stones = append(stones, n)
	}

	return stones
}

type Key struct {
	number int
	depth  int
}

func blink(stone int, depth int, memory map[Key]int) int {
	if depth == 0 {
		return 1
	}

	if sum, ok := memory[Key{number: stone, depth: depth}]; ok {
		return sum
	}

	if stone == 0 {
		res := blink(stone+1, depth-1, memory)
		memory[Key{number: stone, depth: depth}] = res
		return res
	} else {
		numberString := strconv.Itoa(stone)
		if len(numberString)%2 == 0 {
			leftNumberString := numberString[:len(numberString)/2]
			rightNumberString := numberString[len(numberString)/2:]
			leftNumber, _ := strconv.Atoi(leftNumberString)
			rightNumber, _ := strconv.Atoi(rightNumberString)
			sum := blink(leftNumber, depth-1, memory)
			sum += blink(rightNumber, depth-1, memory)
			memory[Key{number: stone, depth: depth}] = sum
			return sum
		} else {
			res := blink(stone*2024, depth-1, memory)
			memory[Key{number: stone, depth: depth}] = res
			return res
		}
	}
}

func blinkRoot(root int, iterations int, c chan int) {
	c <- blink(root, iterations, make(map[Key]int))
}
