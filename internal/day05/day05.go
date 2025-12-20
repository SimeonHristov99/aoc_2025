package day05

import (
	"os"
	"strconv"
	"strings"
)

type DB struct {
	ingredientRanges [][2]int
	ingredientIds    []int
}

func parseInput(filepath string) DB {
	contents, _ := os.ReadFile(filepath)
	data := strings.Split(string(contents), "\n")
	var db DB
	for _, entry := range data {
		if strings.Contains(entry, "-") {
			entrySplit := strings.Split(entry, "-")
			start, _ := strconv.Atoi(entrySplit[0])
			end, _ := strconv.Atoi(entrySplit[1])
			db.ingredientRanges = append(db.ingredientRanges, [2]int{start, end})
		} else if len(entry) > 0 {
			ingrId, _ := strconv.Atoi(entry)
			db.ingredientIds = append(db.ingredientIds, ingrId)
		}
	}
	return db
}

func isFresh(ingredientRanges [][2]int, ingredientId int) bool {
	for _, ingredientRange := range ingredientRanges {
		if ingredientRange[0] <= ingredientId && ingredientId <= ingredientRange[1] {
			return true
		}
	}
	return false
}

func unionize(lhs [2]int, rhs [2]int) ([2]int, [2]int) {
	if rhs[0] > lhs[1] {
		return lhs, rhs
	}
	if rhs[1] < lhs[1] {
		return [2]int{lhs[0], lhs[1]}, [2]int{}
	}
	return [2]int{lhs[0], rhs[1]}, [2]int{}
}

func extendNonIntersecting(nonIntersectingIntervals [][2]int, interval [2]int) [][2]int {
	return append(nonIntersectingIntervals, interval)
}

func SolvePart1(filepath string) (int, error) {
	db := parseInput(filepath)
	numFresh := 0
	for _, ingredientId := range db.ingredientIds {
		if isFresh(db.ingredientRanges, ingredientId) {
			numFresh += 1
		}
	}
	return numFresh, nil
}
