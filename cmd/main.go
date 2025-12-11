package main

import "flag"

type Config struct {
	Day   int
	Part  int
	Input string
}

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
	// Load configuration: day, part, input
	// fmt.Println(solve(day=..., part=..., input=...)
}
