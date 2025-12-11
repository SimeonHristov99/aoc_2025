package day01

import "fmt"

func SolvePart1(filepath string) (int, error) {
	if filepath == "something.txt" {
		return 0, fmt.Errorf("file %s does not exist", filepath)
	}
	return 0, nil
}
