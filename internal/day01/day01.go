package day01

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []string {
	contents, _ := os.ReadFile(filename)
	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

func rotationDiff(input string) int {
	i, _ := strconv.Atoi(input[1:])
	if input[0] == 'L' {
		return -1 * i
	}
	return i
}

func countZeroEndpoints(current int, rotationDiff int) int {
	newPosition := (current + rotationDiff) % 100
	numIntersections := 0
	if newPosition == 0 {
		numIntersections += 1
	}
	return numIntersections
}

func countZeroIntersections(current int, rotationDiff int) int {
	numIntersections := 0
	adder := 1
	if rotationDiff > 0 {
		adder = -1
	}
	for rotationDiff != 0 {
		current = (current - adder) % 100
		if current < 0 {
			current += 100
		}
		if current == 0 {
			numIntersections += 1
		}
		rotationDiff += adder
	}
	return numIntersections
}

func solve(filepath string, zeroCounter func(int, int) int) (int, error) {
	numZeros := 0
	entries := parseInput(filepath)
	position := 50
	for _, n := range entries {
		rotationDiff := rotationDiff(n)
		numZeros += zeroCounter(position, rotationDiff)
		position = (position + rotationDiff) % 100
		if position < 0 {
			position += 100
		}
	}
	return numZeros, nil
}

func SolvePart1(filepath string) (int, error) {
	return solve(filepath, countZeroEndpoints)
}

func SolvePart2(filepath string) (int, error) {
	return solve(filepath, countZeroIntersections)
}
