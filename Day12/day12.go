package Day12

import (
	"AdventOfCode/Utils"
	"fmt"
	"strings"
)

func Run() {
	fmt.Println("Day12:")
	adMap := getAdjacentMap("./Day12/data.txt")
	fmt.Println("Part 1: ", getAllPaths(adMap, true))
	fmt.Println("Part 2: ", getAllPaths(adMap, false))
	fmt.Println("----------")
}

func getAdjacentMap(path string) map[string][]string {
	adjacentMap := make(map[string][]string)

	data := Utils.ReadFileAsString(path)
	for _, d := range data {
		nodes := strings.Split(d, "-")
		adjacentMap[nodes[0]] = append(adjacentMap[nodes[0]], nodes[1])
		adjacentMap[nodes[1]] = append(adjacentMap[nodes[1]], nodes[0])
	}

	return adjacentMap
}

func getAllPaths(adMap map[string][]string, partOne bool) int {
	curPath := []string{"start"}

	paths := getPaths(adMap, curPath, partOne)

	return len(paths)
}

func getPaths(adMap map[string][]string, path []string, partOne bool) [][]string {
	lastNode := path[len(path)-1]
	if lastNode == "end" {
		return [][]string{path}
	}

	var listOfPaths [][]string

	for _, s := range adMap[lastNode] {

		if partOne {
			if isLower(s) && Utils.StringArrayContains(path, s) {
				continue
			}
		}
		if isLower(s) && isTwiceVisited(path, s) || s == "start" {
			continue
		}
		listOfPaths = append(listOfPaths, getPaths(adMap, append(path, s), partOne)...)
	}

	return listOfPaths
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}

func isTwiceVisited(path []string, newLower string) bool {
	visited := false

	lowerMap := make(map[string]int)

	twoCounter := 0

	for _, s := range path {
		if isLower(s) {
			lowerMap[s]++
			if s == newLower {
				lowerMap[s]++
			}
			if lowerMap[s] >= 2 {
				twoCounter++
			}
		}
	}

	if twoCounter > 1 {
		visited = true
	}

	return visited

}
