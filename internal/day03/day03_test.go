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
