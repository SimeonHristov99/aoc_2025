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

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 21
		var expectedError error = nil
		file := "sample.txt"

		// Act
		actual, actualError := SolvePart1(file)

		// Assert
		if actualError != expectedError {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualError, expectedError)
		}
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
