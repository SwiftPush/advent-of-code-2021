package day06

import (
	"aoc/registry"
	"aoc/utils"
	"strconv"

	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[6] = main
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

func main() {
	filename := utils.ParseCommandLineArguments()
	lines := readInput(filename)

	part1Result := part1(lines)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(lines)
	fmt.Println("part2Result", part2Result)
}

func part1(nums []int) int {
	return simulateFish(nums, 80)
}

func part2(nums []int) int {
	return simulateFish(nums, 256)
}

func simulateFish(nums []int, days int) int {
	counts := map[int]int{}
	for _, num := range nums {
		counts[num] += 1
	}

	for dayCounter := 0; dayCounter < days; dayCounter += 1 {
		newCounts := map[int]int{}
		for timer, count := range counts {
			newCounts[timer-1] = count
		}
		deadFishCount := newCounts[-1]
		delete(newCounts, -1)
		newCounts[6] += deadFishCount
		newCounts[8] += deadFishCount

		counts = newCounts
	}

	totalCount := 0
	for _, count := range counts {
		totalCount += count
	}
	return totalCount
}
