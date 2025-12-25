package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/SimeonHristov99/aoc_2025/internal/day01"
	"github.com/SimeonHristov99/aoc_2025/internal/day02"
	"github.com/SimeonHristov99/aoc_2025/internal/day03"
	"github.com/SimeonHristov99/aoc_2025/internal/day04"
	"github.com/SimeonHristov99/aoc_2025/internal/day05"
	"github.com/SimeonHristov99/aoc_2025/internal/day06"
)

type Config struct {
	Day   int
	Part  int
	Input string
}

type Solver func(string) (int, error)

func parseArgs(args []string) (Config, error) {
	var config Config
	fs := flag.NewFlagSet("aoc_2025", flag.ContinueOnError)
	fs.IntVar(&config.Day, "day", 1, "which day to run")
	fs.IntVar(&config.Part, "part", 1, "which part to run")
	fs.StringVar(&config.Input, "input", "internal/day01/input.txt", "path to input file")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	return config, nil
}

func main() {
	solvers := map[int]map[int]Solver{
		1: {
			1: day01.SolvePart1,
			2: day01.SolvePart2,
		},
		2: {
			1: day02.SolvePart1,
			2: day02.SolvePart2,
		},
		3: {
			1: day03.SolvePart1,
			2: day03.SolvePart2,
		},
		4: {
			1: day04.SolvePart1,
			2: day04.SolvePart2,
		},
		5: {
			1: day05.SolvePart1,
			2: day05.SolvePart2,
		},
		6: {
			1: day06.SolvePart1,
			2: day06.SolvePart2,
		},
	}
	config, _ := parseArgs(os.Args[1:])
	if _, err := os.Stat(config.Input); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "file '%s' does not exist\n", config.Input)
		return
	}
	result, _ := solvers[config.Day][config.Part](config.Input)
	fmt.Printf("Day %d, Part=%d, Input='%s': %d\n", config.Day, config.Part, config.Input, result)
}
