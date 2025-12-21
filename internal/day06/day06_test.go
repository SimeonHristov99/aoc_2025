package day06

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("when called then returns the values and operations", func(t *testing.T) {
		// Arrange
		filepath := "sample.txt"
		expectedValues := [][]int{
			{123, 328, 51, 64},
			{45, 64, 387, 23},
			{6, 98, 215, 314},
		}
		expectedOpResults := [4][4]int{
			{2, 3, 6, -1},
			{2, 3, 5, -1},
			{2, 3, 6, -1},
			{2, 3, 5, -1},
		}

		// Act
		actualValues, actualOpFuncs := parseInput(filepath)
		for i := range expectedOpResults {
			expectedOpResults[i][3] = actualOpFuncs[i](expectedOpResults[i][0], expectedOpResults[i][1])
		}

		// Arrange
		if !reflect.DeepEqual(actualValues, expectedValues) {
			t.Fatalf("\nactualValues=\n%#v\nexpectedValues=\n%#v\n", actualValues, expectedValues)
		}
		for i := range expectedOpResults {
			if expectedOpResults[i][2] != expectedOpResults[i][3] {
				t.Fatalf("\nactual=\n%#v\nexpected=\n%#v\n", expectedOpResults[i][2], expectedOpResults[i][3])
			}
		}
	})
}
