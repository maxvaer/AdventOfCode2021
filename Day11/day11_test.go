package Day11

import (
	"testing"
)

func TestCalculateFlashes(t *testing.T) {
	board := getPositionsFromFile("./data.txt")
	flashes := calculateFlashes(board, 100)
	expectedFlashes := 1603
	if flashes != expectedFlashes {
		t.Fatalf(`Expected %d got %d`, expectedFlashes, flashes)
	}
}

func TestCalculateSyncFlash(t *testing.T) {
	board := getPositionsFromFile("./data.txt")
	sync := calculateSyncFlash(board)
	expectedSync := 222
	if sync != expectedSync {
		t.Fatalf(`Expected %d got %d`, expectedSync, sync)
	}
}
