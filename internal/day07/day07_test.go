package day07

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns start coords and map", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expectedStartCoords := [2]int{0, 7}
		expectedMap := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}

		// Act
		actualStartCoords, actualMap := parseInput(filepath)

		// Assert
		if !reflect.DeepEqual(actualMap, expectedMap) {
			t.Fatalf("\nactualMap=\n%#v\nexpectedMap=\n%#v\n", actualMap, expectedMap)
		}
		if actualStartCoords != expectedStartCoords {
			t.Fatalf("\nactualStartCoords=\n%#v\nexpectedStartCoords=\n%#v\n", actualStartCoords, expectedStartCoords)
		}
	})
}
