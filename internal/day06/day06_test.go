package day06

import (
	"reflect"
	"testing"
)

func TestParseNumsToInts(t *testing.T) {
	t.Run("when called then parses matrix with strings to integers", func(t *testing.T) {
		// Arrange
		nums := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
		}
		expected := [][]int{
			{123, 328, 51, 64},
			{45, 64, 387, 23},
			{6, 98, 215, 314},
		}

		// Act
		actual := parseNumsToInts(nums)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestParseInput(t *testing.T) {
	t.Run("when called then returns the values as strings and operations as functions", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expectedValues := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
		}
		expectedOpResults := [4][4]int{
			{2, 3, 6, -1},
			{2, 3, 5, -1},
			{2, 3, 6, -1},
			{2, 3, 5, -1},
		}

		// Act
		actualValues, actualOpFuncs := parseInput(filepath)
		for i := range expectedOpResults {
			expectedOpResults[i][3] = actualOpFuncs[i](expectedOpResults[i][0], expectedOpResults[i][1])
		}

		// Arrange
		if !reflect.DeepEqual(actualValues, expectedValues) {
			t.Fatalf("\nactualValues=\n%#v\nexpectedValues=\n%#v\n", actualValues, expectedValues)
		}
		for i := range expectedOpResults {
			if expectedOpResults[i][2] != expectedOpResults[i][3] {
				t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", expectedOpResults[i][2], expectedOpResults[i][3])
			}
		}
	})
}

func TestParseNumsViaColumns(t *testing.T) {
	t.Run("when called then returns a matrix with the numbers for the operations", func(t *testing.T) {
		// Arrange
		nums := []string{
			"123 328  51 64 ",
			" 45 64  387 23 ",
			"  6 98  215 314",
		}
		expected := [][]int{
			{1, 24, 356},
			{369, 248, 8},
			{32, 581, 175},
			{623, 431, 4},
		}

		// Act
		actual := parseNumsViaColumns(nums)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 4277556
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
		expected := 4583860641327
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
