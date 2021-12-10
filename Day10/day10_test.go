package Day10

import (
	"AdventOfCode/Utils"
	"testing"
)

func TestCalculateCorrupted(t *testing.T) {
	chunks := Utils.ReadFileAsString("./data.txt")
	score := calculateCorrupted(chunks)
	expectedScore := uint(339411)
	if score != expectedScore {
		t.Fatalf(`Expected %d got %d`, expectedScore, score)
	}
}

func TestCalculateIncompleteScore(t *testing.T) {
	chunks := Utils.ReadFileAsString("./data.txt")
	score := calculateIncompleteScore(chunks)
	expectedScore := 2289754624
	if score != expectedScore {
		t.Fatalf(`Expected %d got %d`, expectedScore, score)
	}
}
