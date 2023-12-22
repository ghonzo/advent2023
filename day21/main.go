// Advent of Code 2023, Day 21
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 21: Step Counter
// Part 1 answer: 3562
// Part 2 answer: 592723929260582
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

type growthPoint struct {
	step  int
	plots int
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
	var growthPoints []growthPoint
	// Find the points where we extend beyond the initial grid, twice
	for i := 1; ; i++ {
		for p := range sourcePoints {
			for p2 := range p.SurroundingCardinals() {
				if destPoints[p2] {
					continue
				}
				if v := g.Get(common.NewPoint(posMod(p2.X(), size.X()), posMod(p2.Y(), size.Y()))); v == '.' {
					destPoints[p2] = true
				}
			}
		}
		// Did we reach a growthPoint?
		if i == size.X()/2 || i == 3*size.X()/2 || i == 5*size.X()/2 {
			growthPoints = append(growthPoints, growthPoint{i, len(destPoints)})
			// Do we have 3?
			if len(growthPoints) == 3 {
				return extrapolate(growthPoints, steps)
			}
		}
		sourcePoints = destPoints
		destPoints = make(map[common.Point]bool)
	}
}

func posMod(a, b int) int {
	v := a % b
	if v < 0 {
		return v + b
	}
	return v
}

// Used https://users.rowan.edu/~hassen/NumerAnalysis/Interpolation_and_Approximation.pdf to fit a quadratic equation
func extrapolate(gp []growthPoint, steps int) int {
	l0 := (steps - gp[1].step) * (steps - gp[2].step) / ((gp[0].step - gp[1].step) * (gp[0].step - gp[2].step))
	l1 := (steps - gp[0].step) * (steps - gp[2].step) / ((gp[1].step - gp[0].step) * (gp[1].step - gp[2].step))
	l2 := (steps - gp[0].step) * (steps - gp[1].step) / ((gp[2].step - gp[0].step) * (gp[2].step - gp[1].step))
	return gp[0].plots*l0 + gp[1].plots*l1 + gp[2].plots*l2
}
