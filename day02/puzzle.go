package day02

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	registry.Registry[2] = main
}

type Line struct {
	direction string
	distance  int
}

func readInput(filename string) []Line {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	lines := make([]Line, len(inputLines))
	for i, inputLine := range inputLines {
		lineSplit := strings.Split(inputLine, " ")
		distance, _ := strconv.Atoi(lineSplit[1])
		lines[i] = Line{
			direction: lineSplit[0],
			distance:  distance,
		}
	}
	return lines
}

func main() {
	filename := utils.ParseCommandLineArguments()
	input := readInput(filename)

	part1Result := part1(input)
	part2Result := part2(input)

	fmt.Println(part1Result)
	fmt.Println(part2Result)
}

func part1(lines []Line) int {
	x, y := 0, 0
	for _, line := range lines {
		switch line.direction {
		case "forward":
			x += line.distance
		case "down":
			y -= line.distance
		case "up":
			y += line.distance
		}
	}

	depth := y * -1
	return x * depth
}

func part2(lines []Line) int {
	x, y, aim := 0, 0, 0

	for _, line := range lines {
		switch line.direction {
		case "forward":
			x += line.distance
			y -= aim * line.distance
		case "down":
			aim += line.distance
		case "up":
			aim -= line.distance
		}
	}

	depth := y * -1
	return x * depth
}
