package day01

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns array of strings", func(t *testing.T) {
		// Arrange
		filename := "sample.txt"
		expected := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

		// Act
		actual := ParseInput(filename)

		// Assert
		if reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestProcessSingle(t *testing.T) {
	t.Run("when called with left and single digit then returns negative digit", func(t *testing.T) {
		// Arrange
		input := "L5"
		expected := -5

		// Act
		actual := ProcessSingle(input)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when called with left and number then returns negative number", func(t *testing.T) {
		// Arrange
		input := "L99"
		expected := -99

		// Act
		actual := ProcessSingle(input)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when called with right and number then returns positive number", func(t *testing.T) {
		// Arrange
		input := "R60"
		expected := 60

		// Act
		actual := ProcessSingle(input)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns three", func(t *testing.T) {
		// Arrange
		expected := 3
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
