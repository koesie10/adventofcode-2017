package adventofcode2017

import (
	"strings"
	"strconv"
	"fmt"
	"bytes"
)

func Day10Part1(input string, listSize int) int {
	parts := strings.Split(input, ",")
	lengths := make([]int, len(parts))

	for i, v := range parts {
		length, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		lengths[i] = length
	}

	list := make([]int, listSize)

	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	var pos, skip int

	for _, length := range lengths {
		sublist := make([]int, length)

		for i := 0; i < length; i++ {
			sublist[i] = list[(pos+i)%len(list)]
		}

		sublist = reverseInts(sublist)

		for i := range sublist {
			list[(pos+i)%len(list)] = sublist[i]
		}

		pos = (pos + length + skip) % len(list)
		skip++
	}

	return list[0] * list[1]
}

func Day10Part2(input string) string {
	const listSize = 256

	lengths := make([]int, len(input))

	for i, v := range input {
		lengths[i] = int(v)
	}

	lengths = append(lengths, 17, 31, 73, 47, 23)

	list := make([]int, listSize)

	for i := 0; i < listSize; i++ {
		list[i] = i
	}

	var pos, skip int

	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			sublist := make([]int, length)

			for i := 0; i < length; i++ {
				sublist[i] = list[(pos+i)%len(list)]
			}

			sublist = reverseInts(sublist)

			for i := range sublist {
				list[(pos+i)%len(list)] = sublist[i]
			}

			pos = (pos + length + skip) % len(list)
			skip++
		}
	}

	hash := new(bytes.Buffer)

	for i := 0; i < 16; i++ {
		block := list[i*16:(i+1)*16]

		output := 0

		for _, v := range block {
			output ^= v
		}

		fmt.Fprintf(hash, "%02x", output)
	}

	return hash.String()
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}

	return append(reverseInts(input[1:]), input[0])
}
