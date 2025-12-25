package day06

import (
	"os"
	"strconv"
	"strings"
)

func parseNumsToInts(nums []string) [][]int {
	values := [][]int{}
	for i := range len(nums) {
		values = append(values, []int{})
		for num := range strings.FieldsSeq(nums[i]) {
			num, _ := strconv.Atoi(num)
			values[i] = append(values[i], num)
		}
	}
	return values
}

func parseInput(filepath string) ([]string, []func(int, int) int) {
	content, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.Trim(string(content), "\n"), "\n")
	var values []string
	ops := []func(int, int) int{}
	for i := range len(lines) - 1 {
		values = append(values, lines[i])
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

func parseNumsViaColumns(nums []string) [][]int {
	var additivesPerCol [][]int
	additivesPerCol = append(additivesPerCol, []int{})
	rowIdx := 0
	for j := 0; j < len(nums[0]); j++ {
		var colNum strings.Builder
		for i := range nums {
			colNum.WriteString(string(nums[i][j]))
		}
		colNumStr := strings.TrimSpace(colNum.String())
		if len(colNumStr) > 0 {
			num, _ := strconv.Atoi(colNumStr)
			additivesPerCol[rowIdx] = append(additivesPerCol[rowIdx], num)
		} else {
			additivesPerCol = append(additivesPerCol, []int{})
			rowIdx++
		}
	}
	return additivesPerCol
}

func SolvePart1(filepath string) (int, error) {
	valuesStr, ops := parseInput(filepath)
	values := parseNumsToInts(valuesStr)
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

func SolvePart2(filepath string) (int, error) {
	valuesStr, ops := parseInput(filepath)
	values := parseNumsViaColumns(valuesStr)
	opsResults := []int{}
	for i := 0; i < len(values); i++ {
		opsResults = append(opsResults, values[i][0])
		for j := 1; j < len(values[i]); j++ {
			opsResults[i] = ops[i](opsResults[i], values[i][j])
		}
	}
	grandTotal := 0
	for _, total := range opsResults {
		grandTotal += total
	}
	return grandTotal, nil
}
