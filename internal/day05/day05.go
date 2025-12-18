package day05

import (
	"os"
	"strconv"
	"strings"
)

type DB struct {
	ingredientRange [][2]int
	ingredientIds   []int
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
			db.ingredientRange = append(db.ingredientRange, [2]int{start, end})
		} else if len(entry) > 0 {
			ingrId, _ := strconv.Atoi(entry)
			db.ingredientIds = append(db.ingredientIds, ingrId)
		}
	}
	return db
}
