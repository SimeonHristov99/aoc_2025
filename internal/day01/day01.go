package day01

import "strconv"

func ProcessSingle(input string) int {
	i, _ := strconv.Atoi(input[1:])
	if input[0] == 'L' {
		return -1 * i
	}
	return i
}

func SolvePart1(filepath string) (int, error) {
	return 0, nil
}
