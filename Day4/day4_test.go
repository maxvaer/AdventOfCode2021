package Day4

import (
	"AdventOfCode/Utils"
	"testing"
)

func TestGetScore(t *testing.T) {
	data := Utils.ReadFileAsString("./test.txt")
	expected := uint(4512)
	score := getScore(data, false)
	if score != expected {
		t.Fatalf(`Expected %d got %d`, expected, score)
	}
}

func TestGetScoreLast(t *testing.T) {
	data := Utils.ReadFileAsString("./test.txt")
	expected := uint(1924)
	score := getScore(data, true)
	if score != expected {
		t.Fatalf(`Expected %d got %d`, expected, score)
	}
}

func BenchmarkGetScore(b *testing.B) {
	data := Utils.ReadFileAsString("./test.txt")
	for i := 0; i < b.N; i++ {
		getScore(data, false)
	}
}

func BenchmarkGetScoreLast(b *testing.B) {
	data := Utils.ReadFileAsString("./test.txt")
	for i := 0; i < b.N; i++ {
		getScore(data, true)
	}
}
