package adventofcode2017

import (
	"strings"
	"strconv"
)

func Day12Part1(input string) int {
	graph := makeGraph(input)

	discovered := make(map[int]struct{})
	s := stack{0}
	var v int

	// DFS
	for len(s) > 0 {
		s, v = s.Pop()
		if _, ok := discovered[v]; ok {
			continue
		}

		discovered[v] = struct{}{}
		for _, e := range graph[v] {
			s = s.Push(e)
		}
	}

	return len(discovered)
}

func Day12Part2(input string) int {
	graph := makeGraph(input)
	discovered := make(map[int]struct{})
	var groups int

	for node := range graph {
		if _, ok := discovered[node]; ok {
			continue
		}

		groups++

		s := stack{node}
		var v int

		// DFS
		for len(s) > 0 {
			s, v = s.Pop()
			if _, ok := discovered[v]; ok {
				continue
			}

			discovered[v] = struct{}{}
			for _, e := range graph[v] {
				s = s.Push(e)
			}
		}
	}

	return groups
}

func makeGraph(input string) map[int][]int {
	lines := strings.Split(input, "\n")

	graph := make(map[int][]int, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "<->")

		n, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		nodeParts := strings.Split(parts[1], ",")
		connected := make([]int, len(nodeParts))
		for i, part := range nodeParts {
			v, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				panic(err)
			}

			connected[i] = v
		}

		graph[n] = connected
	}

	return graph
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}
