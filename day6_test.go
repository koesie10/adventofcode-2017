package adventofcode2017

import "testing"

func TestDay6Example(t *testing.T) {
	input := "0\t2\t7\t0"
	expectedCycles := 5
	expectedRecount := 4
	cycles, recount := Day6(input)

	if cycles != expectedCycles {
		t.Errorf("Expected cycles to be %v, got %v for input %q", expectedCycles, cycles, input)
	}

	if recount != expectedRecount {
		t.Errorf("Expected recount to be %v, got %v for input %q", expectedRecount, recount, input)
	}
}

func TestDay6Input(t *testing.T) {
	input := day6Input
	expectedCycles := 6681
	expectedRecount := 2392
	cycles, recount := Day6(input)

	if cycles != expectedCycles {
		t.Errorf("Expected cycles to be %v, got %v for input %q", expectedCycles, cycles, input)
	}

	if recount != expectedRecount {
		t.Errorf("Expected recount to be %v, got %v for input %q", expectedRecount, recount, input)
	}
}

const day6Input = `4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3`
