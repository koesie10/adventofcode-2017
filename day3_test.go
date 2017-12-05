package adventofcode2017

import (
	"testing"
)

func TestDay3Part1Examples(t *testing.T) {
	runs := []struct {
		input    int64
		expected int64
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for i, r := range runs {
		out := Day3Part1(r.input)

		if out != r.expected {
			t.Errorf("run %d: Expected %d, got %d for input %d", i, r.expected, out, r.input)
		}
	}
}

const day3Input int64 = 325489

func TestDay3Part1Input(t *testing.T) {
	var expected int64 = 552

	out := Day3Part1(day3Input)
	if out != expected {
		t.Errorf("Expected %d, got %d for input %d", expected, out, day3Input)
	}
}
