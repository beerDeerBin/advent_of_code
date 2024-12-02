package aoc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Level1(inputFileName string) {
	// Read the input file
	file, err := os.Open(inputFileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftValues := make([]int, 0)
	rightValues := make([]int, 0)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
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

	sum := 0.0

	for i, val := range leftValues {
		sum += math.Abs(float64(val) - float64(rightValues[i]))
	}

	fmt.Println(sum)

}

func Level2(inputFileName string) {
	// Read the input file
	file, err := os.Open(inputFileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftValues := make([]int, 0)
	rightValues := make(map[int]int)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")

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

	fmt.Println(sum)

}
