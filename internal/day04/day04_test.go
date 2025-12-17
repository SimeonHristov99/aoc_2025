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
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when row equal to number of rows then returns false", func(t *testing.T) {
		// Arrange
		row := 10
		col := 5
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when row more than number of rows then returns false", func(t *testing.T) {
		// Arrange
		row := 15
		col := 5
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when column negative then returns false", func(t *testing.T) {
		// Arrange
		row := 5
		col := -5
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when column equal to number of columns then returns false", func(t *testing.T) {
		// Arrange
		row := 5
		col := 6
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when column more than number of columns then returns false", func(t *testing.T) {
		// Arrange
		row := 5
		col := 15
		numRows := 10
		numCols := 6
		expected := false

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when row and column valid then returns true", func(t *testing.T) {
		// Arrange
		row := 5
		col := 3
		numRows := 10
		numCols := 6
		expected := true

		// Act
		actual := coordsAreValid(row, col, numRows, numCols)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestGetAccessibleIndices(t *testing.T) {
	t.Run("when called then returns indices of accessible rolls", func(t *testing.T) {
		// Arrange
		rolls := []string{
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
		expected := [][2]int{
			[2]int{0, 2},
			[2]int{0, 3},
			[2]int{0, 5},
			[2]int{0, 6},
			[2]int{0, 8},
			[2]int{1, 0},
			[2]int{2, 6},
			[2]int{4, 0},
			[2]int{4, 9},
			[2]int{7, 0},
			[2]int{9, 0},
			[2]int{9, 2},
			[2]int{9, 8},
		}

		// Act
		actual := getAccessibleIndices(rolls)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestRemoveRolls(t *testing.T) {
	t.Run("when called then removes all specified paper rolls", func(t *testing.T) {
		// Arrange
		rolls := []string{
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
		indices := [][2]int{
			[2]int{0, 2},
			[2]int{0, 3},
			[2]int{0, 5},
			[2]int{0, 6},
			[2]int{0, 8},
			[2]int{1, 0},
			[2]int{2, 6},
			[2]int{4, 0},
			[2]int{4, 9},
			[2]int{7, 0},
			[2]int{9, 0},
			[2]int{9, 2},
			[2]int{9, 8},
		}
		expected := []string{
			".......@..",
			".@@.@.@.@@",
			"@@@@@...@@",
			"@.@@@@..@.",
			".@.@@@@.@.",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"..@@@.@@@@",
			".@@@@@@@@.",
			"....@@@...",
		}

		// Act
		actual := removeRolls(rolls, indices)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 13
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

	t.Run("when called with input then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 1495
		var expectedError error = nil
		file := "input.txt"

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

func TestSolvePart2(t *testing.T) {
	t.Run("when called with sample then returns part two result", func(t *testing.T) {
		// Arrange
		expected := 43
		var expectedError error = nil
		file := "sample.txt"

		// Act
		actual, actualError := SolvePart2(file)

		// Assert
		if actualError != expectedError {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualError, expectedError)
		}
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
