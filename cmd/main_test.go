package main

import "testing"

func TestParseArgs(t *testing.T) {
	t.Run("when nothing set then should return default values", func(t *testing.T) {
		// Arrange
		expected := Config{
			Day:   1,
			Part:  1,
			Input: "internal/day01/part01/input.txt",
		}

		// Act
		actual := ParseArgs([]string{})

		// Assert
		if actual != expected {
			t.Fatalf("expected=\n%#v\nactual=\n%#v\n", expected, actual)
		}
	})
}
