// Advent of Code 2023, Day 1
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 1: Trebuchet?!
// Part 1 answer: 55386
// Part 2 answer: 54824
func main() {
	fmt.Println("Advent of Code 2023, Day 1")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for _, line := range entries {
		var first, last int
		for _, c := range []byte(line) {
			if c > '0' && c <= '9' {
				last = int(c - '0')
				if first == 0 {
					first = last
				}
			}
		}
		total += first*10 + last
	}
	return total
}

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2(entries []string) int {
	var total int
	for _, line := range entries {
		var first, last int
		for i := 0; i < len(line); i++ {
			var d int
			c := line[i]
			if c > '0' && c <= '9' {
				d = int(c - '0')
			} else {
				for di, ds := range digits {
					if strings.HasPrefix(line[i:], ds) {
						d = di + 1
						break
					}
				}
			}
			if d > 0 {
				last = d
				if first == 0 {
					first = d
				}
			}
		}
		total += first*10 + last
	}
	return total
}
