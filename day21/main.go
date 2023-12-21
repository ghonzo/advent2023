// Advent of Code 2023, Day 21
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 21:
// Part 1 answer: 3562
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 21")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	return numberOfPlots(lines, 64)
}

func numberOfPlots(lines []string, steps int) int {
	g := common.ArraysGridFromLines(lines)
	// Find the starting spot
	var start common.Point
	for p := range g.AllPoints() {
		if g.Get(p) == 'S' {
			start = p
			g.Set(p, '.')
			break
		}
	}
	sourcePoints := map[common.Point]bool{start: true}
	destPoints := make(map[common.Point]bool)
	for i := 0; i < steps; i++ {
		for p := range sourcePoints {
			for p2 := range p.SurroundingCardinals() {
				if v, _ := g.CheckedGet(p2); v == '.' {
					destPoints[p2] = true
				}
			}
		}
		sourcePoints = destPoints
		destPoints = make(map[common.Point]bool)
	}
	return len(sourcePoints)
}

func part2(lines []string) int {
	return numberOfPlotsInfinite(lines, 26501365)
}

func numberOfPlotsInfinite(lines []string, steps int) int {
	g := common.ArraysGridFromLines(lines)
	// Find the starting spot
	var start common.Point
	for p := range g.AllPoints() {
		if g.Get(p) == 'S' {
			start = p
			g.Set(p, '.')
			break
		}
	}
	sourcePoints := map[common.Point]bool{start: true}
	destPoints := make(map[common.Point]bool)
	size := g.Size()
	for i := 0; i < steps; i++ {
		for p := range sourcePoints {
			for p2 := range p.SurroundingCardinals() {
				if v := g.Get(common.NewPoint(posMod(p2.X(), size.X()), posMod(p2.Y(), size.Y()))); v == '.' {
					destPoints[p2] = true
				}
			}
		}
		sourcePoints = destPoints
		destPoints = make(map[common.Point]bool)
	}
	return len(sourcePoints)
}

func posMod(a, b int) int {
	v := a % b
	if v < 0 {
		return v + b
	}
	return v
}
