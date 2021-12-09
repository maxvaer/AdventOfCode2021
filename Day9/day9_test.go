package Day9

import "testing"

func TestSumLowPoints(t *testing.T) {
	heightMap := getHeightMapFromFile("./data.txt")
	sum := sumLowPoints(heightMap)
	expectedCount := 508
	if sum != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, sum)
	}
}

func TestBasinScore(t *testing.T) {
	heightMap := getHeightMapFromFile("./data.txt")
	score := calculateBasinScore(heightMap)
	expectedScore := 1564640
	if score != expectedScore {
		t.Fatalf(`Expected %d got %d`, expectedScore, score)
	}
}
