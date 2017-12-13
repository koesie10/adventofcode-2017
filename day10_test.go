package adventofcode2017

import "testing"

func TestDay10Part1(t *testing.T) {
	runs := []struct {
		name     string
		input    string
		listSize int
		expected int
	}{
		{"Example", "3,4,1,5", 5, 12},
		{"Input", day10Input, 256, 20056},
	}

	for _, r := range runs {
		t.Run(r.name, func(t *testing.T) {
			out := Day10Part1(r.input, r.listSize)

			if out != r.expected {
				t.Errorf("Expected %v, got %v", r.expected, out)
			}
		})
	}
}

func TestDay10Part2(t *testing.T) {
	runs := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
		{"Input", day10Input, "d9a7de4a809c56bf3a9465cb84392c8e"},
	}

	for _, r := range runs {
		t.Run(r.name, func(t *testing.T) {
			out := Day10Part2(r.input)

			if out != r.expected {
				t.Errorf("Expected %q, got %q", r.expected, out)
			}
		})
	}
}

const day10Input = "83,0,193,1,254,237,187,40,88,27,2,255,149,29,42,100"
