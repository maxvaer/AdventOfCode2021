package Day8

import "testing"

func TestCountSimpleDigits(t *testing.T) {
	_, digitOutput := GetSegmentData("./data.txt")
	count := countSimpleDigits(digitOutput)
	expectedCount := uint(247)
	if count != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, count)
	}
}

func TestCountOutput(t *testing.T) {
	signalPattern, digitOutput := GetSegmentData("./data.txt")
	count := calculateOutput(signalPattern, digitOutput)
	expectedCount := uint(933305)
	if count != expectedCount {
		t.Fatalf(`Expected %d got %d`, expectedCount, count)
	}
}
