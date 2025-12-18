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
}
