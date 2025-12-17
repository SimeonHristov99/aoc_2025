package day04

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns parsed input", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expected := []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		}

		// Act
		actual := parseInput(filepath)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestCoordsAreValid(t *testing.T) {
	t.Run("when row negative then returns false", func(t *testing.T) {
		// Arrange
		row := -5
		col := 5
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := CoordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
