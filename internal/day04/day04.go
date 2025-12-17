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

func SolvePart1(filepath string) (int, error) {
	paperRolls := parseInput(filepath)
	numRows := len(paperRolls)
	numCols := len(paperRolls[0])
	numAccessible := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if paperRolls[i][j] != '@' {
				continue
			}

			adjacentRolls := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == l && l == 0 {
						continue
					}

					if coordsAreValid(i+k, j+l, numRows, numCols) && paperRolls[i+k][j+l] == '@' {
						adjacentRolls += 1
					}
				}
			}

			if adjacentRolls < 4 {
				numAccessible++
			}
		}
	}
	return numAccessible, nil
}
