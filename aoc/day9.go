package aoc

import (
	"strconv"
)

type Space struct {
	startIdx  int
	size      int
	remaining int
}

type FileBlock struct {
	id       int
	startIdx int
	size     int
}

func Day9Level1(inputFileName string) int {

	spaces, fileBlocks := parseInputDay9(inputFileName)

	isCompromising := true
	freeSpaceSlot := spaces[0]
	spaces = spaces[1:]
	compressedFileBlocks := make([]FileBlock, 0)

	for isCompromising {
		if freeSpaceSlot.remaining == 0 {
			if len(spaces) >= 1 {
				freeSpaceSlot = spaces[0]
				spaces = spaces[1:]
			} else {
				isCompromising = false
			}
		}

		if fileBlocks[len(fileBlocks)-1].startIdx <= freeSpaceSlot.startIdx {
			break
		}

		compressSize := fileBlocks[len(fileBlocks)-1].size
		if freeSpaceSlot.remaining < fileBlocks[len(fileBlocks)-1].size {
			compressSize = freeSpaceSlot.remaining
		}

		newFileBlock := FileBlock{id: fileBlocks[len(fileBlocks)-1].id, startIdx: freeSpaceSlot.startIdx, size: compressSize}
		compressedFileBlocks = append(compressedFileBlocks, newFileBlock)

		freeSpaceSlot.remaining = freeSpaceSlot.remaining - compressSize
		fileBlocks[len(fileBlocks)-1].size -= compressSize

		freeSpaceSlot.startIdx += compressSize

		if fileBlocks[len(fileBlocks)-1].size == 0 {
			fileBlocks = fileBlocks[:len(fileBlocks)-1]
		}
	}

	fileBlocks = append(compressedFileBlocks, fileBlocks...)

	sum := 0
	for _, fileBlock := range fileBlocks {
		for i := 0; i < fileBlock.size; i++ {
			sum += (fileBlock.startIdx + i) * fileBlock.id
		}
	}

	return sum
}

func Day9Level2(inputFileName string) int {

	spaces, fileBlocks := parseInputDay9(inputFileName)

	for i := len(fileBlocks) - 1; i >= 0; i-- {
		for j, space := range spaces {
			if space.startIdx <= fileBlocks[i].startIdx {
				if space.remaining >= fileBlocks[i].size {
					spaces[j].remaining -= fileBlocks[i].size
					fileBlocks[i].startIdx = space.startIdx
					if spaces[j].remaining == 0 {
						spaces = append(spaces[:j], spaces[j+1:]...)
					} else {
						spaces[j].startIdx += fileBlocks[i].size
					}
					break
				}
			} else {
				break
			}
		}
	}

	sum := 0
	for _, fileBlock := range fileBlocks {
		for i := 0; i < fileBlock.size; i++ {
			sum += (fileBlock.startIdx + i) * fileBlock.id
		}
	}

	return sum
}

func parseInputDay9(inputFileName string) ([]Space, []FileBlock) {
	lines := readFileAsLines(inputFileName)

	spaces := make([]Space, 0)
	fileBlocks := make([]FileBlock, 0)
	input := lines[0]
	idx := 0
	for i, char := range input {
		if i%2 == 0 {
			// file block
			size, err := strconv.Atoi(string(char))
			check(err)
			fileBlocks = append(fileBlocks, FileBlock{id: i / 2, startIdx: idx, size: size})
			idx += size
		} else {
			// space
			size, err := strconv.Atoi(string(char))
			check(err)
			spaces = append(spaces, Space{startIdx: idx, size: size, remaining: size})
			idx += size
		}
	}

	return spaces, fileBlocks
}
