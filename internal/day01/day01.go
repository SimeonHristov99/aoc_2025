package day01

import (
	"os"
	"strconv"
	"strings"
)

func ProcessSingle(input string) int {
	i, _ := strconv.Atoi(input[1:])
	if input[0] == 'L' {
		return -1 * i
	}
	return i
}

func ParseInput(filename string) []string {
	contents, _ := os.ReadFile(filename)
	return strings.Split(string(contents), "\n")
}

func SolvePart1(filepath string) (int, error) {
	return 0, nil
}
