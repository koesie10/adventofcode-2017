package adventofcode2017

import "testing"

func TestDay13Part1(t *testing.T) {
	runs := []struct {
		name     string
		input    string
		expected int
	}{
		{"Example", "0: 3\n1: 2\n4: 4\n6: 4", 24},
		{"Input", day13Input, 2164},
	}

	for _, r := range runs {
		t.Run(r.name, func(t *testing.T) {
			out := Day13Part1(r.input)

			if out != r.expected {
				t.Errorf("Expected %v, got %v", r.expected, out)
			}
		})
	}
}

func TestDay13Part2(t *testing.T) {
	runs := []struct {
		name     string
		input    string
		expected int
	}{
		{"Example", "0: 3\n1: 2\n4: 4\n6: 4", 10},
		//{"Input", day13Input, 0},
	}

	for _, r := range runs {
		t.Run(r.name, func(t *testing.T) {
			out := Day13Part2(r.input)

			if out != r.expected {
				t.Errorf("Expected %v, got %v", r.expected, out)
			}
		})
	}
}

const day13Input = `0: 3
1: 2
2: 4
4: 4
6: 5
8: 6
10: 6
12: 8
14: 6
16: 6
18: 9
20: 8
22: 8
24: 8
26: 12
28: 8
30: 12
32: 12
34: 12
36: 10
38: 14
40: 12
42: 10
44: 8
46: 12
48: 14
50: 12
52: 14
54: 14
56: 14
58: 12
62: 14
64: 12
66: 12
68: 14
70: 14
72: 14
74: 17
76: 14
78: 18
84: 14
90: 20
92: 14`
