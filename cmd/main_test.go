package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdOutErr(f func()) (string, string) {
	rStdOut, wStdOut, _ := os.Pipe()
	rStdErr, wStdErr, _ := os.Pipe()
	oldStdOut := os.Stdout
	oldStdErr := os.Stderr
	os.Stdout = wStdOut
	os.Stderr = wStdErr
	f()
	_ = wStdOut.Close()
	_ = wStdErr.Close()
	os.Stdout = oldStdOut
	os.Stderr = oldStdErr
	stdOut, _ := io.ReadAll(rStdOut)
	stdErr, _ := io.ReadAll(rStdErr)
	_ = rStdOut.Close()
	_ = rStdErr.Close()
	return string(stdOut), string(stdErr)
}

func TestParseArgs(t *testing.T) {
	t.Run("when nothing set then should return default values", func(t *testing.T) {
		// Arrange
		expected := Config{
			Day:   1,
			Part:  1,
			Input: "internal/day01/input.txt",
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
		input := "../internal/day01/sample.txt"
		result := 0
		expectedOutput := fmt.Sprintf("Day %d, Part=%d, Input='%s': %d\n", day, part, input, result)

		initialArgs := os.Args
		defer func() { os.Args = initialArgs }()
		os.Args = []string{initialArgs[0], fmt.Sprintf("-day=%d", day), fmt.Sprintf("-part=%d", part), fmt.Sprintf("-input=%s", input)}

		// Act
		actualOutput, _ := captureStdOutErr(main)

		// Assert
		if actualOutput != expectedOutput {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualOutput, expectedOutput)
		}
	})

	t.Run("when file does not exist then raise an error", func(t *testing.T) {
		// Arrange
		day := 1
		part := 1
		input := "something.txt"
		expectedErrorMessage := fmt.Sprintf("file '%s' does not exist\n", input)
		expectedStdOut := ""

		initialArgs := os.Args
		defer func() { os.Args = initialArgs }()
		os.Args = []string{initialArgs[0], fmt.Sprintf("-day=%d", day), fmt.Sprintf("-part=%d", part), fmt.Sprintf("-input=%s", input)}

		// Act
		actualStdOut, actualStdErr := captureStdOutErr(main)

		// Assert
		if actualStdOut != expectedStdOut {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualStdOut, expectedStdOut)
		}
		if !strings.Contains(actualStdErr, expectedErrorMessage) {
			t.Fatalf("\nexpected=\n%#v\nactual=\n%#v\n", expectedErrorMessage, actualStdErr)
		}
	})
}
