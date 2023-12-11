// Advent of Code 2023, Day 11
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 11: Cosmic Expansion
// Part 1 answer: 9509330
// Part 2 answer: 635832237682
func main() {
	fmt.Println("Advent of Code 2023, Day 11")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	return findSolution(g, 2)
}

func part2(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	return findSolution(g, 1000000)
}

func findSolution(g common.Grid, expansionFactor int) int {
	// Figure out empty rows and columns
	var emptyRows, emptyColumns []int
	size := g.Size()
outerRows:
	for y := 0; y < size.Y(); y++ {
		for x := 0; x < size.X(); x++ {
			if g.Get(common.NewPoint(x, y)) != '.' {
				continue outerRows
			}
		}
		// Blank row
		emptyRows = append(emptyRows, y)
	}
outerColumns:
	for x := 0; x < size.X(); x++ {
		for y := 0; y < size.Y(); y++ {
			if g.Get(common.NewPoint(x, y)) != '.' {
				continue outerColumns
			}
		}
		// Blank column
		emptyColumns = append(emptyColumns, x)
	}
	// Now find all the galaxy points
	var galaxyPoints []common.Point
	for p := range g.AllPoints() {
		if g.Get(p) == '#' {
			galaxyPoints = append(galaxyPoints, adjustPoint(p, emptyRows, emptyColumns, expansionFactor))
		}
	}
	// Now sum all the lengths
	var total int
	for _, p1 := range galaxyPoints {
		for _, p2 := range galaxyPoints {
			if p1 != p2 {
				total += p1.Sub(p2).ManhattanDistance()
			}
		}
	}
	// We counted every path twice ... sue me
	return total / 2
}

func adjustPoint(p common.Point, emptyRows, emptyColums []int, expansionFactor int) common.Point {
	var deltaX, deltaY int
	for _, row := range emptyRows {
		if row < p.Y() {
			deltaY++
		}
	}
	for _, col := range emptyColums {
		if col < p.X() {
			deltaX++
		}
	}
	return p.Add(common.NewPoint(deltaX*(expansionFactor-1), deltaY*(expansionFactor-1)))
}
