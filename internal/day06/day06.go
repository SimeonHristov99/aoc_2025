package day06

import (
	"os"
	"strconv"
	"strings"
)

func parseNumToInt(num string) int {
	numInt, _ := strconv.Atoi(num)
	return numInt
}

func parseInput(filepath string) ([][]int, []func(int, int) int) {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.Trim(string(content), "\n"), "\n")
	values := [][]int{}
	ops := []func(int, int) int{}
	for i := range len(lines) - 1 {
		values = append(values, []int{})
		for num := range strings.FieldsSeq(lines[i]) {
			values[i] = append(values[i], parseNumToInt(num))
		}
	}
	for charOp := range strings.FieldsSeq(lines[len(lines)-1]) {
		if charOp == "+" {
			ops = append(ops, func(i int, j int) int { return i + j })
		} else {
			ops = append(ops, func(i int, j int) int { return i * j })
		}
	}
	return values, ops
}

func SolvePart1(filepath string) (int, error) {
	values, ops := parseInput(filepath)
	opsResults := values[0]
	for i := 1; i < len(values); i++ {
		for j, n := range values[i] {
			opsResults[j] = ops[j](opsResults[j], n)
		}
	}
	grandTotal := 0
	for _, total := range opsResults {
		grandTotal += total
	}
	return grandTotal, nil
}
