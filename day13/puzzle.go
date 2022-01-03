package day13

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"strconv"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[13] = main
}

type Point struct {
	x, y int
}

type Fold struct {
	dir      string
	location int
}

func readInput(filename string) ([]Point, []Fold) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	points := []Point{}
	folds := []Fold{}

	parsingPoints := true
	for _, inputLine := range inputLines {
		if inputLine == "" {
			parsingPoints = false
			continue
		}
		if parsingPoints {
			parts := strings.Split(inputLine, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points = append(points, Point{x: x, y: y})
		} else {
			parts := strings.Split(inputLine[11:], "=")
			loc, _ := strconv.Atoi(parts[1])
			folds = append(folds, Fold{dir: parts[0], location: loc})
		}
	}

	return points, folds
}

func main() {
	filename := utils.ParseCommandLineArguments()
	points, folds := readInput(filename)

	part1Result := part1(points, folds)
	fmt.Println("part1Result", part1Result)

	fmt.Println("part2Result")
	part2(points, folds)
}

func part1(points []Point, folds []Fold) int {
	fold := folds[0]

	newPoints := map[Point]bool{}
	for _, point := range points {
		newPoint := calcPos(point, fold)
		if newPoint != nil {
			newPoints[*newPoint] = true
		}
	}

	return len(newPoints)
}

func part2(points []Point, folds []Fold) {
	newPoints := map[Point]bool{}
	for _, point := range points {
		newPoints[point] = true
	}

	for _, fold := range folds {
		tempPoints := map[Point]bool{}
		for point := range newPoints {
			newPoint := calcPos(point, fold)
			if newPoint != nil {
				tempPoints[*newPoint] = true
			}
		}
		newPoints = tempPoints
	}

	printGrid(newPoints)
}

func calcPos(point Point, fold Fold) *Point {
	x, y := point.x, point.y

	switch fold.dir {
	case "x":
		if point.x == fold.location {
			return nil
		}
		if point.x < fold.location {
			x = point.x
		} else {
			x = 2*fold.location - point.x
		}
	case "y":
		if point.y == fold.location {
			return nil
		}
		if point.y < fold.location {
			y = point.y
		} else {
			y = 2*fold.location - point.y
		}
	}

	return &Point{x: x, y: y}
}

func printGrid(points map[Point]bool) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 39; x++ {
			point := Point{x: x, y: y}
			if _, ok := points[point]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
