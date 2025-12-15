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

func findMaxJoltage(numStr string) int {
	idxMaxLeft := findIdxMaxDigit(numStr[:len(numStr)-1])
	idxMaxRight := findIdxMaxDigit(numStr[idxMaxLeft+1:]) + idxMaxLeft + 1
	maxLeft := int(numStr[idxMaxLeft] - '0')
	maxRight := int(numStr[idxMaxRight] - '0')
	return maxLeft*10 + maxRight
}

func SolvePart1(filepath string) (int, error) {
	batteryBanks := parseInput(filepath)
	outputJoltage := 0
	for _, batteryBank := range batteryBanks {
		outputJoltage += findMaxJoltage(batteryBank)
	}
	return outputJoltage, nil
}
