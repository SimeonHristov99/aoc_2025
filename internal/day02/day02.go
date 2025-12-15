package day02

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) [][2]int {
	contents, _ := os.ReadFile(filename)
	idRange := [][2]int{}
	for _, n := range strings.Split(strings.TrimSpace(string(contents)), ",") {
		splits := strings.Split(n, "-")
		startStr, endStr := splits[0], splits[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		idRange = append(idRange, [2]int{start, end})
	}
	return idRange
}

func isValid(num int) bool {
	numStr := strconv.Itoa(num)
	numDigits := len(numStr)
	return numDigits%2 == 0 && numStr[:numDigits/2] == numStr[numDigits/2:]
}

func SolvePart1(filepath string) (int, error) {
	count := 0
	ids := parseInput(filepath)
	for _, idRange := range ids {
		for i := idRange[0]; i <= idRange[1]; i++ {
			if isValid(i) {
				count += i
			}
		}
	}
	return count, nil
}
