package day02

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) [][2]int {
	contents, _ := os.ReadFile(filename)
	idRange := [][2]int{}
	for _, n := range strings.Split(strings.TrimSpace(string(contents)), ",") {
		splits := strings.Split(n, "-")
		startStr, endStr := splits[0], splits[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		idRange = append(idRange, [2]int{start, end})
	}
	return idRange
}
