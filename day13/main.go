// Advent of Code 2023, Day 13
package main

import (
	"fmt"
	"slices"

	"github.com/ghonzo/advent2023/common"
)

// Day 13: Point of Incidence
// Part 1 answer: 34772
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 13")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var total int
	lastBlank := -1
	for n, line := range append(lines, "") {
		if len(line) == 0 {
			total += scoreForPattern(lines[lastBlank+1 : n])
			lastBlank = n
		}
	}
	return total
}

func scoreForPattern(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	// Score -> sum of column
	colScore := make([]uint64, g.Size().X())
	for x := 0; x < g.Size().X(); x++ {
		var score uint64
		for y := 0; y < g.Size().Y(); y++ {
			score <<= 1
			if g.Get(common.NewPoint(x, y)) == '#' {
				score++
			}
		}
		colScore[x] = score
	}
	if col, ok := findSymmetry(colScore); ok {
		return col
	}
	// Let's look for rows
	rowScore := make([]uint64, g.Size().Y())
	for y := 0; y < g.Size().Y(); y++ {
		var score uint64
		for x := 0; x < g.Size().X(); x++ {
			score <<= 1
			if g.Get(common.NewPoint(x, y)) == '#' {
				score++
			}
		}
		rowScore[y] = score
	}
	if row, ok := findSymmetry(rowScore); ok {
		return row * 100
	}
	panic("none")
}

func findSymmetry(scores []uint64) (int, bool) {
	for i := 0; i < len(scores)-1; i++ {
		size := min(i+1, len(scores)-i-1)
		leftSlice := scores[i+1-size : i+1]
		rightSlice := scores[i+1 : i+1+size]
		slices.Reverse(leftSlice)
		if slices.Equal(leftSlice, rightSlice) {
			return i + 1, true
		}
		slices.Reverse(leftSlice)
	}
	return 0, false
}

func part2(lines []string) int {
	var total int
	return total
}
