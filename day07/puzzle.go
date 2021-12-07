package day07

import (
	"aoc/registry"
	"aoc/utils"
	"strconv"

	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[7] = main
}

func readInput(filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	numStrs := strings.Split(inputText, ",")
	nums := []int{}
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return nums
}

type fuelFn func(int) int

func main() {
	filename := utils.ParseCommandLineArguments()
	nums := readInput(filename)

	part1Result := part1(nums)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(nums)
	fmt.Println("part2Result", part2Result)
}

func part1(nums []int) int {
	return findLowestPos(nums, abs)
}

func part2(nums []int) int {
	return findLowestPos(nums, triangleNum)
}

func findLowestPos(nums []int, f fuelFn) int {
	minPos, maxPos := getMinMax(nums)
	var minTotalFuel *int
	for targetPos := minPos; targetPos < maxPos; targetPos++ {
		totalFuel := 0
		for _, num := range nums {
			totalFuel += f(num - targetPos)
		}
		if minTotalFuel == nil || totalFuel < *minTotalFuel {
			minTotalFuel = &totalFuel
		}
	}

	return *minTotalFuel
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func getMinMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min, max
}

func triangleNum(n int) int {
	n = abs(n)
	return n * (n + 1) / 2
}
