package adventofcode2017

import (
	"strings"
	"strconv"
)

func Day5(input string, part2 bool) int {
	lines := strings.Split(input, "\n")
	instructions := make([]int, len(lines))

	for i, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		instructions[i] = v
	}

	pointer := 0
	steps := 0

	for {
		if pointer < 0 || pointer >= len(instructions) {
			break
		}

		offset := instructions[pointer]

		if part2 && offset >= 3 {
			instructions[pointer]--
		} else {
			instructions[pointer]++
		}

		pointer += offset

		steps++
	}

	return steps
}
