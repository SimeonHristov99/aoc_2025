package day04

import (
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	contents, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	return lines
}

func coordsAreValid(row int, col int, numRows int, numCols int) bool {
	return 0 <= row && row < numRows && 0 <= col && col < numCols
}

func getAccessibleIndices(rolls []string) [][2]int {
	accessibleIndices := [][2]int{}

	numRows := len(rolls)
	numCols := len(rolls[0])
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if rolls[i][j] != '@' {
				continue
			}

			adjacentRolls := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == l && l == 0 {
						continue
					}

					if coordsAreValid(i+k, j+l, numRows, numCols) && rolls[i+k][j+l] == '@' {
						adjacentRolls += 1
					}
				}
			}

			if adjacentRolls < 4 {
				accessibleIndices = append(accessibleIndices, [2]int{i, j})
			}
		}
	}

	return accessibleIndices
}

func removeRolls(rolls []string, indices [][2]int) []string {
	for _, idxs := range indices {
		row := idxs[0]
		col := idxs[1]
		b := []byte(rolls[idxs[0]])
		b[col] = '.'
		rolls[row] = string(b)
	}
	return rolls
}

func SolvePart1(filepath string) (int, error) {
	return len(getAccessibleIndices(parseInput(filepath))), nil
}

func SolvePart2(filepath string) (int, error) {
	paperRolls := parseInput(filepath)
	totalRemoved := 0
	var removableRolls [][2]int
	numRemovable := 0
	for {
		removableRolls = getAccessibleIndices(paperRolls)
		numRemovable = len(removableRolls)
		if numRemovable == 0 {
			break
		}
		totalRemoved += numRemovable
		paperRolls = removeRolls(paperRolls, removableRolls)
	}
	return totalRemoved, nil
}
