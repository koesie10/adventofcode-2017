package adventofcode2017

import (
	"strings"
	"strconv"
	"bytes"
)

func Day6(input string) (int, int) {
	blockStrings := strings.Split(input, "\t")
	blocks := make([]int, len(blockStrings))
	for i, str := range blockStrings {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		blocks[i] = v
	}

	keyGenerator := func(input []int) string {
		var result bytes.Buffer
		for _, v := range input {
			result.WriteRune(rune(v))
			result.WriteRune('\t')
		}

		return result.String()
	}

	cycles := 0
	seen := make(map[string]int)
	for {
		key := keyGenerator(blocks)
		if count, ok := seen[key]; ok {
			return cycles, cycles - count
		}

		seen[key] = cycles

		var max int
		var maxIndex int
		for i, v := range blocks {
			if v > max {
				max = v
				maxIndex = i
			}
		}

		blocks[maxIndex] = 0

		for i := 1; i <= max; i++ {
			blocks[(maxIndex + i) % len(blocks)]++
		}

		cycles++
	}
}
