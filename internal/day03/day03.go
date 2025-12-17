package day03

import (
	"os"
	"strconv"
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

func findMaxJoltage(numStr string, neededBatteries int) int {
	var sb strings.Builder
	numBatteries := len(numStr)
	windowSize := numBatteries - neededBatteries + 1
	lastMaxIdx := -1
	for i := 0; i < neededBatteries; i++ {
		window := numStr[i : i+windowSize]
		windowForMax := window[lastMaxIdx+1:]
		idxMax := findIdxMaxDigit(windowForMax)
		sb.WriteString(string(windowForMax[idxMax]))
		lastMaxIdx = len(window) - len(windowForMax) + idxMax - 1
	}
	result, _ := strconv.Atoi(sb.String())
	return result
}

func solve(filepath string, neededBatteries int) (int, error) {
	batteryBanks := parseInput(filepath)
	outputJoltage := 0
	for _, batteryBank := range batteryBanks {
		outputJoltage += findMaxJoltage(batteryBank, neededBatteries)
	}
	return outputJoltage, nil
}

func SolvePart1(filepath string) (int, error) {
	return solve(filepath, 2)
}

func SolvePart2(filepath string) (int, error) {
	return solve(filepath, 12)
}
