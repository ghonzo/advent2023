// Advent of Code 2023, Day 18
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 18: Lavaduct Lagoon
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
	var vertices []common.Point
	var p common.Point
	var total int
	for _, line := range lines {
		parts := strings.Fields(line)
		dir := dirMap[parts[0][0]]
		length := common.Atoi(parts[1])
		vertices = append(vertices, p)
		// Edges count!
		total += int(length)
		p = p.Add(dir.Times(int(length)))
	}
	// Trapezoid forumula for the area of a polygon
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		total += (vertices[i].Y() + vertices[j].Y()) * (vertices[i].X() - vertices[j].X())
	}
	return total/2 + 1
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
	// Trapezoid forumula for the area of a polygon
	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		total += (vertices[i].Y() + vertices[j].Y()) * (vertices[i].X() - vertices[j].X())
	}
	return total/2 + 1
}
