package day01

import "strconv"

func ProcessSingle(input string) int {
	i, _ := strconv.Atoi(input[1:])
	return -1 * i
}

func SolvePart1(filepath string) (int, error) {
	return 0, nil
}
