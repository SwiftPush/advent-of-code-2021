package day15

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"strconv"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[15] = main
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

type Point struct {
	x, y int
}

type Elem struct {
	minDistance int
	p           Point
}

func (e Elem) Less(a QItem) bool {
	return e.minDistance < a.(Elem).minDistance
}

func part1(grid [][]int) int {
	return search(grid)
}

func part2(grid [][]int) int {
	newGrid := make([][]int, len(grid)*5)
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid[0])*5)
	}
	for y, row := range grid {
		for x, v := range row {
			newGrid[y][x] = v
		}
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			nextGrid := incrementGrid(grid, x+y)
			assignGrid(newGrid, nextGrid, x*len(grid[0]), y*len(grid))
		}
	}

	return search(newGrid)
}

func incrementGrid(grid [][]int, val int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid[0]))
	}
	for y := range grid {
		for x := range grid[y] {
			newGrid[y][x] = (grid[y][x]+val-1)%9 + 1
		}
	}
	return newGrid
}

func assignGrid(grid [][]int, newGrid [][]int, x, y int) {
	for i := 0; i < len(newGrid); i++ {
		for j := 0; j < len(newGrid[i]); j++ {
			grid[y+i][x+j] = newGrid[i][j]
		}
	}
}

func search(grid [][]int) int {
	pq := NewPriorityQueue()
	e := Elem{
		p:           Point{x: 0, y: 0},
		minDistance: grid[0][0],
	}
	pq.Push(e)
	distances := map[Point]int{}
	distances[Point{x: 0, y: 0}] = grid[0][0]

	for pq.Length() != 0 {
		u := pq.Pop().(Elem)

		neighbours := []Point{
			{x: u.p.x + 1, y: u.p.y},
			{x: u.p.x - 1, y: u.p.y},
			{x: u.p.x, y: u.p.y + 1},
			{x: u.p.x, y: u.p.y - 1},
		}
		for _, n := range neighbours {
			if !isPointInRange(grid, n.x, n.y) {
				continue
			}
			alt := u.minDistance + grid[n.y][n.x]
			dist, ok := distances[n]
			if !ok {
				dist = 99999
			}
			if alt < dist {
				distances[n] = alt
				pq.Push(Elem{p: n, minDistance: alt})
			}
		}
	}

	return distances[Point{y: len(grid) - 1, x: len(grid[0]) - 1}] - grid[0][0]
}

func isPointInRange(grid [][]int, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if y >= len(grid) {
		return false
	}
	if x >= len(grid[0]) {
		return false
	}
	return true
}
