package day03

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns parsed input", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expected := []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}

		// Act
		actual := parseInput(filepath)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestFindIdxMaxDigit(t *testing.T) {
	t.Run("when called then returns max digit", func(t *testing.T) {
		// Arrange
		input := "818181911112111"
		expected := 6

		// Act
		actual := findIdxMaxDigit(input)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestFindMaxJoltage(t *testing.T) {
	t.Run("when max digits at edges then returns maximum joltage", func(t *testing.T) {
		// Arrange
		input := "811111111111119"
		neededBatteries := 2
		expected := 89

		// Act
		actual := findMaxJoltage(input, neededBatteries)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when left max digit not at edge then returns maximum joltage", func(t *testing.T) {
		// Arrange
		input := "1111111811119"
		neededBatteries := 2
		expected := 89

		// Act
		actual := findMaxJoltage(input, neededBatteries)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when temp then returns maximum joltage", func(t *testing.T) {
		// Arrange
		input := "234234234234278"
		neededBatteries := 12
		expected := 434234234278

		// Act
		actual := findMaxJoltage(input, neededBatteries)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 357
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
		expected := 17109
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
		expected := 3121910778619
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

