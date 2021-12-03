package day01

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	registry.Registry[1] = main
}

func readInput(filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputSplit := strings.Split(inputText, "\n")

	nums := make([]int, len(inputSplit))
	for i, inputString := range inputSplit {
		nums[i], _ = strconv.Atoi(inputString)
	}
	return nums
}

func main() {
	filename := utils.ParseCommandLineArguments()
	input := readInput(filename)

	part1Result := part1(input)
	part2Result := part2(input)

	fmt.Println(part1Result)
	fmt.Println(part2Result)
}

func part1(nums []int) int {
	return solver(nums, 1)
}

func part2(nums []int) int {
	return solver(nums, 3)
}

func solver(nums []int, windowSize int) int {
	increments := 0

	i, j := 0, windowSize
	for j < len(nums) {
		if nums[j] > nums[i] {
			increments += 1
		}
		i += 1
		j += 1
	}

	return increments
}
