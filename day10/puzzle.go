package day10

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"sort"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[10] = main
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
	lines := readInput(filename)

	part1Result := part1(lines)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(lines)
	fmt.Println("part2Result", part2Result)
}

func part1(lines []string) int {
	syntaxErrorScore := 0
	for _, line := range lines {
		cs := corruptionScore(line)
		syntaxErrorScore += cs
	}
	return syntaxErrorScore
}

func corruptionScore(line string) int {
	stack := []rune{}
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack = append(stack, c)
			continue
		}

		l := len(stack)
		popped := stack[l-1]
		stack = stack[:l-1]
		switch {
		case c == ')' && popped != '(':
			return 3
		case c == ']' && popped != '[':
			return 57
		case c == '}' && popped != '{':
			return 1197
		case c == '>' && popped != '<':
			return 25137
		}
	}
	return 0
}

func part2(lines []string) int {
	scores := []int{}
	for _, line := range lines {
		score, ok := completionScore(line)
		if !ok {
			continue
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func completionScore(line string) (int, bool) {
	stack := []rune{}
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack = append(stack, c)
			continue
		}

		l := len(stack)
		popped := stack[l-1]
		stack = stack[:l-1]
		switch {
		case c == ')' && popped != '(':
			return 0, false
		case c == ']' && popped != '[':
			return 0, false
		case c == '}' && popped != '{':
			return 0, false
		case c == '>' && popped != '<':
			return 0, false
		}
	}

	return scoreStack(stack), true
}

func scoreStack(s []rune) int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	score := 0
	for _, c := range s {
		score *= 5
		switch c {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
}
