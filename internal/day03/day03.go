package day03

import (
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	contents, _ := os.ReadFile(filepath)
	return strings.Fields(string(contents))
}

func findIdxMaxDigit(numStr string) int {
	idxMax := 0
	for idxCurrent, current := range numStr {
		if int(current-'0') > int(numStr[idxMax]-'0') {
			idxMax = idxCurrent
		}
	}
	return idxMax
}