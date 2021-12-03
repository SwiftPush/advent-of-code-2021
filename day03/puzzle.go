package day00

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	registry.Registry[3] = main
}

func readInput(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	return inputLines
}

func main() {
	filename := utils.ParseCommandLineArguments()
	input := readInput(filename)

	part1Result := part1(input)
	part2Result := part2(input)

	fmt.Println("part1", part1Result)
	fmt.Println("part2", part2Result)
}

func part1(lines []string) int {
	gammaRateString, epsilonRateString := calcGammaEpsilon(lines)

	gammaRate, _ := strconv.ParseInt(gammaRateString, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateString, 2, 64)

	return int(gammaRate) * int(epsilonRate)
}

func part2(lines []string) int {
	oxygenString := calcOxygen(lines)
	co2String := calcCO2(lines)

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 64)
	co2, _ := strconv.ParseInt(co2String, 2, 64)

	return int(oxygen) * int(co2)
}

func calcOxygen(lines []string) string {
	return calcRating(lines, calcGamma)
}

func calcCO2(lines []string) string {
	return calcRating(lines, calcEpsilon)
}

type ratingFn func([]string) string

func calcRating(lines []string, rf ratingFn) string {
	for bitCounter := range lines[0] {
		if len(lines) <= 1 {
			break
		}

		ratingStr := rf(lines)

		newLines := []string{}
		for _, line := range lines {
			if line[bitCounter] == ratingStr[bitCounter] {
				newLines = append(newLines, line)
			}
		}

		lines = newLines
	}

	return lines[0]
}

func calcGamma(lines []string) string {
	gamma, _ := calcGammaEpsilon(lines)
	return gamma
}

func calcEpsilon(lines []string) string {
	_, epsilon := calcGammaEpsilon(lines)
	return epsilon
}

func calcGammaEpsilon(lines []string) (gammaRate string, epsilonRate string) {
	counters := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, char := range line {
			if char == '1' {
				counters[i] += 1
			}
		}
	}

	for i := range counters {
		if float64(counters[i]) >= float64(len(lines))/2.0 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	return
}
