package adventofcode2017

import (
	"strings"
	"strconv"
)

func Day13Part1(input string) int {
	lines := strings.Split(input, "\n")
	type layerSpec struct {
		depth     int
		scanRange int
	}

	var maxDepth int

	layerSpecs := make([]layerSpec, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")

		depth, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		scanRange, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(err)
		}

		layerSpecs[i] = layerSpec{depth, scanRange}

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	type layer struct {
		scanRange int
		pos       int
		down      bool
	}

	layers := make([]*layer, maxDepth+1)

	for _, v := range layerSpecs {
		layers[v.depth] = &layer{v.scanRange, 0, true}
	}

	var severity int

	for depth, scanner := range layers {
		if scanner != nil {
			if scanner.pos == 0 {
				severity += depth * scanner.scanRange
			}
		}

		for _, s := range layers {
			if s == nil {
				continue
			}

			if s.down {
				s.pos++
			} else {
				s.pos--
			}

			if s.pos == 0 {
				s.down = true
			} else if s.pos == s.scanRange-1 {
				s.down = false
			}
		}
	}

	return severity
}

func Day13Part2(input string) int {
	lines := strings.Split(input, "\n")
	type layerSpec struct {
		depth     int
		scanRange int
	}

	var maxDepth int

	layerSpecs := make([]layerSpec, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")

		depth, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		scanRange, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(err)
		}

		layerSpecs[i] = layerSpec{depth, scanRange}

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	type layer struct {
		scanRange int
		pos       int
		down      bool
	}

	var delay int

	next:
	for {
		layers := make([]*layer, maxDepth+1)

		simulate := func() {
			for _, s := range layers {
				if s == nil {
					continue
				}

				if s.down {
					s.pos++
				} else {
					s.pos--
				}

				if s.pos == 0 {
					s.down = true
				} else if s.pos == s.scanRange-1 {
					s.down = false
				}
			}
		}

		for _, v := range layerSpecs {
			layers[v.depth] = &layer{v.scanRange, 0, true}
		}

		for i := 0; i < delay; i++ {
			simulate()
		}

		for _, scanner := range layers {
			if scanner != nil {
				if scanner.pos == 0 {
					delay++
					continue next
				}
			}

			simulate()
		}

		return delay
	}

	return 0
}
