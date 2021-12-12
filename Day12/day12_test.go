package Day12

import "testing"

func TestGetAllPathsPartOne(t *testing.T) {
	adMap := getAdjacentMap("./data.txt")
	count := getAllPaths(adMap, true)
	expectedCount := 4885
	if count != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, count)
	}
}

func TestGetAllPathsPartTwo(t *testing.T) {
	adMap := getAdjacentMap("./data.txt")
	count := getAllPaths(adMap, false)
	expectedCount := 117095
	if count != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, count)
	}
}
