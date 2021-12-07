package Day7

import "testing"

func TestCalculateFuelCostForPositions(t *testing.T) {
	positions, maxPos := getPositionsFromFile("./data.txt")
	cost, position := calculateFuelCostForPositions(positions, maxPos, false)
	expectedCost := uint(336040)
	expectedPos := uint(323)
	if cost != expectedCost {
		t.Fatalf(`Expected %d got %d`, expectedCost, cost)
	}
	if position != expectedPos {
		t.Fatalf(`Expected %d got %d`, expectedPos, position)
	}
}

func TestCalculateFuelCostForPositionsSum(t *testing.T) {
	positions, maxPos := getPositionsFromFile("./data.txt")
	cost, position := calculateFuelCostForPositions(positions, maxPos, true)
	expectedCost := uint(94813675)
	expectedPos := uint(463)
	if cost != expectedCost {
		t.Fatalf(`Expected %d got %d`, expectedCost, cost)
	}
	if position != expectedPos {
		t.Fatalf(`Expected %d got %d`, expectedPos, position)
	}
}

func BenchmarkCalculateFuelCostForPositions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		positions, maxPos := getPositionsFromFile("./data.txt")
		calculateFuelCostForPositions(positions, maxPos, false)
	}
}

func BenchmarkCalculateFuelCostForPositionsSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		positions, maxPos := getPositionsFromFile("./data.txt")
		calculateFuelCostForPositions(positions, maxPos, true)
	}
}
