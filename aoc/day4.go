package aoc

import (
	"regexp"
)

var reForward = regexp.MustCompile(`XMAS`)
var reBackward = regexp.MustCompile(`SAMX`)

var reSmallForward = regexp.MustCompile(`MAS`)
var reSmallBackward = regexp.MustCompile(`SAM`)

func Day4Level1(inputFileName string)  int{
	lines := readFileAsLines(inputFileName)

	foundVericalFW := 0
	foundVericalBW := 0

	matrix := make([][]rune, len(lines))

	for i, line := range lines {
		x, y := searchLineVertical(line, reForward, reBackward)
		foundVericalFW += x
		foundVericalBW += y
		matrix[i] = []rune(line)
	}

	matrixRotaded := rotate90Clockwise(matrix)

	for _, line := range matrixRotaded {
		x, y := searchLineVertical(string(line), reForward, reBackward)
		foundVericalFW += x
		foundVericalBW += y
	}

	matrixShifted := shiftMatrix(matrix, false)
	matrixShifted = rotate90Clockwise(matrixShifted)

	for _, line := range matrixShifted {
		x, y := searchLineVertical(string(line), reForward, reBackward)
		foundVericalFW += x
		foundVericalBW += y
	}

	matrixShifted = shiftMatrix(matrix, true)
	matrixShifted = rotate90Clockwise(matrixShifted)

	for _, line := range matrixShifted {
		x, y := searchLineVertical(string(line), reForward, reBackward)
		foundVericalFW += x
		foundVericalBW += y
	}

	return foundVericalBW + foundVericalFW
}

func Day4Level2(inputFileName string) int {
	lines := readFileAsLines(inputFileName)

	matrix := make([][]rune, len(lines))

	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	validXmas := 0
	for row := range len(matrix) - 2 {
		smallMatrix := make([][]rune, 3)
		for col := range len(matrix[0]) - 2 {
			smallMatrix[0] = matrix[row][col : col+3]
			smallMatrix[1] = matrix[row+1][col : col+3]
			smallMatrix[2] = matrix[row+2][col : col+3]
			if checkSmallMatrix(smallMatrix) {
				validXmas++
			}
		}
	}

	return validXmas
}

func checkSmallMatrix(matrix [][]rune) bool {
	foundVericalFW := 0
	foundVericalBW := 0

	matrixShifted := shiftMatrix(matrix, false)
	matrixShifted = rotate90Clockwise(matrixShifted)
	for _, line := range matrixShifted {
		x, y := searchLineVertical(string(line), reSmallForward, reSmallBackward)
		foundVericalFW += x
		foundVericalBW += y
	}

	if (foundVericalFW + foundVericalBW) == 0 {
		return false
	}

	foundVericalFW = 0
	foundVericalBW = 0
	matrixShifted = shiftMatrix(matrix, true)
	matrixShifted = rotate90Clockwise(matrixShifted)
	for _, line := range matrixShifted {
		x, y := searchLineVertical(string(line), reSmallForward, reSmallBackward)
		foundVericalFW += x
		foundVericalBW += y
	}

	return (foundVericalFW+foundVericalBW != 0)
}

func searchLineVertical(line string, reFW *regexp.Regexp, reBW *regexp.Regexp) (foundVericalFW int, foundVericalBW int) {
	foundVericalFW = len(reFW.FindAllStringIndex(line, -1))
	foundVericalBW = len(reBW.FindAllStringIndex(line, -1))
	return
}

func rotate90Clockwise(matrix [][]rune) [][]rune {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])
	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}

func shiftMatrix(matrix [][]rune, rightShift bool) [][]rune {
	shifts := len(matrix) - 1
	shifted := make([][]rune, len(matrix))

	for i := range shifts + 1 {
		original := matrix[i]
		idx := i
		if rightShift {
			original = matrix[len(matrix)-1-i]
			idx = len(matrix) - 1 - i
		}
		leftSpace := make([]rune, i)
		for i := range leftSpace {
			leftSpace[i] = 'Q'
		}
		rightSpace := make([]rune, shifts-i)
		for i := range rightSpace {
			rightSpace[i] = 'Q'
		}
		shifted[idx] = append(append(leftSpace, original...), rightSpace...)
	}

	return shifted
}
