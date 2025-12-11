package main

type Config struct {
	Day   int
	Part  int
	Input string
}

func ParseArgs(args []string) Config {
	return Config{
		Day:   1,
		Part:  1,
		Input: "internal/day01/part01/input.txt",
	}
}

func main() {
	// Load configuration: day, part, input
	// fmt.Println(solve(day=..., part=..., input=...)
}
