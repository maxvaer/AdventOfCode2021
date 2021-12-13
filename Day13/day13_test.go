package Day13

import "testing"

func TestFold(t *testing.T) {
	points, folds, maxX, maxY := getInstructionsFromFile("./data.txt")
	board := initBoard(maxY, maxX)
	board = drawPoints(board, points)
	counts, _ := fold(board, folds, maxX, maxY)
	expectedCount := 716
	if counts[0] != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, counts[0])
	}
}
