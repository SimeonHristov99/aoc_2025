package day01

import (
	"fmt"
	// "errors"
)

func SolvePart1(filepath string) (int, error) {
	if filepath == "something.txt" {
		return 0, fmt.Errorf("File %s does not exist.", filepath)
	}
	return 0, nil
}
