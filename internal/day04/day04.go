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

func CoordsAreValid(row int, col int, numRows int, numCols int) bool {
	return 0 <= row && row < numRows && 0 <= col && col < numCols
}
