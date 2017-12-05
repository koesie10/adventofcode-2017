package adventofcode2017

import (
	"math"
	"strconv"
	"strings"
)

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")
	var checksum int64

	for _, line := range lines {
		rows := strings.Split(line, "\t")

		var min int64 = math.MaxInt64
		var max int64 = math.MinInt64

		for _, v := range rows {
			row, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}

			if row > max {
				max = row
			}
			if row < min {
				min = row
			}
		}

		checksum += max - min
	}

	return int(checksum)
}

func Day2Part2(input string) int {
	lines := strings.Split(input, "\n")
	var checksum int64

	for _, line := range lines {
		rows := strings.Split(line, "\t")

		values := make([]int64, len(rows))
		for i, v := range rows {
			var err error
			values[i], err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
		}

	Line:
		for _, v1 := range values {
			for _, v2 := range values {
				if v1 == v2 {
					continue
				}

				if v1%v2 == 0 {
					checksum += v1 / v2
					break Line
				}
				if v2%v1 == 0 {
					checksum += v2 / v1
					break Line
				}
			}
		}
	}

	return int(checksum)
}
