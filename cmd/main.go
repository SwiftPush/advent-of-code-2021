package main

import (
	_ "aoc/day01"
	_ "aoc/day02"

	"aoc/registry"

	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected day number")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("unable to parse day number")
		os.Exit(1)
	}

	f, ok := registry.Registry[day]
	if !ok {
		fmt.Println("no function for day")
		os.Exit(1)
	}

	f()
}
