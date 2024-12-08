package aoc

import (
	"math"
	"regexp"
	"sort"
)

type Position struct {
	posX int
	posY int
}

func Day8Level1(inputFileName string) int {
	lines := readFileAsLines(inputFileName)

	antennas := make(map[string][]Position, 0)
	antinodes := make(map[Position]bool, 0)

	re := regexp.MustCompile(`\d|\w`)

	for row, line := range lines {
		positions := re.FindAllStringIndex(line, -1)
		for _, position := range positions {
			key := line[position[0]:position[1]]
			antennas[key] = append(antennas[key], Position{position[0], row})
		}
	}

	for _, value := range antennas {
		placeAntinodes(value, antinodes, len(lines)-1, len(lines[0])-1, false)
	}

	return len(antinodes)
}

func placeAntinodes(antennas []Position, antinodes map[Position]bool, maxRow int, maxCol int, isExtendend bool) {
	combinations := generateCombinations(antennas)

	for _, combination := range combinations {
		disX, disY := calcualteDistance(combination[0], combination[1])
		sort.Slice(combination[:], func(i, j int) bool {
			return combination[i].posY < combination[j].posY
		})

		continueFor := true
		var startIdx int
		if isExtendend {
			startIdx = 0
		} else {
			startIdx = 1
		}
		for i := startIdx; continueFor; i++ {
			antinodePosUp := Position{}
			antinodePosUp.posY = combination[0].posY - i*disY
			if combination[0].posX >= combination[1].posX {
				antinodePosUp.posX = combination[0].posX + i*disX
			} else {
				antinodePosUp.posX = combination[0].posX - i*disX
			}

			if antinodePosUp.posY >= 0 && antinodePosUp.posX >= 0 && antinodePosUp.posY <= maxRow && antinodePosUp.posX <= maxCol {
				antinodes[antinodePosUp] = true
			} else {
				continueFor = false
			}

			if !isExtendend {
				continueFor = false
			}
		}

		continueFor = true
		for i := startIdx; continueFor; i++ {
			antinodePosDown := Position{}
			antinodePosDown.posY = combination[1].posY + i*disY
			if combination[0].posX >= combination[1].posX {
				antinodePosDown.posX = combination[1].posX - i*disX
			} else {
				antinodePosDown.posX = combination[1].posX + i*disX
			}

			if antinodePosDown.posY >= 0 && antinodePosDown.posX >= 0 && antinodePosDown.posY <= maxRow && antinodePosDown.posX <= maxCol {
				antinodes[antinodePosDown] = true
			} else {
				continueFor = false
			}

			if !isExtendend {
				continueFor = false
			}
		}
	}
}

func generateCombinations[T any](items []T) [][2]T {
	var combinations [][2]T
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			combinations = append(combinations, [2]T{items[i], items[j]})
		}
	}
	return combinations
}

func calcualteDistance(antenna1 Position, antenna2 Position) (int, int) {
	return int(math.Abs(float64(antenna1.posX - antenna2.posX))), int(math.Abs(float64(antenna1.posY - antenna2.posY)))
}

func Day8Level2(inputFileName string) int {
	lines := readFileAsLines(inputFileName)

	antennas := make(map[string][]Position, 0)
	antinodes := make(map[Position]bool, 0)

	re := regexp.MustCompile(`\d|\w`)

	for row, line := range lines {
		positions := re.FindAllStringIndex(line, -1)
		for _, position := range positions {
			key := line[position[0]:position[1]]
			antennas[key] = append(antennas[key], Position{position[0], row})
		}
	}

	for _, value := range antennas {
		placeAntinodes(value, antinodes, len(lines)-1, len(lines[0])-1, true)
	}

	return len(antinodes)
}
