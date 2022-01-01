package day12

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"

	"io/ioutil"
	"strings"
)

func init() {
	registry.Registry[12] = main
}

type connection struct {
	a, b string
}

type node struct {
	name       string
	neighbours map[string]bool
}

func readInput(filename string) map[string]node {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")
	connections := []connection{}
	for _, inputLine := range inputLines {
		split := strings.Split(inputLine, "-")
		connection := connection{
			a: split[0],
			b: split[1],
		}
		connections = append(connections, connection)
	}

	nodes := map[string]node{}
	for _, connection := range connections {
		nodes[connection.a] = node{name: connection.a, neighbours: map[string]bool{}}
		nodes[connection.b] = node{name: connection.b, neighbours: map[string]bool{}}
	}

	for _, connection := range connections {
		nodes[connection.a].neighbours[connection.b] = true
		nodes[connection.b].neighbours[connection.a] = true
	}
	return nodes
}

func main() {
	filename := utils.ParseCommandLineArguments()
	nodes := readInput(filename)

	part1Result := part1(nodes)
	fmt.Println("part1Result", part1Result)

	part2Result := part2(nodes)
	fmt.Println("part2Result", part2Result)
}

func part1(nodes map[string]node) int {
	paths := dfs1(nodes, "start", "end", []string{})
	return len(paths)
}

func part2(nodes map[string]node) int {
	paths := dfs2(nodes, "start", "end", []string{})
	return len(paths)
}

func dfs1(nodes map[string]node, current, target string, path []string) [][]string {
	newPath := append([]string{}, path...)
	newPath = append(newPath, current)

	if current == target {
		return [][]string{newPath}
	}

	paths := [][]string{}
	for n := range nodes[current].neighbours {
		if isBig(n) || !containsString(path, n) {
			paths = append(paths, dfs1(nodes, n, target, newPath)...)
		}
	}
	return paths
}

func dfs2(nodes map[string]node, current, target string, path []string) [][]string {
	newPath := append([]string{}, path...)
	newPath = append(newPath, current)

	if current == target {
		return [][]string{newPath}
	}

	paths := [][]string{}
	for n := range nodes[current].neighbours {
		if isBig(n) || canVisitSmallCave(newPath, n) {
			paths = append(paths, dfs2(nodes, n, target, newPath)...)
		}
	}
	return paths
}

func containsString(slice []string, s string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}

func canVisitSmallCave(path []string, c string) bool {
	if c == "start" {
		return false
	}

	alreadyRevisited := false
	counts := map[string]int{}
	for _, p := range path {
		counts[p] += 1
	}
	for k, v := range counts {
		if !isBig(k) && v > 1 {
			alreadyRevisited = true
		}
	}

	switch counts[c] {
	case 0:
		return true
	case 1:
		return !alreadyRevisited
	}

	return false
}

func isBig(s string) bool {
	return strings.ToUpper(s) == s
}
