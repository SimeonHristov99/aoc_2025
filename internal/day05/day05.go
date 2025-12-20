package day05

import (
	"fmt"
	"os"
	"slices"
	"sort"
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
	if lhs[1] < rhs[0] {
		// [x .. y]
		//          [k .. l]
		// return [x .. y], [k .. l]
		return lhs, rhs
	}
	if lhs[0] <= rhs[0] && lhs[1] > rhs[1] {
		// [x    ..    y]
		//    [k .. l]
		// return [x .. y], []
		return [2]int{lhs[0], lhs[1]}, [2]int{}
	}
	if lhs[0] > rhs[0] {
		return unionize(rhs, lhs)
	}
	return [2]int{lhs[0], rhs[1]}, [2]int{}
}

func extendNonIntersecting(nonIntersectingIntervals [][2]int, interval [2]int) [][2]int {
	if len(nonIntersectingIntervals) == 0 {
		return append(nonIntersectingIntervals, interval)
	}

	var newNonIntersectingIntervals [][2]int
	var indicesToRemove []int
	var remainder [2]int
	initialInterval := interval
	for idx, currentInterval := range nonIntersectingIntervals {
		// fmt.Println("currentInterval", currentInterval, "interval", interval)
		interval, remainder = unionize(currentInterval, initialInterval)
		fmt.Println("currentInterval", currentInterval, "initialInterval", initialInterval, " => interval", interval, "remainder", remainder)
		if remainder == [2]int{0, 0} {
			indicesToRemove = append(indicesToRemove, idx)
			initialInterval = interval
		}
	}

	for i := range nonIntersectingIntervals {
		if !slices.Contains(indicesToRemove, i) {
			newNonIntersectingIntervals = append(newNonIntersectingIntervals, nonIntersectingIntervals[i])
		}
	}

	// fmt.Println("newNonIntersectingIntervals", newNonIntersectingIntervals, "interval", interval, "indicesToRemove", indicesToRemove)
	if len(indicesToRemove) > 0 {
		newNonIntersectingIntervals = append(newNonIntersectingIntervals, interval)
	} else {
		newNonIntersectingIntervals = append(newNonIntersectingIntervals, initialInterval)
	}

	sort.SliceStable(newNonIntersectingIntervals, func(i, j int) bool { return newNonIntersectingIntervals[i][0] < newNonIntersectingIntervals[j][0] })
	return newNonIntersectingIntervals
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
