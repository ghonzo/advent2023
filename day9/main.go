// Advent of Code 2023, Day 9
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 9: Mirage Maintenance
// Part 1 answer: 1834108701
// Part 2 answer: 993
func main() {
	fmt.Println("Advent of Code 2023, Day 9")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var total int
	for _, line := range lines {
		total += nextValue(convertToInts(line))
	}
	return total
}

func convertToInts(s string) []int {
	fields := strings.Fields(s)
	ret := make([]int, len(fields))
	for i, numStr := range fields {
		ret[i] = common.Atoi(numStr)
	}
	return ret
}

func nextValue(seq []int) int {
	var sequences [][]int = [][]int{seq}
	for i := 0; ; i++ {
		allZeroes := true
		seq = sequences[i]
		var nextSeq []int
		for j := 0; j < len(seq)-1; j++ {
			diff := seq[j+1] - seq[j]
			nextSeq = append(nextSeq, diff)
			if diff != 0 {
				allZeroes = false
			}
		}
		if allZeroes {
			// Unroll it
			var val int
			for k := i; k >= 0; k-- {
				seq = sequences[k]
				val += seq[len(seq)-1]
			}
			return val
		}
		sequences = append(sequences, nextSeq)
	}
}

func part2(lines []string) int {
	var total int
	for _, line := range lines {
		total += prevValue(convertToInts(line))
	}
	return total
}

func prevValue(seq []int) int {
	var sequences [][]int = [][]int{seq}
	for i := 0; ; i++ {
		allZeroes := true
		seq = sequences[i]
		var nextSeq []int
		for j := 0; j < len(seq)-1; j++ {
			diff := seq[j+1] - seq[j]
			nextSeq = append(nextSeq, diff)
			if diff != 0 {
				allZeroes = false
			}
		}
		if allZeroes {
			// Unroll it
			var val int
			for k := i; k >= 0; k-- {
				seq = sequences[k]
				val = seq[0] - val
			}
			return val
		}
		sequences = append(sequences, nextSeq)
	}
}
