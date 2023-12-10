// Advent of Code 2023, Day 10
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 10: Pipe Maze
// Part 1 answer: 6773
// Part 2 answer: 493
func main() {
	fmt.Println("Advent of Code 2023, Day 10")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	grid := common.ArraysGridFromLines(lines)
	startingPoint := findStartingPoint(grid)
	// Not strictly necessary, but helps us down the road
	grid.Set(startingPoint, determineActualTile(grid, startingPoint))
	pathPoints := findPath(grid, startingPoint)
	return len(pathPoints) / 2
}

func findStartingPoint(g common.Grid) common.Point {
	for p := range g.AllPoints() {
		if g.Get(p) == 'S' {
			return p
		}
	}
	panic("no starting point")
}

func determineActualTile(g common.Grid, startingPoint common.Point) byte {
	// Set of directions that feed into this space
	cardinalMap := make(map[common.Point]bool)
	// North first
	direction := common.N
	if v, ok := g.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '|' || v == '7' || v == 'F')
	}
	direction = common.E
	if v, ok := g.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '-' || v == 'J' || v == '7')
	}
	direction = common.S
	if v, ok := g.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '|' || v == 'L' || v == 'J')
	}
	// West can be derived
	var actualTile byte
	if cardinalMap[common.N] {
		if cardinalMap[common.E] {
			actualTile = 'L'
		} else if cardinalMap[common.S] {
			actualTile = '|'
		} else {
			actualTile = 'J'
		}
	} else if cardinalMap[common.E] {
		if cardinalMap[common.S] {
			actualTile = 'F'
		} else {
			actualTile = '-'
		}
	} else {
		actualTile = '7'
	}
	return actualTile
}

func findPath(g common.Grid, startingPoint common.Point) map[common.Point]bool {
	pathPoints := make(map[common.Point]bool)
	// Which direction should we start with?
	var direction common.Point
	switch g.Get(startingPoint) {
	case '|', 'L', 'J':
		direction = common.N
	case '-', 'F':
		direction = common.E
	case '7':
		direction = common.S
	}
	currentPoint := startingPoint
	for !pathPoints[currentPoint] {
		pathPoints[currentPoint] = true
		currentPoint = currentPoint.Add(direction)
		direction = findNextDirection(g.Get(currentPoint), direction)
	}
	return pathPoints
}

func findNextDirection(tile byte, lastDir common.Point) common.Point {
	switch tile {
	case '|':
		if lastDir == common.N {
			return common.N
		} else {
			return common.S
		}
	case '-':
		if lastDir == common.E {
			return common.E
		} else {
			return common.W
		}
	case 'L':
		if lastDir == common.S {
			return common.E
		} else {
			return common.N
		}
	case 'J':
		if lastDir == common.S {
			return common.W
		} else {
			return common.N
		}
	case '7':
		if lastDir == common.E {
			return common.S
		} else {
			return common.W
		}
	case 'F':
		if lastDir == common.N {
			return common.E
		} else {
			return common.S
		}
	}
	panic("invalid path")
}

func part2(lines []string) int {
	grid := common.ArraysGridFromLines(lines)
	startingPoint := findStartingPoint(grid)
	// Not strictly necessary, but helps us down the road
	grid.Set(startingPoint, determineActualTile(grid, startingPoint))
	pathPoints := findPath(grid, startingPoint)
	// Now we need to walk East from the left side and find the number of intersections
	// Fortunately, that's the way AllPoints() returns them
	interior := false
	var interiorPoints int
	var lastElbow byte
	for p := range grid.AllPoints() {
		// Is it part of the path?
		if pathPoints[p] {
			v := grid.Get(p)
			switch v {
			case '|':
				interior = !interior
			case 'F', 'L':
				lastElbow = v
			case 'J':
				if lastElbow == 'F' {
					interior = !interior
				}
				lastElbow = v
			case '7':
				if lastElbow == 'L' {
					interior = !interior
				}
				lastElbow = v
			}
		} else if interior {
			interiorPoints++
		}
	}
	return interiorPoints
}
