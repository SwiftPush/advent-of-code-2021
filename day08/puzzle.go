package day08

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
	registry.Registry[8] = main
}

type Entry struct {
	signals []string
	outputs []string
}

func readInput(filename string) []Entry {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")
	entries := []Entry{}
	for _, inputLine := range inputLines {
		parts := strings.Split(inputLine, " | ")
		signals := strings.Fields(parts[0])
		outputVals := strings.Fields(parts[1])

		signalsSorted := []string{}
		for _, s := range signals {
			signalsSorted = append(signalsSorted, sortString(s))
		}
		outputsSorted := []string{}
		for _, o := range outputVals {
			outputsSorted = append(outputsSorted, sortString(o))
		}

		e := Entry{
			signals: signalsSorted,
			outputs: outputsSorted,
		}
		entries = append(entries, e)
	}

	return entries
}

func main() {
	filename := utils.ParseCommandLineArguments()
	entries := readInput(filename)

	part1Result := part1(entries)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(entries)
	fmt.Println("part2Result", part2Result)
}

func part1(entries []Entry) int {
	uniqueSegments := 0
	for _, e := range entries {
		for _, output := range e.outputs {
			switch len(output) {
			case 2, 4, 3, 7:
				uniqueSegments += 1
			}
		}
	}
	return uniqueSegments
}

func part2(entries []Entry) int {
	totalSum := 0
	for _, e := range entries {
		oneStr, fourStr := "", ""
		for _, s := range e.signals {
			switch len(s) {
			case 4: // digit 4
				fourStr = s
			case 2: // digit 1
				oneStr = s
			}
		}

		signalMapping := map[string]int{}
		for _, s := range e.signals {
			sLen := len(s)
			c1 := digitsInCommon(oneStr, s)
			c4 := digitsInCommon(fourStr, s)

			switch {
			case sLen == 2 && c1 == 2 && c4 == 2:
				signalMapping[s] = 1
			case sLen == 5 && c1 == 1 && c4 == 2:
				signalMapping[s] = 2
			case sLen == 5 && c1 == 2 && c4 == 3:
				signalMapping[s] = 3
			case sLen == 4 && c1 == 2 && c4 == 4:
				signalMapping[s] = 4
			case sLen == 5 && c1 == 1 && c4 == 3:
				signalMapping[s] = 5
			case sLen == 6 && c1 == 1 && c4 == 3:
				signalMapping[s] = 6
			case sLen == 3 && c1 == 2 && c4 == 2:
				signalMapping[s] = 7
			case sLen == 7 && c1 == 2 && c4 == 4:
				signalMapping[s] = 8
			case sLen == 6 && c1 == 2 && c4 == 4:
				signalMapping[s] = 9
			case sLen == 6 && c1 == 2 && c4 == 3:
				signalMapping[s] = 0
			}
		}

		result := ""
		for _, o := range e.outputs {
			result += strconv.Itoa(signalMapping[o])
		}
		resultInt, _ := strconv.Atoi(result)

		totalSum += resultInt
	}

	return totalSum
}

func digitsInCommon(a, b string) int {
	i := intersect(a, b)
	return len(i)
}

func intersect(a, b string) string {
	result := ""
	for _, c := range a {
		if strings.ContainsRune(b, c) {
			result += string(c)
		}
	}
	return result
}

func sortString(a string) string {
	s := strings.Split(a, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
