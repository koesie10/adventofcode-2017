package adventofcode2017

func Day9(input string) (int, int) {
	var delim []rune

	push := func(r rune) {
		delim = append(delim, r)
	}

	pull := func() {
		delim = delim[:len(delim)-1]
	}

	var score = 1
	var sum int
	var removed int

	for _, c := range input {
		if len(delim) == 0 {
			push(c)
			continue
		}

		state := delim[len(delim)-1]

		if state == '!' {
			pull()
			continue
		}

		if c == '!' {
			push('!')
			continue
		}

		if state == '<' {
			if c == '>' {
				pull()
			} else {
				removed++
			}
			continue
		}

		if state == '{' {
			if c == ',' {
				continue
			}
		}

		if c == '}' {
			pull()
			sum += score
			score--
			continue
		}

		if c == '{' {
			push('{')
			score++
			continue
		} else if c == '<' {
			push('<')
			continue
		}

		panic("unknown character " + string(c))
	}

	return sum, removed
}
