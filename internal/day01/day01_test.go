package day01

import "testing"

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
