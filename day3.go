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
