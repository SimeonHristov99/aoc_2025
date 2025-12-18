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
			ingredientRange: [][2]int{
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
