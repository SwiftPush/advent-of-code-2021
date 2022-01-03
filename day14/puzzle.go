package day14

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[14] = main
}

func readInput(filename string) (string, map[string]string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	template := inputLines[0]
	rules := map[string]string{}
	for _, inputLine := range inputLines[2:] {
		parts := strings.Split(inputLine, " -> ")
		rules[parts[0]] = parts[1]
	}

	return template, rules
}

func main() {
	filename := utils.ParseCommandLineArguments()
	template, rules := readInput(filename)

	part1Result := process(template, rules, 10)
	fmt.Println("part1Result", part1Result)

	part2Result := process(template, rules, 40)
	fmt.Println("part2Result", part2Result)
}

func process(template string, rules map[string]string, numSteps int) int {
	pairs := initPairs(template)

	for step := 0; step < numSteps; step++ {
		newPairs := map[string]int{}
		for pair, count := range pairs {
			insertion := rules[pair]

			if insertion == "" {
				newPairs[pair] = count
			} else {
				newPair1 := string(pair[0]) + insertion
				newPair2 := insertion + string(pair[1])
				newPairs[newPair1] += count
				newPairs[newPair2] += count
			}
		}
		pairs = newPairs
	}

	charCounts := countChars(pairs)

	return getLeastMostDifference(charCounts)
}

func initPairs(template string) map[string]int {
	template = template + "$"

	pairs := map[string]int{}
	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		pairs[pair] += 1
	}

	return pairs
}

func getLeastMostDifference(letterCounts map[string]int) int {
	mostCommon, leastCommon := -1, -1
	for _, v := range letterCounts {
		if mostCommon == -1 {
			mostCommon, leastCommon = v, v
		}
		if v > mostCommon {
			mostCommon = v
		}
		if v < leastCommon {
			leastCommon = v
		}
	}
	return mostCommon - leastCommon
}

func countChars(pairs map[string]int) map[string]int {
	charCounts := map[string]int{}
	for pair, count := range pairs {
		c := string(pair[0])
		if c == "$" {
			continue
		}
		charCounts[c] += count
	}
	return charCounts
}
