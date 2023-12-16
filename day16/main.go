// Advent of Code 2023, Day 16
package main

import (
	"fmt"

	"github.com/ghonzo/advent2023/common"
)

// Day 16: The Floor Will Be Lava
// Part 1 answer: 8551
// Part 2 answer: 8754
func main() {
	fmt.Println("Advent of Code 2023, Day 16")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type beam struct {
	pos common.Point
	dir common.Point
}

func part1(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	return calcEnergized(g, beam{pos: common.NewPoint(-1, 0), dir: common.R})
}

var forwardSlashMirror = map[common.Point]common.Point{common.U: common.R, common.R: common.U, common.D: common.L, common.L: common.D}
var backSlashMirror = map[common.Point]common.Point{common.U: common.L, common.R: common.D, common.D: common.R, common.L: common.U}

func calcEnergized(g common.Grid, initialBeam beam) int {
	energized := make(map[common.Point]bool)
	beamsSeen := make(map[beam]bool)
	beams := []beam{initialBeam}
	for len(beams) > 0 {
		// Pop off the top
		b := beams[0]
		beams = beams[1:]
		// Move
		b.pos = b.pos.Add(b.dir)
		v, ok := g.CheckedGet(b.pos)
		if !ok || beamsSeen[b] {
			continue
		}
		energized[b.pos] = true
		beamsSeen[b] = true
		switch v {
		case '/':
			b.dir = forwardSlashMirror[b.dir]
		case '\\':
			b.dir = backSlashMirror[b.dir]
		case '|':
			if b.dir == common.R || b.dir == common.L {
				beams = append(beams, beam{pos: b.pos, dir: common.U}, beam{pos: b.pos, dir: common.D})
				continue
			}
		case '-':
			if b.dir == common.U || b.dir == common.D {
				beams = append(beams, beam{pos: b.pos, dir: common.R}, beam{pos: b.pos, dir: common.L})
				continue
			}
		}
		beams = append(beams, b)
	}
	return len(energized)
}

func part2(lines []string) int {
	g := common.ArraysGridFromLines(lines)
	var maxEnergized int
	for x := 0; x < g.Size().X(); x++ {
		maxEnergized = max(maxEnergized, calcEnergized(g, beam{pos: common.NewPoint(x, -1), dir: common.D}),
			calcEnergized(g, beam{pos: common.NewPoint(x, g.Size().Y()), dir: common.U}))
	}
	for y := 0; y < g.Size().Y(); y++ {
		maxEnergized = max(maxEnergized, calcEnergized(g, beam{pos: common.NewPoint(-1, y), dir: common.R}),
			calcEnergized(g, beam{pos: common.NewPoint(g.Size().X(), y), dir: common.L}))
	}
	return maxEnergized
}
