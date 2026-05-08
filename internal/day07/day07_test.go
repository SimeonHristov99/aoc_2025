package day07

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns start coords and map", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expectedTimeline := Timeline{
			x:  0,
			y:  7,
			id: 1,
		}
		expectedMap := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}

		// Act
		actualTimeline, actualMap := parseInput(filepath)

		// Assert
		if !reflect.DeepEqual(actualMap, expectedMap) {
			t.Fatalf("\nactualMap=\n%#v\nexpectedMap=\n%#v\n", actualMap, expectedMap)
		}
		if !reflect.DeepEqual(actualTimeline, expectedTimeline) {
			t.Fatalf("\nactualTimeline=\n%#v\nexpectedTimeline=\n%#v\n", actualTimeline, expectedTimeline)
		}
	})
}

func TestDoBFS(t *testing.T) {
	t.Run("when called then does BFS", func(t *testing.T) {
		// Arrange
		startTimeline := Timeline{
			x:  0,
			y:  7,
			id: 1,
		}
		manifold := []string{
			".......S.......",
			"...............",
			".......^.......",
			"...............",
			"......^.^......",
			"...............",
			".....^.^.^.....",
			"...............",
			"....^.^...^....",
			"...............",
			"...^.^...^.^...",
			"...............",
			"..^...^.....^..",
			"...............",
			".^.^.^.^.^...^.",
			"...............",
		}
		var expectedError error = nil
		expectedSplits := 21

		// Act
		actualSplits, actualError := doBFS(startTimeline, manifold)

		// Assert
		if actualSplits != expectedSplits {
			t.Fatalf("\nactualSplits=\n%#v\nexpectedSplits=\n%#v\n", actualSplits, expectedSplits)
		}
		if actualError != expectedError {
			t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", actualError, expectedError)
		}
	})
}

func TestSolvePart1(t *testing.T) {
	t.Run("when called with sample then returns part one result", func(t *testing.T) {
		// Arrange
		expected := 21
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
		expected := 1585
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
