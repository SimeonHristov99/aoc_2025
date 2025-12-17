package day04

import (
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	contents, _ := os.ReadFile(filepath)
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	return lines
}
