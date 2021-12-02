package Day2

import "testing"

func TestCalculatePosition(t *testing.T) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expected := 150
	result := CalculatePosition(data, false)

	if result != expected {
		t.Fatalf(`Expected %d got %d`, expected, result)
	}
}

func TestCalculatePositionAimmode(t *testing.T) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expected := 900
	result := CalculatePosition(data, true)

	if result != expected {
		t.Fatalf(`Expected %d got %d`, expected, result)
	}
}

func BenchmarkGetDepths(b *testing.B) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	for i := 0; i < b.N; i++ {
		CalculatePosition(data, false)
	}
}

func BenchmarkGetDepthsWindow(b *testing.B) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	for i := 0; i < b.N; i++ {
		CalculatePosition(data, true)
	}
}
