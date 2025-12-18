package day05

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns input parsed", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expected := DB{
			ingredientRanges: [][2]int{
				{3, 5},
				{10, 14},
				{16, 20},
				{12, 18},
			},
			ingredientIds: []int{1, 5, 8, 11, 17, 32},
		}

		// Act
		actual := parseInput(filepath)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestIsFresh(t *testing.T) {
	t.Run("when in range then returns true", func(t *testing.T) {
		// Arrange
		ingredientRanges := [][2]int{
			{3, 5},
			{10, 14},
			{16, 20},
			{12, 18},
		}
		ingredientId := 5
		expected := true

		// Act
		actual := isFresh(ingredientRanges, ingredientId)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when not in range then returns false", func(t *testing.T) {
		// Arrange
		ingredientRanges := [][2]int{
			{3, 5},
			{10, 14},
			{16, 20},
			{12, 18},
		}
		ingredientId := 8
		expected := false

		// Act
		actual := isFresh(ingredientRanges, ingredientId)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestUnionSize(t *testing.T) {
	t.Run("when no intersection then returns sum of elements", func(t *testing.T) {
		// Arrange
		lhs := [2]int{3, 5}
		rhs := [2]int{6, 10}
		expected := 8

		// Act
		actual := unionSize(lhs, rhs)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when intersection on the right then returns union size without duplicates", func(t *testing.T) {
		// Arrange
		lhs := [2]int{3, 5}
		rhs := [2]int{4, 10}
		expected := 8

		// Act
		actual := unionSize(lhs, rhs)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

	t.Run("when equal right element then returns union size without duplicates", func(t *testing.T) {
		// Arrange
		lhs := [2]int{3, 5}
		rhs := [2]int{5, 10}
		expected := 8

		// Act
		actual := unionSize(lhs, rhs)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})

}

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
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

	t.Run("when called with input then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 611
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
