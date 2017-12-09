package adventofcode2017

import (
	"math"
)

func Day3Part1(input int64) int64 {
	// In which spiral are we. The bottom right corner is always the a square root: 9, 25, 49. This is increasing by
	// 2 every time. So just find the closest root and divide by 2 to get the spiral we are in.
	spiral := int64(math.Floor(math.Ceil(math.Sqrt(float64(input))) / 2.0))

	// Now find the distance from the square to the nearest axis
	var off int64
	if spiral == 0 {
		off = int64(float64(input) - math.Pow(2*float64(spiral)-1, 2))
	} else {
		off = int64(float64(input)-math.Pow(2*float64(spiral)-1, 2)) % int64(2*spiral)
	}

	// And finally calculate the number of steps needed
	return spiral + int64(math.Abs(float64(off-spiral)))
}

func Day3Part2(input int64) int64 {
	spiral := int64(math.Floor(math.Ceil(math.Sqrt(float64(input))) / 2.0))
	table := make([][]int64, spiral)

	for i := int64(0); i < spiral; i++ {
		table[i] = make([]int64, spiral)
	}

	x := spiral / 2
	y := spiral / 2
	table[y][x] = 1 // center

	type direction int
	const (
		up direction = iota
		down
		left
		right
	)

	dir := right
	var sum int64

	for {
		if sum > input {
			break
		}

		switch dir {
		case right:
			sum = table[y][x] + table[y-1][x] + table[y-1][x+1] + table[y-1][x+2]
			x++
			table[y][x] = sum

			if table[y-1][x] == 0 {
				dir = up
			}
		case left:
			sum = table[y][x] + table[y+1][x] + table[y+1][x-1] + table[y+1][x-2]
			x--
			table[y][x] = sum

			if table[y+1][x] == 0 {
				dir = down
			}
		case up:
			sum = table[y][x] + table[y][x-1] + table[y-1][x-1] + table[y-2][x-1]
			y--
			table[y][x] = sum

			if table[y][x-1] == 0 {
				dir = left
			}
		case down:
			// sum += current + one-right + one-right & one-down + one-right & two-down
			sum = table[y][x] + table[y][x+1] + table[y+1][x+1] + table[y+2][x+1]
			y++
			table[y][x] = sum

			if table[y][x+1] == 0 {
				dir = right
			}
		}
	}

	return sum
}
