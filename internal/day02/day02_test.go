package day02

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called with sample then returns parsed input", func(t *testing.T) {
		// Arrange
		filename := "sample.txt"
		expected := [][2]int{
			[2]int{11, 22},
			[2]int{95, 115},
			[2]int{998, 1012},
			[2]int{1188511880, 1188511890},
			[2]int{222220, 222224},
			[2]int{1698522, 1698528},
			[2]int{446443, 446449},
			[2]int{38593856, 38593862},
			[2]int{565653, 565659},
			[2]int{824824821, 824824827},
			[2]int{2121212118, 2121212124},
		}

		// Act
		actual := parseInput(filename)

		// Assert
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("when no sequence repeated twice then returns false", func(t *testing.T) {
		// Arrange
		input := 101
		expected := false

		// Act
		actual := isValid(input)

		// Assert
		if actual != expected {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actual, expected)
		}
	})
}
