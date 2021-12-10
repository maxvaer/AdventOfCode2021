package Day10

import (
	"AdventOfCode/Utils"
	"fmt"
	"sort"
	"strings"
)

func Run() {
	fmt.Println("Day10:")
	chunks := Utils.ReadFileAsString("./Day10/data.txt")
	fmt.Println("Corrupted Score:", calculateCorrupted(chunks))
	fmt.Println("Incomplete Score:", calculateIncompleteScore(chunks))
	fmt.Println("----------")
}

func calculateIncompleteScore(chunks []string) int {
	var incompleteCut []string
	var scores []int

	for _, chunk := range chunks {
		expected, found, cutChunk := checkChunk(chunk)

		if expected == "" && found == "" {

			rev := revCutChunk(cutChunk)
			incompleteCut = append(incompleteCut, cutChunk)

			score := 0

			for _, s := range rev {
				char := string(s)
				score *= 5
				if char == ")" {
					score += 1
				}
				if char == "]" {
					score += 2
				}
				if char == "}" {
					score += 3
				}
				if char == ">" {
					score += 4
				}
			}
			scores = append(scores, score)
		}

	}

	sort.Ints(scores)

	scoreIndex := len(scores) / 2

	return scores[scoreIndex]
}

func calculateCorrupted(chunks []string) uint {
	foundParenthesisCount := uint(0)
	foundBracketCount := uint(0)
	foundCurleyCount := uint(0)
	foundTagCount := uint(0)

	for _, chunk := range chunks {
		expected, found, _ := checkChunk(chunk)

		if expected != "" && found != "" {
			if found == ")" {
				foundParenthesisCount++
			}
			if found == "]" {
				foundBracketCount++
			}
			if found == "}" {
				foundCurleyCount++
			}
			if found == ">" {
				foundTagCount++
			}
		}
	}
	return (foundParenthesisCount * 3) + (foundBracketCount * 57) + (foundCurleyCount * 1197) + (foundTagCount * 25137)
}

func revCutChunk(cutChunk string) string {
	result := ""

	for i := len(cutChunk); i > 0; i-- {
		closing := getClosing(string(cutChunk[i-1]))

		result += closing
	}

	return result
}

func checkChunk(chunk string) (string, string, string) {
	cut := true
	for {
		startLength := len(chunk)
		if !cut {
			break
		}
		chunk = strings.Replace(chunk, "[]", "", -1)
		chunk = strings.Replace(chunk, "()", "", -1)
		chunk = strings.Replace(chunk, "{}", "", -1)
		chunk = strings.Replace(chunk, "<>", "", -1)

		cutLength := len(chunk)

		if startLength == cutLength {
			cut = false
		}

	}

	found := ""
	expected := ""

	for i := 0; i < len(chunk)-1; i++ {
		if isClosing(string(chunk[i+1])) {
			if isCorrupt(string(chunk[i]), string(chunk[i+1])) {
				found = string(chunk[i+1])
				expected = getClosing(string(chunk[i]))
				break
			}
		}
	}

	return expected, found, chunk
}

func isClosing(s string) bool {
	return s == ")" || s == ">" || s == "}" || s == "]"
}

func isCorrupt(openning string, closing string) bool {
	if openning == "(" && closing == ")" {
		return false
	}
	if openning == "[" && closing == "]" {
		return false
	}
	if openning == "<" && closing == ">" {
		return false
	}
	if openning == "{" && closing == "}" {
		return false
	}

	return true
}

func getClosing(openning string) string {
	if openning == "(" {
		return ")"
	} else if openning == "<" {
		return ">"
	} else if openning == "{" {
		return "}"
	} else {
		return "]"
	}
}
