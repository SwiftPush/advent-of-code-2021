package day05

import (
	"aoc/registry"
	"aoc/utils"
	"strconv"

	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[5] = main
}

type Point struct {
	x, y int
}

type Line struct {
	src, dst Point
}

func readInput(filename string) []Line {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	lines := []Line{}
	for _, inputLine := range inputLines {
		line := parseInputLine(inputLine)
		lines = append(lines, line)
	}

	return lines
}

func parseInputLine(inputLine string) Line {
	points := strings.Split(inputLine, " -> ")
	src := parsePoint(points[0])
	dst := parsePoint(points[1])
	return Line{
		src: src,
		dst: dst,
	}
}

func parsePoint(input string) Point {
	numStrs := strings.Split(input, ",")
	nums := []int{}
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return Point{
		x: nums[0],
		y: nums[1],
	}
}

func main() {
	filename := utils.ParseCommandLineArguments()
	lines := readInput(filename)

	part1Result := part1(lines)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(lines)
	fmt.Println("part2Result", part2Result)
}

func part1(lines []Line) int {
	return calcOverlaps(lines, false)
}

func part2(lines []Line) int {
	return calcOverlaps(lines, true)
}

func calcOverlaps(lines []Line, includeDiagonalLines bool) int {
	pointCount := map[Point]int{}

	for _, line := range lines {
		xDir := convToDir(line.dst.x - line.src.x)
		yDir := convToDir(line.dst.y - line.src.y)

		if !includeDiagonalLines {
			if xDir != 0 && yDir != 0 {
				continue
			}
		}

		xPos, yPos := line.src.x, line.src.y
		for (xPos != line.dst.x) || (yPos != line.dst.y) {
			point := Point{x: xPos, y: yPos}
			pointCount[point] += 1

			xPos += xDir
			yPos += yDir
		}
		point := Point{x: xPos, y: yPos}
		pointCount[point] += 1
	}

	return countOverlaps(pointCount)
}

func convToDir(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	}
	return 1
}

func countOverlaps(pointCounts map[Point]int) (overlaps int) {
	for _, count := range pointCounts {
		if count > 1 {
			overlaps += 1
		}
	}
	return overlaps
}
