package Day1

import "testing"

func TestGetDepths(t *testing.T) {
	data := []int{ 199,200,208,210,200,207,240,269,260,263 }
	expected := 7
	count := GetDepths(data)

	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func TestGetDepthsWindow(t *testing.T) {
	data := []int{ 199,200,208,210,200,207,240,269,260,263 }
	expected := 5
	count := GetDepthsWindow(data)

	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func BenchmarkGetDepths(b *testing.B) {
	data := []int{ 199,200,208,210,200,207,240,269,260,263 }
	for i := 0; i < b.N; i++ {
		GetDepths(data)
	}
}

func BenchmarkGetDepthsWindow(b *testing.B) {
	data := []int{ 199,200,208,210,200,207,240,269,260,263 }
	for i := 0; i < b.N; i++ {
		GetDepthsWindow(data)
	}
}