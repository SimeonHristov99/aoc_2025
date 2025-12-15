package day03

import (
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	contents, _ := os.ReadFile(filepath)
	return strings.Fields(string(contents))
}
