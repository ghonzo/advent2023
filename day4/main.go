// Advent of Code 2023, Day 4
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 4: Scratchcards
// Part 1 answer: 26914
// Part 2 answer: 13080971
func main() {
	fmt.Println("Advent of Code 2023, Day 4")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for _, line := range entries {
		var cardTotal int
		colonIndex := strings.Index(line, ":")
		pipeIndex := strings.Index(line, "|")
		winning := make(map[int]bool)
		for _, numStr := range strings.Fields(line[colonIndex+1 : pipeIndex-1]) {
			winning[common.Atoi(numStr)] = true
		}
		for _, numStr := range strings.Fields(line[pipeIndex+1:]) {
			if winning[common.Atoi(numStr)] {
				if cardTotal == 0 {
					cardTotal = 1
				} else {
					cardTotal <<= 1
				}
			}
		}
		total += cardTotal
	}
	return total
}

func part2(entries []string) int {
	// First figure out the card totals
	cardTotals := make([]int, len(entries))
	for i, line := range entries {
		colonIndex := strings.Index(line, ":")
		pipeIndex := strings.Index(line, "|")
		winning := make(map[int]bool)
		for _, numStr := range strings.Fields(line[colonIndex+1 : pipeIndex-1]) {
			winning[common.Atoi(numStr)] = true
		}
		for _, numStr := range strings.Fields(line[pipeIndex+1:]) {
			if winning[common.Atoi(numStr)] {
				cardTotals[i]++
			}
		}
	}
	// Now figure out how many additional cards of each there will be
	additionalCards := make([]int, len(entries))
	for card, ct := range cardTotals {
		for i := card + 1; i <= card+ct; i++ {
			additionalCards[i] += additionalCards[card] + 1
		}
	}
	// Sum them, don't forget the original cards
	var total int
	for _, ac := range additionalCards {
		total += ac + 1
	}
	return total
}
