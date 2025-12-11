package main

import (
	"testing"
	"fmt"
)

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
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expected, actual)
		}
	})

	t.Run("when all arguments passed then should return them", func(t *testing.T) {
		// Arrange
		day := 12
		part := 2
		input := "tmp1.txt"

		expected := Config{
			Day:   day,
			Part:  part,
			Input: input,
		}

		// Act
		actual := ParseArgs([]string{
			fmt.Sprintf("-day=%d", day),
			fmt.Sprintf("-part=%d", part),
			fmt.Sprintf("-input=%s", input),
		})

		// Assert
		if actual != expected {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expected, actual)
		}
	})
}
