// Advent of Code 2023, Day 14
package main

import (
	"fmt"
	"hash/maphash"

	"github.com/ghonzo/advent2023/common"
)

// Day 14: Parabolic Reflector Dish
// Part 1 answer: 107053
// Part 2 answer: 88371
func main() {
	fmt.Println("Advent of Code 2023, Day 14")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	tilt(g, common.N)
	return calculateLoad(g)
}

func tilt(g common.Grid, dir common.Point) {
	for move(g, dir) > 0 {
	}
}

// return the number of round rocks that moved
func move(g common.Grid, dir common.Point) int {
	moved := 0
	for p := range g.AllPoints() {
		if g.Get(p) == 'O' {
			if nv, ok := g.CheckedGet(p.Add(dir)); ok && nv == '.' {
				g.Set(p.Add(dir), 'O')
				g.Set(p, '.')
				moved++
			}
		}
	}
	return moved
}

func calculateLoad(g common.Grid) int {
	var load int
	height := g.Size().Y()
	for x := 0; x < g.Size().X(); x++ {
		for y := 0; y < height; y++ {
			if g.Get(common.NewPoint(x, y)) == 'O' {
				load += height - y
			}
		}
	}
	return load
}

func part2(lines []string) int {
	goalCycle := 1000000000
	g := common.ArraysGridFromLines(lines)
	var h maphash.Hash
	var loads []int
	// hash of grid pointing to cycle number
	gridCycle := make(map[uint64]int)
	for c := 0; ; c++ {
		cycle(g)
		h.WriteString(common.RenderGrid(g))
		hash := h.Sum64()
		if prevCycle, ok := gridCycle[hash]; ok {
			// Okay we saw it before. Figure out the load for the
			cycleLength := c - prevCycle
			targetIndex := goalCycle - 1 - prevCycle
			return loads[(targetIndex%cycleLength)+prevCycle]
		}
		gridCycle[hash] = c
		loads = append(loads, calculateLoad(g))
		h.Reset()
	}
}

func cycle(g common.Grid) {
	tilt(g, common.N)
	tilt(g, common.W)
	tilt(g, common.S)
	tilt(g, common.E)
}
