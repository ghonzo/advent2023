// Advent of Code 2023, Day 3
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 3: Gear Ratios
// Part 1 answer: 539433
// Part 2 answer: 75847567
func main() {
	fmt.Println("Advent of Code 2023, Day 3")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	size := grid.Size()
	for y := 0; y < size.Y(); y++ {
		// The value of the number "in progress"
		var number int
		// True if the current number in progress has been found to be adjacent to a symbol
		var adjacent bool
		for x := 0; x < size.X(); x++ {
			p := common.NewPoint(x, y)
			b := grid.Get(p)
			if b >= '0' && b <= '9' {
				// update number
				number = number*10 + int(b-'0')
				// have we found an adjacent yet?
				if !adjacent {
					// Look at all the adjacent points
					for ap := range p.SurroundingPoints() {
						ab, found := grid.CheckedGet(ap)
						if found && !(ab >= '0' && ab <= '9') && ab != '.' {
							adjacent = true
						}
					}
				}
			} else {
				if adjacent {
					total += number
				}
				number = 0
				adjacent = false
			}
		}
		if adjacent {
			total += number
		}
	}
	return total
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	size := grid.Size()
	// Key is the gear points, values are all the adjacent numbers
	gears := make(map[common.Point][]int)
	for y := 0; y < size.Y(); y++ {
		var number int
		// Set of all the gear points around the number in progress
		gearPointsSet := make(map[common.Point]bool)
		for x := 0; x < size.X(); x++ {
			p := common.NewPoint(x, y)
			b := grid.Get(p)
			if b >= '0' && b <= '9' {
				// update number
				number = number*10 + int(b-'0')
				// Look at all the adjacent points to find any gears
				for ap := range p.SurroundingPoints() {
					ab, found := grid.CheckedGet(ap)
					if found && ab == '*' {
						gearPointsSet[ap] = true
					}
				}
			} else {
				// Add the number to all the gear points
				for k := range gearPointsSet {
					gears[k] = append(gears[k], number)
				}
				clear(gearPointsSet)
				number = 0
			}
		}
		for k := range gearPointsSet {
			gears[k] = append(gears[k], number)
		}
	}
	// Now find all gear points that have exactly 2 adjacent numbers
	for _, v := range gears {
		if len(v) == 2 {
			total += v[0] * v[1]
		}
	}
	return total
}
