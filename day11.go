package adventofcode2017

import (
	"strings"
)

func Day11Part1(input string) (int, int) {
	steps := strings.Split(input, ",")

	directions := map[string]vector3{
		"se": {1, -1, 0},
		"ne": {1, 0, -1},
		"n":  {0, 1, -1},
		"nw": {-1, 1, 0},
		"sw": {-1, 0, 1},
		"s":  {0, -1, 1},
	}

	origin := vector3{0, 0, 0}

	var pos vector3
	var furthest int

	for _, step := range steps {
		pos = pos.add(directions[step])

		dist := pos.dist(origin)
		if dist > furthest {
			furthest = dist
		}
	}

	return pos.dist(origin), furthest
}

type vector3 [3]int

func (v vector3) add(p vector3) vector3 {
	return [3]int{
		v[0] + p[0],
		v[1] + p[1],
		v[2] + p[2],
	}
}

func (v vector3) dist(p vector3) int {
	return (AbsInt(v[0]-p[0]) + AbsInt(v[1]-p[1]) + AbsInt(v[2]-p[2])) / 2
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
