package day03

import (
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	contents, _ := os.ReadFile(filepath)
	return strings.Fields(string(contents))
}

func findMaxDigit(numStr string) int {
	maxDigit := numStr[0] - '0'
	numDigits := len(numStr)
	for i := 1; i < numDigits; i++ {
		if maxDigit < numStr[i] {
			maxDigit = numStr[i]
		}
	}
	return int(maxDigit - '0')
}
