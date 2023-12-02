// Advent of Code 2023, Day 2
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 2: Cube Conundrum
// Part 1 answer: 2285
// Part 2 answer: 77021
func main() {
	fmt.Println("Advent of Code 2023, Day 2")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

// The number cubes that make up a "pull"
type pull struct {
	r, g, b int
}

func part1(entries []string) int {
	var total int
outer:
	for i, line := range entries {
		pulls := parseGame(line)
		for _, p := range pulls {
			if p.r > 12 || p.g > 13 || p.b > 14 {
				continue outer
			}
		}
		// Possible
		total += i + 1
	}
	return total
}

// Match the number and the first letter of the color
var cubeRegex = regexp.MustCompile(`(\d+) (.).+`)

func parseGame(line string) []pull {
	var ret []pull
	colonIndex := strings.IndexByte(line, ':')
	for _, str := range strings.Split(line[colonIndex+1:], ";") {
		var p pull
		for _, cube := range strings.Split(str, ", ") {
			group := cubeRegex.FindStringSubmatch(cube)
			// How many cubes
			n := common.Atoi(group[1])
			// What color
			switch group[2] {
			case "r":
				p.r = n
			case "g":
				p.g = n
			case "b":
				p.b = n
			}
		}
		ret = append(ret, p)
	}
	return ret
}

func part2(entries []string) int {
	var total int
	for _, line := range entries {
		pulls := parseGame(line)
		var minpull pull
		for _, p := range pulls {
			minpull.r = max(minpull.r, p.r)
			minpull.g = max(minpull.g, p.g)
			minpull.b = max(minpull.b, p.b)
		}
		// Possible
		total += minpull.r * minpull.g * minpull.b
	}
	return total
}
