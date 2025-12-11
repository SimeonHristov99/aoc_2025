package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	f()
	_ = w.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	_ = r.Close()
	return string(b)
}

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

func TestMain(t *testing.T) {
	t.Run("when called then returns result of day 1 part 1 with input", func(t *testing.T) {
		// Arrange
		day := 1
		part := 1
		input := "internal/day01/part01/input.txt"
		result := 0
		expectedOutput := fmt.Sprintf("Day %d, Part=%d, Input='%s': %d\n", day, part, input, result)

		initialArgs := os.Args
		defer func() { os.Args = initialArgs }()
		os.Args = []string{initialArgs[0], fmt.Sprintf("-day=%d", day), fmt.Sprintf("-part=%d", part), fmt.Sprintf("-input=%s", input)}

		// Act
		actualOutput := captureStdout(main)

		// Assert
		if actualOutput != expectedOutput {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualOutput, expectedOutput)
		}
	})
}
