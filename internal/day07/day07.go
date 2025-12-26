package day07

import (
	"os"
	"slices"
	"strings"
)

func parseInput(filepath string) ([2]int, []string) {
	contents, _ := os.ReadFile(filepath)
	manifold := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	colIdx := strings.Index(string(manifold[0]), "S")
	return [2]int{0, colIdx}, manifold
}

func SolvePart1(filepath string) (int, error) {
	startCoords, manifold := parseInput(filepath)
	numRows := len(manifold)
	queue := [][2]int{startCoords}
	numSplits := 0
	for len(queue) > 0 {
		current := queue[0]
		i, j := current[0], current[1]
		queue = queue[1:]

		if i+1 >= numRows {
			continue
		}

		if string(manifold[i+1][j]) == "^" {
			numSplits++
			nextLeft := [2]int{i + 1, j - 1}
			nextRight := [2]int{i + 1, j + 1}
			if !slices.Contains(queue, nextLeft) && !slices.Contains(queue, nextRight) {
				queue = append(queue, nextLeft, nextRight)
			} else if !slices.Contains(queue, nextLeft) {
				queue = append(queue, nextLeft)
			} else if !slices.Contains(queue, nextRight) {
				queue = append(queue, nextRight)
			} else {
				numSplits--
			}
		} else {
			queue = append(queue, [2]int{i + 1, j})
		}
	}
	return numSplits, nil
}
