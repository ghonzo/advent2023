// Advent of Code 2023, Day 10
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 10:
// Part 1 answer: 6773
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 10")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	grid := common.ArraysGridFromLines(lines)
	// Find starting point
	var startingPoint common.Point
	for startingPoint = range grid.AllPoints() {
		if grid.Get(startingPoint) == 'S' {
			break
		}
	}
	// First, figure out which direction we can follow from the starting position
	lastDirection, nextDirection := findFirstDirections(grid, startingPoint)
	// Okay now we can just starting following it and remembering the points
	currentPoint := startingPoint.Add(lastDirection)
	var allPoints []common.Point
	for currentPoint != startingPoint {
		allPoints = append(allPoints, currentPoint)
		currentPoint = currentPoint.Add(nextDirection)
		nextDirection = findNextDirection(grid.Get(currentPoint), nextDirection)
	}
	return (len(allPoints) + 1) / 2
}

func findFirstDirections(g common.Grid, startingPoint common.Point) (lastDirection common.Point, nextDirection common.Point) {
	lastDirection = common.N
	p := startingPoint.Add(lastDirection)
	v, ok := g.CheckedGet(p)
	if ok {
		switch v {
		case '|':
			nextDirection = common.N
			return
		case '7':
			nextDirection = common.W
			return
		case 'F':
			nextDirection = common.E
			return
		}
	}
	lastDirection = common.E
	p = startingPoint.Add(lastDirection)
	v, ok = g.CheckedGet(p)
	if ok {
		switch v {
		case '-':
			nextDirection = common.E
			return
		case 'J':
			nextDirection = common.N
			return
		case '7':
			nextDirection = common.S
			return
		}
	}
	lastDirection = common.S
	p = startingPoint.Add(lastDirection)
	v, ok = g.CheckedGet(p)
	if ok {
		switch v {
		case '|':
			nextDirection = common.S
			return
		case 'J':
			nextDirection = common.W
			return
		case 'L':
			nextDirection = common.E
			return
		}
	}
	panic("nope")
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
	return common.Point{}
}

func part2(lines []string) int {
	// First, we need to replace the "S" with the appropirate other shape
	grid := common.ArraysGridFromLines(lines)
	// Find starting point
	var startingPoint common.Point
	for startingPoint = range grid.AllPoints() {
		if grid.Get(startingPoint) == 'S' {
			break
		}
	}
	// Find which directions feed into this space
	cardinalMap := make(map[common.Point]bool)
	// North first
	direction := common.N
	if v, ok := grid.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '|' || v == '7' || v == 'F')
	}
	direction = common.E
	if v, ok := grid.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '-' || v == 'J' || v == '7')
	}
	direction = common.S
	if v, ok := grid.CheckedGet(startingPoint.Add(direction)); ok {
		cardinalMap[direction] = (v == '|' || v == 'L' || v == 'J')
	}
	// West can be derived
	var actualShape byte
	if cardinalMap[common.N] {
		if cardinalMap[common.E] {
			actualShape = 'L'
		} else if cardinalMap[common.S] {
			actualShape = '|'
		} else {
			actualShape = 'J'
		}
	} else if cardinalMap[common.E] {
		if cardinalMap[common.S] {
			actualShape = 'F'
		} else {
			actualShape = '-'
		}
	} else {
		actualShape = '7'
	}
	// Now replace it in the grid
	grid.Set(startingPoint, actualShape)
	// Put all the points of the path in a map
	pathPoints := make(map[common.Point]bool)
	pathPoints[startingPoint] = true
	lastDirection, nextDirection := findFirstDirections(grid, startingPoint)
	// Okay now we can just starting following it and remembering the points
	currentPoint := startingPoint.Add(lastDirection)
	for currentPoint != startingPoint {
		pathPoints[currentPoint] = true
		currentPoint = currentPoint.Add(nextDirection)
		nextDirection = findNextDirection(grid.Get(currentPoint), nextDirection)
	}
	// Now we need to walk East from the left side and find the number of intersections
	// Fortunately, that's the way AllPoints() returns them
	interior := false
	var interiorPoints int
	var lastElbow byte
	for p := range grid.AllPoints() {
		v := grid.Get(p)
		// Is it part of the path?
		if pathPoints[p] {
			switch v {
			case '|':
				interior = !interior
			case 'F', 'L':
				lastElbow = v
			case 'J':
				if lastElbow == 'F' {
					interior = !interior
				}
				lastElbow = 0
			case '7':
				if lastElbow == 'L' {
					interior = !interior
				}
				lastElbow = 0
			}
		} else if interior {
			interiorPoints++
		}
	}
	return interiorPoints
}
