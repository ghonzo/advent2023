// Advent of Code 2023, Day 13
package main

import (
	"fmt"
	"slices"

	"github.com/ghonzo/advent2023/common"
)

// Day 13: Point of Incidence
// Part 1 answer: 34772
// Part 2 answer: 35554
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
			total += scoreForPattern(common.ArraysGridFromLines(lines[lastBlank+1:n]), 0)
			lastBlank = n
		}
	}
	return total
}

// If we should ignore a score (for part 2), include it here. Otherwise 0 is fine
func scoreForPattern(g common.Grid, omitScore int) int {
	// Score (unique pattern) for each column and row
	colScore := make([]uint64, g.Size().X())
	rowScore := make([]uint64, g.Size().Y())
	for p := range g.AllPoints() {
		colScore[p.X()] <<= 1
		rowScore[p.Y()] <<= 1
		if g.Get(p) == '#' {
			colScore[p.X()]++
			rowScore[p.Y()]++
		}
	}
	if col, ok := findSymmetry(colScore, omitScore%100); ok {
		return col
	}
	if row, ok := findSymmetry(rowScore, omitScore/100); ok {
		return row * 100
	}
	return 0
}

// If omit is non-zero, ignore that line of reflection
func findSymmetry(scores []uint64, omit int) (int, bool) {
	for i := 0; i < len(scores)-1; i++ {
		if i == omit-1 {
			continue
		}
		size := min(i+1, len(scores)-i-1)
		leftSlice := scores[i+1-size : i+1]
		rightSlice := scores[i+1 : i+1+size]
		slices.Reverse(leftSlice)
		if slices.Equal(leftSlice, rightSlice) {
			return i + 1, true
		}
		// Undo the reverse
		slices.Reverse(leftSlice)
	}
	return 0, false
}

func part2(lines []string) int {
	var total int
	lastBlank := -1
	for n, line := range append(lines, "") {
		if len(line) == 0 {
			total += scoreForSmudgePattern(common.ArraysGridFromLines(lines[lastBlank+1 : n]))
			lastBlank = n
		}
	}
	return total
}

func scoreForSmudgePattern(g common.Grid) int {
	// First figure out the original score
	originalScore := scoreForPattern(g, 0)
	// Now brute force all smudge variations
	for p := range g.AllPoints() {
		v := g.Get(p)
		if v == '.' {
			g.Set(p, '#')
		} else {
			g.Set(p, '.')
		}
		if newScore := scoreForPattern(g, originalScore); newScore > 0 {
			return newScore
		}
		g.Set(p, v)
	}
	panic("none")
}
