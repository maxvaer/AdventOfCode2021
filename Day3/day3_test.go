package Day3

import "testing"

func TestCalculatePower(t *testing.T) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected := int64(198)
	power := CalculatePower(data)
	if power != expected {
		t.Fatalf(`Expected %d got %d`, expected, power)
	}
}

func TestCalculateLifeSupport(t *testing.T) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected := int64(230)
	rating := CalculateLifeSupport(data)
	if rating != expected {
		t.Fatalf(`Expected %d got %d`, expected, rating)
	}
}

func BenchmarkCalculatePower(b *testing.B) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	for i := 0; i < b.N; i++ {
		CalculatePower(data)
	}
}

func BenchmarkCalculateLifeSupport(b *testing.B) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	for i := 0; i < b.N; i++ {
		CalculateLifeSupport(data)
	}
}
