package main

import "flag"

type Config struct {
	Day   int
	Part  int
	Input string
}

func ParseArgs(args []string) Config {
	var config Config
	fs := flag.NewFlagSet("aoc_2025", flag.PanicOnError)
	fs.IntVar(&config.Day, "day", 1, "which day to run")
	fs.IntVar(&config.Part, "part", 1, "which part to run")
	fs.StringVar(&config.Input, "input", "internal/day01/part01/input.txt", "path to input file")
	fs.Parse(args)
	return config
}

func main() {
	// Load configuration: day, part, input
	// fmt.Println(solve(day=..., part=..., input=...)
}
