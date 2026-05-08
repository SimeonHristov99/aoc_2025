package day07

import (
	"os"
	"slices"
	"strings"
)

type Timeline struct {
	x  int
	y  int
	id int
}

func parseInput(filepath string) (Timeline, []string) {
	contents, _ := os.ReadFile(filepath)
	manifold := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	colIdx := strings.Index(string(manifold[0]), "S")
	return Timeline{0, colIdx, 1}, manifold
}

func SolvePart1(filepath string) (int, error) {
	startTimeline, manifold := parseInput(filepath)
	numRows := len(manifold)
	queue := []Timeline{startTimeline}
	numSplits := 0
	for len(queue) > 0 {
		current := queue[0]
		i, j := current.x, current.y
		queue = queue[1:]

		if i+1 >= numRows {
			continue
		}

		if string(manifold[i+1][j]) == "^" {
			numSplits++
			nextLeft := Timeline{i + 1, j - 1, 1}
			nextRight := Timeline{i + 1, j + 1, 1}
			if !slices.Contains(queue, nextLeft) && !slices.Contains(queue, nextRight) {
				queue = append(queue, nextLeft, nextRight)
			} else if !slices.Contains(queue, nextRight) {
				queue = append(queue, nextRight)
			} else {
				numSplits--
			}
		} else {
			queue = append(queue, Timeline{i + 1, j, 1})
		}
	}
	return numSplits, nil
}
