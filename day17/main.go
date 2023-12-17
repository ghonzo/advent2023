// Advent of Code 2023, Day 17
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
	"github.com/oleiade/lane/v2"
)

// Day 17: Clumsy Crucible
// Part 1 answer: 942
// Part 2 answer: 1082
func main() {
	fmt.Println("Advent of Code 2023, Day 17")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type state struct {
	pos     common.Point
	prevDir common.Point
	// Number of times we've move this direction
	times int
}

type continueFlag int

const (
	FINISHED continueFlag = iota
	ABORT
	CONTINUE
)

func part1(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	finishPoint := g.Size().Add(common.NW)
	return findShortestPath(g, func(s state) continueFlag {
		if s.pos == finishPoint {
			return FINISHED
		} else {
			return CONTINUE
		}
	}, func(s state) []common.Point {
		var nextDirections []common.Point
		for _, dir := range []common.Point{common.R, common.D, common.L, common.U} {
			if dir == s.prevDir.Reflect() {
				continue
			}
			if dir == s.prevDir && s.times == 3 {
				continue
			}
			nextDirections = append(nextDirections, dir)
		}
		return nextDirections
	})
}

// Return minimum heat loss
func findShortestPath(g common.Grid, completeFn func(s state) continueFlag, nextMovesFn func(s state) []common.Point) int {
	// Stores the minimum heat loss we've seen at each point
	minHeatLoss := make(map[state]int)
	// State, with accumulated heat loss as our scoring function
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{}, 0)
	for !pq.Empty() {
		curState, heatLoss, _ := pq.Pop()
		cf := completeFn(curState)
		if cf == FINISHED {
			return heatLoss
		}
		if cf == ABORT {
			continue
		}
		// Have we been here before?
		if mhl, ok := minHeatLoss[curState]; ok && heatLoss >= mhl {
			// Yes, so we can forget this path
			continue
		}
		// Remember we've been here
		minHeatLoss[curState] = heatLoss
		// Now find all the next moves
		for _, dir := range nextMovesFn(curState) {
			np := curState.pos.Add(dir)
			if heatLossChar, ok := g.CheckedGet(np); ok {
				newHeatLoss := int(heatLossChar - '0')
				newState := state{pos: np}
				if curState.prevDir == dir {
					newState.times = curState.times + 1
				} else {
					newState.times = 1
				}
				newState.prevDir = dir
				pq.Push(newState, heatLoss+newHeatLoss)
			}
		}
	}
	panic("rats")
}

func part2(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	finishPoint := g.Size().Add(common.NW)
	return findShortestPath(g, func(s state) continueFlag {
		if s.pos == finishPoint {
			if s.times < 4 {
				return ABORT
			} else {
				return FINISHED
			}
		}
		return CONTINUE
	}, func(s state) []common.Point {
		if s.times > 0 && s.times < 4 {
			return []common.Point{s.prevDir}
		}
		var nextDirections []common.Point
		for _, dir := range []common.Point{common.R, common.D, common.L, common.U} {
			if dir == s.prevDir.Reflect() {
				continue
			}
			if dir == s.prevDir && s.times == 10 {
				continue
			}
			nextDirections = append(nextDirections, dir)
		}
		return nextDirections
	})
}
