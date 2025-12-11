package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Run("when nothing set then should return default values", func(t *testing.T) {
		// Arrange
		expected := Config{
			Day:   1,
			Part:  1,
			Input: "internal/day01/part01/input.txt",
		}
		var expectedErr error = nil

		// Act
		actual, actualErr := ParseArgs([]string{})

		// Assert
		if actualErr != expectedErr {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedErr, actualErr)
		}
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
		var expectedErr error = nil

		// Act
		actual, actualErr := ParseArgs([]string{
			fmt.Sprintf("-day=%d", day),
			fmt.Sprintf("-part=%d", part),
			fmt.Sprintf("-input=%s", input),
		})

		// Assert
		if actualErr != expectedErr {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedErr, actualErr)
		}
		if actual != expected {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expected, actual)
		}
	})

	t.Run("when some arguments passed then they are parsed and remaining are default", func(t *testing.T) {
		// Arrange
		input := "tmp1.txt"

		expected := Config{
			Day:   1,
			Part:  1,
			Input: input,
		}
		var expectedErr error = nil

		// Act
		actual, actualErr := ParseArgs([]string{
			fmt.Sprintf("-input=%s", input),
		})

		// Assert
		if actualErr != expectedErr {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedErr, actualErr)
		}
		if actual != expected {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expected, actual)
		}
	})

	t.Run("when parse errors then returns error message", func(t *testing.T) {
		// Arrange
		day := 12
		part := "something"
		input := "tmp1.txt"

		expectedConfig := Config{}
		expectedErrorMessage := "invalid value \"something\" for flag -part: parse error"

		// Act
		actualConfig, err := ParseArgs([]string{
			fmt.Sprintf("-day=%d", day),
			fmt.Sprintf("-part=%s", part),
			fmt.Sprintf("-input=%s", input),
		})
		actualErrorMessage := err.Error()

		// Assert
		if actualConfig != expectedConfig {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedConfig, actualConfig)
		}
		if !strings.Contains(actualErrorMessage, expectedErrorMessage) {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedErrorMessage, actualErrorMessage)
		}
	})
}
