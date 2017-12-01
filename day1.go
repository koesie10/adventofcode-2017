package adventofcode2017

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

func Day1Part2(input string) int {
	if len(input) % 2 != 0 {
		panic("invalid input, is odd")
	}

	half := len(input) / 2

	// just duplicate the input so we avoid case distinction later on
	doubleInput := []byte(input + input)

	var sum int

	for i := 0; i < len([]byte(input)); i++ {
		c := doubleInput[i]
		look := doubleInput[i + half]

		if c == look {
			sum += int(c - '0')
		}
	}

	return sum
}
