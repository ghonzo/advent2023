// Advent of Code 2023, Day 17
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
	"github.com/oleiade/lane/v2"
)

// Day 17: Clumsy Crucible
// Part 1 answer: 942
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 17")
	lines := common.ReadStringsFromFile("input.txt")
	//lines := common.ReadStringsFromFile("testdata/example.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type state struct {
	pos      common.Point
	prevDir  []common.Point
	heatLoss int
}

func (s state) validNextDirections() []common.Point {
	var nextDirections []common.Point
	skipDirection := determineSkipDirection(s.prevDir)
	for _, dir := range []common.Point{common.R, common.D, common.L, common.U} {
		if dir == skipDirection {
			continue
		}
		// Also can't go back the way we came
		if len(s.prevDir) > 0 && dir == s.prevDir[len(s.prevDir)-1].Reflect() {
			continue
		}
		nextDirections = append(nextDirections, dir)
	}
	return nextDirections
}

func part1(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	return findShortestPath(g)
}

type memoState struct {
	pos                common.Point
	lastDirectionsHash int
}

var dirHash = map[common.Point]int{common.U: 1, common.R: 2, common.D: 4, common.L: 8}

func (s state) memo() memoState {
	var h int
	size := len(s.prevDir)
	for i := 0; i < min(size, 3); i++ {
		h *= 16
		h += dirHash[s.prevDir[size-i-1]]
	}
	return memoState{s.pos, h}
}

// Return minimum heat loss
func findShortestPath(g common.Grid) int {
	finishPoint := g.Size().Add(common.NW)
	// Stores the minimum heat loss we've seen at each point
	minHeatLoss := make(map[memoState]int)
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{}, 0)
	for !pq.Empty() {
		curState, _, _ := pq.Pop()
		// Did we finish?
		if curState.pos == finishPoint {
			//printSolution(g, curState)
			return curState.heatLoss
		}
		ms := curState.memo()
		// Have we been here before?
		if mhl, ok := minHeatLoss[ms]; ok && curState.heatLoss >= mhl {
			// Yes, so we can forget this path
			continue
		}
		// Remember we've been here
		minHeatLoss[ms] = curState.heatLoss
		// Now find all the next moves, but skip a direction if we've used it the last three times
		for _, dir := range curState.validNextDirections() {
			np := curState.pos.Add(dir)
			if heatLossChar, ok := g.CheckedGet(np); ok {
				newHeatLoss := curState.heatLoss + int(heatLossChar-'0')
				newState := state{pos: np, heatLoss: newHeatLoss}
				newState.prevDir = append(newState.prevDir, curState.prevDir...)
				newState.prevDir = append(newState.prevDir, dir)
				pq.Push(newState, newHeatLoss)
			}
		}
	}
	panic("rats")
}

func determineSkipDirection(prevDir []common.Point) common.Point {
	size := len(prevDir)
	if size >= 3 && prevDir[size-1] == prevDir[size-2] && prevDir[size-1] == prevDir[size-3] {
		return prevDir[size-1]
	}
	return common.Point{}
}

var dirChars = map[common.Point]byte{common.U: '^', common.R: '>', common.D: 'v', common.L: '<'}

func printSolution(g common.Grid, s state) {
	var p common.Point
	for _, dir := range s.prevDir {
		p = p.Add(dir)
		g.Set(p, dirChars[dir])
	}
	fmt.Println(common.RenderGrid(g))
}

func part2(lines []string) int {
	var total int
	return total
}
