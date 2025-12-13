package day01

import "testing"

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
}

func TestSolvePart1(t *testing.T) {
	t.Run("when called then returns 0", func(t *testing.T) {
		// Arrange
		expected := 0
		var expectedError error = nil
		file := "internal/day01/sample.txt"

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
