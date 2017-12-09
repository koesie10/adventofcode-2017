package adventofcode2017

import (
	"strings"
	"sort"
)

func Day4(input string, validFunc func(string) bool) int64 {
	var count int64

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if validFunc(line) {
			count++
		}
	}

	return count
}

func Day4Part1IsValid(input string) bool {
	words := strings.Split(input, " ")
	wordMap := make(map[string]struct{}, len(words))

	for _, word := range words {
		if _, ok := wordMap[word]; ok {
			return false
		}

		wordMap[word] = struct{}{}
	}

	return true
}

func Day4Part2IsValid(input string) bool {
	words := strings.Split(input, " ")
	wordMap := make(map[string]struct{}, len(words))

	for _, word := range words {
		r := []rune(word)
		sort.Sort(sortRunes(r))
		word = string(r)

		if _, ok := wordMap[word]; ok {
			return false
		}

		wordMap[word] = struct{}{}
	}

	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
