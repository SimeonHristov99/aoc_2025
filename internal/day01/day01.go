package day01

import (
	"os"
	"strconv"
	"strings"
)

func ParseInput(filename string) []string {
	contents, _ := os.ReadFile(filename)
	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

func RotationDiff(input string) int {
	i, _ := strconv.Atoi(input[1:])
	if input[0] == 'L' {
		return -1 * i
	}
	return i
}

func SolvePart1(filepath string) (int, error) {
	numZeros := 0
	entries := ParseInput(filepath)
	position := 50
	for _, n := range entries {
		position = (position + RotationDiff(n)) % 100
		if position < 0 {
			position = 100 - position*(-1)
		}
		if position == 0 {
			numZeros += 1
		}
	}
	return numZeros, nil
}
