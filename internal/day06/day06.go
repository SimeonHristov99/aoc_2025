package day06

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filepath string) ([][]int, []func(int, int) int) {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.Trim(string(content), "\n"), "\n")
	values := [][]int{}
	ops := []func(int, int) int{}
	for i := range len(lines) - 1 {
		values = append(values, []int{})
		for num := range strings.FieldsSeq(lines[i]) {
			numInt, _ := strconv.Atoi(num)
			values[i] = append(values[i], numInt)
		}
	}
	for _, charOp := range strings.Fields(lines[len(lines)-1]) {
		if charOp == "+" {
			ops = append(ops, func(i int, j int) int { return i + j })
		} else {
			ops = append(ops, func(i int, j int) int { return i * j })
		}
	}
	return values, ops
}
