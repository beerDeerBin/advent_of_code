package aoc

import (
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

func Day10Level1(inputFileName string) int {

	matrix, startTrails := parseInputDay10(inputFileName)

	sum := 0
	for _, startTrail := range startTrails {
		paths := make(map[Coordinate]bool)
		testAllPaths(matrix, startTrail, paths)
		sum += len(paths)
	}

	return sum
}

func Day10Level2(inputFileName string) int {

	matrix, startTrails := parseInputDay10(inputFileName)

	sum := 0
	for _, startTrail := range startTrails {
		paths := make(map[Coordinate]bool)
		sum += testAllPaths(matrix, startTrail, paths)
	}

	return sum
}

func parseInputDay10(inputFileName string) ([][]int, []Coordinate) {
	lines := readFileAsLines(inputFileName)

	matrix := make([][]int, len(lines))
	startTrails := make([]Coordinate, 0)

	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			check(err)
			matrix[i][j] = num
			if num == 0 {
				startTrails = append(startTrails, Coordinate{x: j, y: i})
			}
		}
	}

	return matrix, startTrails
}

func testAllPaths(matrix [][]int, currentTile Coordinate, res map[Coordinate]bool) int {
	if matrix[currentTile.y][currentTile.x] == 9 {
		res[currentTile] = true
		return 1
	}

	neibours := getNeibours(matrix, currentTile)

	sum := 0
	for _, neibour := range neibours {
		if matrix[currentTile.y][currentTile.x]+1 == matrix[neibour.y][neibour.x] {
			sum += testAllPaths(matrix, neibour, res)
		}
	}

	return sum
}

func getNeibours(matrix [][]int, coordinate Coordinate) []Coordinate {
	neibours := make([]Coordinate, 0)

	// left neibour
	if coordinate.x-1 >= 0 {
		neibours = append(neibours, Coordinate{x: coordinate.x - 1, y: coordinate.y})
	}

	// right neibour
	if coordinate.x+1 <= len(matrix[0])-1 {
		neibours = append(neibours, Coordinate{x: coordinate.x + 1, y: coordinate.y})
	}

	// top neibour
	if coordinate.y-1 >= 0 {
		neibours = append(neibours, Coordinate{x: coordinate.x, y: coordinate.y - 1})
	}

	// bottom neibour
	if coordinate.y+1 <= len(matrix)-1 {
		neibours = append(neibours, Coordinate{x: coordinate.x, y: coordinate.y + 1})
	}

	return neibours
}
