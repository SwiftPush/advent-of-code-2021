package day04

import (
	"aoc/registry"
	"aoc/utils"

	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	registry.Registry[4] = main
}

func readInput(filename string) ([]int, []BingoBoard) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	numberDraw := parseNumberDraw(inputLines[0])

	bingoGrids := []BingoBoard{}
	for i := 2; i < len(inputLines); i += 6 {
		bingoGrid := parseBingoGrid(inputLines[i:])
		bingoGrids = append(bingoGrids, bingoGrid)
	}

	return numberDraw, bingoGrids
}

func parseNumberDraw(line string) []int {
	numsStr := strings.Split(line, ",")
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		num, _ := strconv.Atoi(numStr)
		nums[i] = num
	}
	return nums
}

type BingoBoard struct {
	grid              map[int]BingoCellProps
	markedCounterRows []int
	markedCounterCols []int
}

type BingoCellProps struct {
	point  Point
	marked bool
}

type Point struct {
	x, y int
}

func parseBingoGrid(lines []string) BingoBoard {
	board := BingoBoard{}
	board.grid = map[int]BingoCellProps{}
	board.markedCounterRows = make([]int, 5)
	board.markedCounterCols = make([]int, 5)

	for i := 0; i < 5; i += 1 {
		line := strings.TrimSpace(lines[i])
		lineSplit := strings.Fields(line)

		for j, numStr := range lineSplit {
			num, _ := strconv.Atoi(numStr)

			board.grid[num] = BingoCellProps{
				point:  Point{x: j, y: i},
				marked: false,
			}
		}
	}

	return board
}

func main() {
	filename := utils.ParseCommandLineArguments()

	numDraw, bingoGrids := readInput(filename)
	part1Result := part1(numDraw, bingoGrids)
	fmt.Println("part1", part1Result)

	numDraw, bingoGrids = readInput(filename)
	part2Result := part2(numDraw, bingoGrids)
	fmt.Println("part2", part2Result)
}

func part1(numDraw []int, bingoGrids []BingoBoard) int {
	winningNum, winningBoard := findWinningBoard(numDraw, bingoGrids)
	boardTotal := sumUnmarkedNumbers(winningBoard)
	return boardTotal * winningNum
}

func part2(numDraw []int, bingoGrids []BingoBoard) int {
	losingNum, losingBoard := findLosingBoard(numDraw, bingoGrids)
	boardTotal := sumUnmarkedNumbers(losingBoard)
	return boardTotal * losingNum
}

func findWinningBoard(numDraw []int, bingoGrids []BingoBoard) (int, BingoBoard) {
	for _, num := range numDraw {
		for i := range bingoGrids {
			props, ok := bingoGrids[i].grid[num]
			if !ok {
				continue
			}

			props.marked = true
			bingoGrids[i].grid[num] = props
			bingoGrids[i].markedCounterCols[props.point.x] += 1
			bingoGrids[i].markedCounterRows[props.point.y] += 1

			if bingoGrids[i].markedCounterCols[props.point.x] >= 5 {
				return num, bingoGrids[i]
			}
			if bingoGrids[i].markedCounterRows[props.point.y] >= 5 {
				return num, bingoGrids[i]
			}
		}
	}

	return -1, BingoBoard{}
}

func findLosingBoard(numDraw []int, bingoGrids []BingoBoard) (int, BingoBoard) {
	boards := map[int]bool{}
	for i := range bingoGrids {
		boards[i] = true
	}
	losingNum := -1
	lastBoard := -1

	for _, num := range numDraw {
		for i := range boards {
			lastBoard = i
		}

		for i := range bingoGrids {
			props, ok := bingoGrids[i].grid[num]
			if !ok {
				continue
			}

			props.marked = true
			bingoGrids[i].grid[num] = props
			bingoGrids[i].markedCounterCols[props.point.x] += 1
			bingoGrids[i].markedCounterRows[props.point.y] += 1

			if bingoGrids[i].markedCounterCols[props.point.x] >= 5 {
				delete(boards, i)
			}
			if bingoGrids[i].markedCounterRows[props.point.y] >= 5 {
				delete(boards, i)
			}
		}

		if len(boards) == 0 {
			losingNum = num
			break
		}
	}

	return losingNum, bingoGrids[lastBoard]
}

func sumUnmarkedNumbers(bingoGrid BingoBoard) int {
	sum := 0

	for num, props := range bingoGrid.grid {
		if !props.marked {
			sum += num
		}
	}

	return sum
}
