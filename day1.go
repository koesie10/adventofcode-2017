package main

func Day1Part1(input string) int {
	var sum int

	for i, c := range []byte(input) {
		var next byte
		if i + 1 == len(input) {
			next = input[0]
		} else {
			next = input[i + 1]
		}

		if c == next {
			sum += int(c - '0')
		}
	}

	return sum
}
