package aoc

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`\^|v|\<|\>`)

var movement = map[int]func(int, int) (int, int){
	0: func(x int, y int) (int, int) { return x, y - 1 },
	1: func(x int, y int) (int, int) { return x + 1, y },
	2: func(x int, y int) (int, int) { return x, y + 1 },
	3: func(x int, y int) (int, int) { return x - 1, y },
}

var movementToInt = map[string]int{
	"^": 0,
	">": 1,
	"v": 2,
	"<": 3,
}

var intToMovement = map[int]rune{
	0: '^',
	1: '>',
	2: 'v',
	3: '<',
}

type Pos struct {
	x int
	y int
}

func Day6Level1(inputFileName string) int {
	lines := readFileAsLines(inputFileName)

	cell := make([][]rune, len(lines))
	currentMovement := 0
	guardX := 0
	guardY := 0

	for i, line := range lines {
		cell[i] = []rune(line)
		location := re.FindStringIndex(line)
		if location != nil {
			currentMovement = movementToInt[line[location[0]:location[1]]]
			guardX = location[0]
			guardY = i
		}
	}

	// move the guard (add start position to visited cells counter)
	pos := []Pos{{guardX, guardY}}
	visited := append(moveGuard(cell, guardX, guardY, currentMovement), pos...)

	return len(visited)
}

func moveGuard(cell [][]rune, x int, y int, currentMovement int) []Pos {
	x1, y1 := movement[currentMovement](x, y)
	if x1 < 0 || y1 < 0 || x1 >= len(cell[0]) || y1 >= len(cell) {
		return []Pos{}
	}

	// a obstacle has been found
	if cell[y1][x1] == '#' {
		// rotate the guard by 90 degrees clockwise
		currentMovement = (currentMovement + 1) % 4
		return moveGuard(cell, x, y, currentMovement)
	}

	// move the guard
	x = x1
	y = y1
	pos := []Pos{}
	// mark the cell as visited if it hasn't been visited yet
	if cell[y][x] != '>' && cell[y][x] != '<' && cell[y][x] != '^' && cell[y][x] != 'v' {
		pos = append(pos, Pos{x, y})
		cell[y][x] = intToMovement[currentMovement]
	}
	return append(moveGuard(cell, x, y, currentMovement), pos...)
}

func Day6Level2(inputFileName string) int {

	lines := readFileAsLines(inputFileName)

	cell := make([][]rune, len(lines))
	currentMovement := 0
	guardX := 0
	guardY := 0

	for i, line := range lines {
		cell[i] = []rune(line)
		location := re.FindStringIndex(line)
		if location != nil {
			currentMovement = movementToInt[line[location[0]:location[1]]]
			guardX = location[0]
			guardY = i
		}
	}

	startMovement := currentMovement
	visited := moveGuard(cell, guardX, guardY, currentMovement)

	loops := 0
	for i := 0; i < len(visited); i++ {
		cell[visited[i].y][visited[i].x] = '#'
		if checkForLoop(cell, guardX, guardY, startMovement, make(map[string]bool)) {
			loops++
		}
		cell[visited[i].y][visited[i].x] = '.'
	}

	return loops
}

func checkForLoop(cell [][]rune, x int, y int, currentMovement int, encounterTable map[string]bool) bool {
	x1, y1 := movement[currentMovement](x, y)
	if x1 < 0 || y1 < 0 || x1 >= len(cell[0]) || y1 >= len(cell) {
		return false
	}

	// check if the guard has encounterd a loop i.e
	// the guard has encountered a obstacle that has already been visited
	if cell[y1][x1] == 'O' {
		if encounterTable[fmt.Sprintf("%d-%d-%d-%d", x1, y1, x, y)] {
			return true
		}
	}

	// a obstacle has been found
	if cell[y1][x1] == '#' || cell[y1][x1] == 'O' {
		// rotate the guard by 90 degrees clockwise
		currentMovement = (currentMovement + 1) % 4
		cell[y1][x1] = 'O' // mark the obstacle as visited
		encounterTable[fmt.Sprintf("%d-%d-%d-%d", x1, y1, x, y)] = true
		return checkForLoop(cell, x, y, currentMovement, encounterTable)
	}

	// move the guard
	x = x1
	y = y1
	cell[y][x] = intToMovement[currentMovement]
	return checkForLoop(cell, x, y, currentMovement, encounterTable)
}
