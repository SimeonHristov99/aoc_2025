package day01

import "testing"

func TestSolvePart1(t *testing.T) {
	t.Run("when called then returns 0", func(t *testing.T) {
		// Arrange
		expected := 0

		// Act
		actual := solvePart1()

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
