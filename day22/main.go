// Advent of Code 2023, Day 22
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 22: Sand Slabs
// Part 1 answer: 454
// Part 2 answer: 74287
func main() {
	fmt.Println("Advent of Code 2023, Day 22")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type point3 [3]int

type brick struct {
	// start z-coord will always be less than or equal to end z-coord
	start, end point3
	restingOn  []*brick
}

func (b *brick) points() <-chan point3 {
	ch := make(chan point3)
	go func() {
		for x := min(b.start[0], b.end[0]); x <= max(b.start[0], b.end[0]); x++ {
			for y := min(b.start[1], b.end[1]); y <= max(b.start[1], b.end[1]); y++ {
				for z := b.start[2]; z <= b.end[2]; z++ {
					ch <- point3{x, y, z}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func sortByZ(a, b *brick) int {
	return a.start[2] - b.start[2]
}

func part1(lines []string) int {
	bricks := readBricks(lines)
	for {
		var moved bool
		// Sort the bricks by z-coord
		slices.SortFunc(bricks, sortByZ)
		// Record where every brick is in 3d space
		brickSpace := make(map[point3]*brick)
		// Now make each brick fall in z-coord
		for _, b := range bricks {
			moved = b.fall(brickSpace) || moved
		}
		// If at least one moved, do it again
		if !moved {
			break
		}
	}
	disintegrateSet := make(map[*brick]bool)
	for _, b := range bricks {
		disintegrateSet[b] = true
	}
	// We can't disintegrate bricks are supporting exactly one other brick
	for _, b := range bricks {
		if len(b.restingOn) == 1 {
			delete(disintegrateSet, b.restingOn[0])
		}
	}
	return len(disintegrateSet)
}

func readBricks(lines []string) []*brick {
	bricks := make([]*brick, len(lines))
	for i, line := range lines {
		tildeIndex := strings.Index(line, "~")
		b := &brick{start: strtop3(line[0:tildeIndex]), end: strtop3(line[tildeIndex+1:])}
		if b.start[2] > b.end[2] {
			b.start, b.end = b.end, b.start
		}
		bricks[i] = b
	}
	return bricks
}

func strtop3(s string) point3 {
	var p point3
	for i, coordStr := range strings.Split(s, ",") {
		p[i] = common.Atoi(coordStr)
	}
	return p
}

func (b *brick) fall(brickSpace map[point3]*brick) bool {
	moved := false
	// Move the start and end down unless we are on the floor
	if b.start[2] != 1 {
		// Move start and end down
		b.start[2]--
		b.end[2]--
		var restingOn []*brick
		for p := range b.points() {
			// Are we occupying the space of another brick?
			if otherBrick, found := brickSpace[p]; found && !slices.Contains(restingOn, otherBrick) {
				restingOn = append(restingOn, otherBrick)
			}
		}
		// If we are resting on at least one other brick, undo and remember
		if len(restingOn) > 0 {
			b.start[2]++
			b.end[2]++
			b.restingOn = restingOn
		} else {
			moved = true
		}
	}
	// Fill brickspace
	for p := range b.points() {
		brickSpace[p] = b
	}
	return moved
}

func part2(lines []string) int {
	bricks := readBricks(lines)
	for {
		var moved bool
		// Sort the bricks by z-coord
		slices.SortFunc(bricks, sortByZ)
		// Record where every brick is in 3d space
		brickSpace := make(map[point3]*brick)
		// Now make each brick fall in z-coord
		for _, b := range bricks {
			moved = b.fall(brickSpace) || moved
		}
		// If at least one moved, do it again
		if !moved {
			break
		}
	}
	var total int
	for _, b := range bricks {
		total += countFall(b, bricks)
	}
	return total
}

func countFall(b *brick, bricks []*brick) int {
	goneBricks := map[*brick]bool{b: true}
outer:
	for {
	inner:
		for _, b2 := range bricks {
			if !goneBricks[b2] && len(b2.restingOn) > 0 {
				// Check if all that this was resting on are gone
				for _, b3 := range b2.restingOn {
					if !goneBricks[b3] {
						continue inner
					}
				}
				// Yep, all gone, so this one should be added
				goneBricks[b2] = true
				continue outer
			}
		}
		// Didn't add any
		return len(goneBricks) - 1
	}
}
