package Day5

import (
	"testing"
)

func TestCountOverlap(t *testing.T) {
	lines := getLines("./data.txt", true)
	expected := uint(6548)
	maxX, maxY := getMaxXY(lines)
	board := initBoard(maxX, maxY)
	draw(board, lines)
	count := countOverlaps(board)
	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func TestCountOverlapDiagonal(t *testing.T) {
	lines := getLines("./data.txt", false)
	expected := uint(19663)
	maxX, maxY := getMaxXY(lines)
	board := initBoard(maxX, maxY)
	draw(board, lines)
	count := countOverlaps(board)
	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func BenchmarkOverlap(b *testing.B) {
	lines := getLines("./data.txt", true)
	maxX, maxY := getMaxXY(lines)
	board := initBoard(maxX, maxY)
	draw(board, lines)
	countOverlaps(board)
}

func BenchmarkOverlapDiagonal(b *testing.B) {
	lines := getLines("./data.txt", false)
	maxX, maxY := getMaxXY(lines)
	board := initBoard(maxX, maxY)
	draw(board, lines)
	countOverlaps(board)
}
