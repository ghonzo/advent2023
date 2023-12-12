// Advent of Code 2023, Day 12
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 12:
// Part 1 answer: 7922
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 12")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type groupsAndSum struct {
	groups []int
	sum    int
}

func part1(lines []string) int {
	var total int
	for _, line := range lines {
		c := countAllArrangements(line)
		total += c
		//fmt.Println(line, c)
	}
	return total
}

func countAllArrangements(line string) int {
	parts := strings.Fields(line)
	groups := convertToInts(parts[1])
	sum := 0
	for _, n := range groups {
		sum += n
	}
	return count(parts[0], groupsAndSum{groups, sum})
}

func count(springs string, g groupsAndSum) int {
	unknownIndex := strings.Index(springs, "?")
	if unknownIndex == -1 {
		if validGroup(springs, g) {
			//fmt.Println(springs)
			return 1
		} else {
			return 0
		}
	}
	if !validPartialGroup(springs, g) {
		return 0
	}
	// If there's no way we could make it, stop it too
	if strings.Count(springs, "#")+len(springs)-unknownIndex < g.sum {
		return 0
	}
	return count(springs[0:unknownIndex]+"."+springs[unknownIndex+1:], g) + count(springs[0:unknownIndex]+"#"+springs[unknownIndex+1:], g)
}

func validGroup(springs string, g groupsAndSum) bool {
	if strings.Count(springs, "#") != g.sum {
		return false
	}
	groupIndex := 0
	groupCount := 0
	for _, r := range springs {
		if r == '.' && groupCount > 0 {
			if groupIndex >= len(g.groups) || g.groups[groupIndex] != groupCount {
				return false
			}
			groupCount = 0
			groupIndex++
		} else if r == '#' {
			groupCount++
		}
	}
	return (groupIndex == len(g.groups) && groupCount == 0) || (groupIndex == len(g.groups)-1 && g.groups[groupIndex] == groupCount)
}

func validPartialGroup(springs string, g groupsAndSum) bool {
	if strings.Count(springs, "#") > g.sum {
		return false
	}
	groupIndex := 0
	groupCount := 0
	for _, r := range springs {
		if r == '.' && groupCount > 0 {
			if groupIndex >= len(g.groups) || g.groups[groupIndex] != groupCount {
				return false
			}
			groupCount = 0
			groupIndex++
		} else if r == '#' {
			groupCount++
		} else if r == '?' {
			break
		}
	}
	return true
}

func convertToInts(s string) []int {
	fields := strings.Split(s, ",")
	ret := make([]int, len(fields))
	for i, numStr := range fields {
		ret[i] = common.Atoi(numStr)
	}
	return ret
}

func part2(lines []string) int {
	var total int
	for _, line := range lines {
		c := countAllArrangements2(line)
		total += c
		fmt.Println(line, c)
	}
	return total
}

func countAllArrangements2(line string) int {
	parts := strings.Fields(line)
	origGroups := convertToInts(parts[1])
	// 2
	groups := append(origGroups, origGroups...)
	// 4
	groups = append(groups, groups...)
	// 5
	groups = append(groups, origGroups...)
	sum := 0
	for _, n := range groups {
		sum += n
	}
	return count(strings.Repeat(parts[0]+"?", 4)+parts[0], groupsAndSum{groups, sum})
}
