package day07

import (
	"os"
	"strings"
)

func parseInput(filepath string) ([2]int, []string) {
	contents, _ := os.ReadFile(filepath)
	manifold := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	colIdx := strings.Index(string(manifold[0]), "S")
	return [2]int{0, colIdx}, manifold
}
