package day09

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[9] = main
}

func readInput(filename string) [][]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")
	grid := [][]int{}
	for _, inputLine := range inputLines {
		gridLine := []int{}
		for _, c := range inputLine {
			i, _ := strconv.Atoi(string(c))
			gridLine = append(gridLine, i)
		}
		grid = append(grid, gridLine)
	}

	return grid
}

func main() {
	filename := utils.ParseCommandLineArguments()
	grid := readInput(filename)

	part1Result := part1(grid)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(grid)
	fmt.Println("part2Result", part2Result)
}

func part1(grid [][]int) int {
	sum := 0

	for i := range grid {
		for j := range grid[i] {
			num := grid[i][j]

			switch {
			case (i-1) >= 0 && grid[i-1][j] <= num:
			case (i+1) < len(grid) && grid[i+1][j] <= num:
			case (j-1) >= 0 && grid[i][j-1] <= num:
			case (j+1) < len(grid[i]) && grid[i][j+1] <= num:
			default:
				riskLevel := num + 1
				sum += riskLevel
			}
		}
	}

	return sum
}

func part2(grid [][]int) int {
	basinSizes := []int{}
	for i := range grid {
		for j := range grid[i] {
			basinSize := fill(grid, i, j)
			if basinSize > 0 {
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}

	sort.Slice(basinSizes, func(i, j int) bool { // desc
		return basinSizes[i] > basinSizes[j]
	})

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func fill(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) {
		return 0
	}
	if j < 0 || j >= len(grid[i]) {
		return 0
	}
	if grid[i][j] == -1 {
		return 0
	}
	if grid[i][j] == 9 {
		return 0
	}

	grid[i][j] = -1
	return 1 + fill(grid, i-1, j) + fill(grid, i+1, j) + fill(grid, i, j-1) + fill(grid, i, j+1)
}
