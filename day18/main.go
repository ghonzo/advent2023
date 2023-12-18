// Advent of Code 2023, Day 18
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 18:
// Part 1 answer: 62365
// Part 2 answer: 159485361249806
func main() {
	fmt.Println("Advent of Code 2023, Day 18")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

var dirMap = map[byte]common.Point{'U': common.U, 'R': common.R, 'D': common.D, 'L': common.L}

func part1(lines []string) int {
	g := common.NewSparseGrid()
	var p common.Point
	// Set the trench
	for _, line := range lines {
		parts := strings.Fields(line)
		dir := dirMap[parts[0][0]]
		for i := 0; i < common.Atoi(parts[1]); i++ {
			g.Set(p, '#')
			p = p.Add(dir)
		}
	}
	return calculateArea(g)
}

func calculateArea(g common.Grid) int {
	// Now calculate the contained area. Start by figuring out the bounds
	mmX := new(common.MaxMin[int])
	mmY := new(common.MaxMin[int])
	for p := range g.AllPoints() {
		mmX.Accept(p.X())
		mmY.Accept(p.Y())
	}
	var total int
	for y := mmY.Min; y <= mmY.Max; y++ {
		interior := false
		var lastElbow byte
		for x := mmX.Min; x <= mmX.Max; x++ {
			spaceType := determineSpaceType(g, x, y)
			switch spaceType {
			case '|':
				interior = !interior
			case 'F', 'L':
				lastElbow = spaceType
			case 'J':
				if lastElbow == 'F' {
					interior = !interior
				}
				lastElbow = spaceType
			case '7':
				if lastElbow == 'L' {
					interior = !interior
				}
				lastElbow = spaceType
			}
			if spaceType != '.' || interior {
				total++
			}
		}
	}
	return total
}

func present(g common.Grid, p common.Point) bool {
	v, ok := g.CheckedGet(p)
	return ok && v != 0
}

func determineSpaceType(g common.Grid, x, y int) byte {
	p := common.NewPoint(x, y)
	if !present(g, p) {
		return '.'
	}
	switch {
	case !present(g, p.Add(common.L)) && !present(g, p.Add(common.R)):
		return '|'
	case present(g, p.Add(common.L)) && present(g, p.Add(common.R)):
		return '-'
	case present(g, p.Add(common.R)) && present(g, p.Add(common.D)):
		return 'F'
	case present(g, p.Add(common.L)) && present(g, p.Add(common.D)):
		return '7'
	case present(g, p.Add(common.L)) && present(g, p.Add(common.U)):
		return 'J'
	case present(g, p.Add(common.R)) && present(g, p.Add(common.U)):
		return 'L'
	}
	panic("ugh")
}

var dirMap2 = map[byte]common.Point{'0': common.R, '1': common.D, '2': common.L, '3': common.U}

func part2(lines []string) int {
	var vertices []common.Point
	var p common.Point
	var total int
	for _, line := range lines {
		index := strings.Index(line, "#")
		dir := dirMap2[line[index+6]]
		length, _ := strconv.ParseInt(line[index+1:index+6], 16, 0)
		vertices = append(vertices, p)
		// Edges count!
		total += int(length)
		p = p.Add(dir.Times(int(length)))
	}
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		total += (vertices[i].Y() + vertices[j].Y()) * (vertices[i].X() - vertices[j].X())
	}
	return total/2 + 1
}
