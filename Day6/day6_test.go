package Day6

import "testing"

func TestEvolve(t *testing.T) {
	fish := getInitialFishFromFile("./data.txt")
	count := evolve(fish, 18)
	expected := 372300
	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func TestEvolveFast(t *testing.T) {
	fish := getInitialFishFromFile("./data.txt")
	sortedFish := sortFish(fish)
	count := fastEvolve(sortedFish, 256)
	expected := 1675781200288
	if count != expected {
		t.Fatalf(`Expected %d got %d`, expected, count)
	}
}

func BenchmarkEvolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fish := getInitialFishFromFile("./data.txt")
		fastEvolve(fish, 18)
	}
}

func BenchmarkFastEvolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fish := getInitialFishFromFile("./data.txt")
		sortedFish := sortFish(fish)
		fastEvolve(sortedFish, 256)
	}
}
