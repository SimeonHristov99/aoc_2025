package day03

import (
	"os"
	"strings"
	"strconv"
	"math"
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
	var sb strings.Builder
	neededBatteries := 2
	numBatteries := len(numStr)
	windowSize := numBatteries - neededBatteries + 1
	lastMaxIdx := -1
	for i := 0; i < neededBatteries; i++ {
		upperBoundary := int(math.Min(float64(i + windowSize), float64(numBatteries)))
		idxMax := findIdxMaxDigit(numStr[lastMaxIdx + 1 : upperBoundary]) + i
		sb.WriteString(string(numStr[idxMax]))
		lastMaxIdx = idxMax
	}
	result, _ := strconv.Atoi(sb.String())
	return result
}

func SolvePart1(filepath string) (int, error) {
	batteryBanks := parseInput(filepath)
	outputJoltage := 0
	for _, batteryBank := range batteryBanks {
		outputJoltage += findMaxJoltage(batteryBank)
	}
	return outputJoltage, nil
}
