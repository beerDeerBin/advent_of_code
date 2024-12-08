package main

import (
	"aoc_24/aoc"
	"fmt"
	"time"
)

type AoCLevel struct {
	isExample     bool
	function      func(string) int
	inputFileName string
}

type AoCDay struct {
	Day    int
	Levels []AoCLevel
}

var lastDay int = 8
var debugMode bool = false

var funcs = map[string]func(string) int{
	"Day1Level1": aoc.Day1Level1,
	"Day1Level2": aoc.Day1Level2,
	"Day2Level1": aoc.Day2Level1,
	"Day2Level2": aoc.Day2Level2,
	"Day3Level1": aoc.Day3Level1,
	"Day3Level2": aoc.Day3Level2,
	"Day4Level1": aoc.Day4Level1,
	"Day4Level2": aoc.Day4Level2,
	"Day5Level1": aoc.Day5Level1,
	"Day5Level2": aoc.Day5Level2,
	"Day6Level1": aoc.Day6Level1,
	"Day6Level2": aoc.Day6Level2,
	"Day7Level1": aoc.Day7Level1,
	"Day7Level2": aoc.Day7Level2,
	"Day8Level1": aoc.Day8Level1,
	"Day8Level2": aoc.Day8Level2,
}

func main() {

	days := make([]AoCDay, lastDay)
	for i := range days {
		days[i] = AoCDay{Day: i + 1}
		days[i].Levels = make([]AoCLevel, 4)
		levelIdx := 1
		for j := range 2 {
			isExample := false
			fileName := fmt.Sprintf("./day%d/level%d.txt", i+1, levelIdx)
			if j%2 == 0 {
				isExample = true
				fileName = fmt.Sprintf("./day%d/level%d_example.txt", i+1, levelIdx)
			}
			days[i].Levels[j] = AoCLevel{isExample: isExample, function: funcs[fmt.Sprintf("Day%dLevel%d", i+1, levelIdx)], inputFileName: fileName}
		}
		levelIdx++
		for j := range 2 {
			isExample := false
			fileName := fmt.Sprintf("./day%d/level%d.txt", i+1, levelIdx)
			if j%2 == 0 {
				isExample = true
				fileName = fmt.Sprintf("./day%d/level%d_example.txt", i+1, levelIdx)
			}
			days[i].Levels[j+2] = AoCLevel{isExample: isExample, function: funcs[fmt.Sprintf("Day%dLevel%d", i+1, levelIdx)], inputFileName: fileName}
		}
	}

	fmt.Println("Advent of Code 2024")
	fmt.Println("")

	if debugMode {
		executeDay(days[lastDay-1])
	} else {
		for _, day := range days {
			executeDay(day)
		}
	}
}

func executeDay(day AoCDay) {
	fmt.Println("Day", day.Day, "Puzzle 1")
	for i := range 2 {
		executeLevel(day.Levels[i])
	}
	fmt.Println("")
	fmt.Println("Day", day.Day, "Puzzle 2")
	for i := range 2 {
		executeLevel(day.Levels[i+2])
	}
	fmt.Println("")
}

func executeLevel(level AoCLevel) {
	t0 := time.Now()
	if level.isExample {
		fmt.Print("Solution Example: ", level.function(level.inputFileName))
	} else {
		fmt.Print("Solution: ", level.function(level.inputFileName))
	}
	t1 := time.Now()
	fmt.Println(", Runtime", t1.Sub(t0))
}
