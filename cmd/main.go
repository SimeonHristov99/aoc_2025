package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SimeonHristov99/aoc_2025/internal/day01"
)

type Config struct {
	Day   int
	Part  int
	Input string
}

type Solver func(string) (int, error)

func ParseArgs(args []string) (Config, error) {
	var config Config
	fs := flag.NewFlagSet("aoc_2025", flag.ContinueOnError)
	fs.IntVar(&config.Day, "day", 1, "which day to run")
	fs.IntVar(&config.Part, "part", 1, "which part to run")
	fs.StringVar(&config.Input, "input", "internal/day01/part01/input.txt", "path to input file")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	return config, nil
}

func main() {
	solvers := map[int]map[int]Solver{
		1: {
			1: day01.SolvePart1,
		},
	}
	config, _ := ParseArgs(os.Args[1:])
	result, _ := solvers[config.Day][config.Part](config.Input)
	fmt.Printf("Day %d, Part=%d, Input='%s': %d\n", config.Day, config.Part, config.Input, result)
}
