// Advent of Code 2023, Day 8
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2023/common"
)

// Day 8: Haunted Wasteland
// Part 1 answer: 19783
// Part 2 answer: 9177460370549
func main() {
	fmt.Println("Advent of Code 2023, Day 8")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

var nodeRegex = regexp.MustCompile(`(...) = \((...), (...)\)`)

func part1(lines []string) int {
	var steps int
	nodeMap := make(map[string][2]string)
	for _, line := range lines[2:] {
		group := nodeRegex.FindStringSubmatch(line)
		nodeMap[group[1]] = [2]string{group[2], group[3]}
	}
	curNode := "AAA"
	for {
		for _, dir := range lines[0] {
			steps++
			nextNodes := nodeMap[curNode]
			if dir == 'L' {
				curNode = nextNodes[0]
			} else {
				curNode = nextNodes[1]
			}
			if curNode == "ZZZ" {
				return steps
			}
		}
	}
}

func part2(lines []string) int {
	nodeMap := make(map[string][2]string)
	for _, line := range lines[2:] {
		group := nodeRegex.FindStringSubmatch(line)
		nodeMap[group[1]] = [2]string{group[2], group[3]}
	}
	// Find all the starting nodes
	var curNodes []string
	for node := range nodeMap {
		if node[2] == 'A' {
			curNodes = append(curNodes, node)
		}
	}
	var allSteps []int = make([]int, len(curNodes))
	// Figure out the number of steps for each node
outer:
	for i, curNode := range curNodes {
		for {
			for _, dir := range lines[0] {
				allSteps[i]++
				nextNodes := nodeMap[curNode]
				if dir == 'L' {
					curNode = nextNodes[0]
				} else {
					curNode = nextNodes[1]
				}
				if curNode[2] == 'Z' {
					continue outer
				}
			}
		}
	}
	// The answer is going to be the least common multiples of all the individual steps
	return common.LCM(allSteps[0], allSteps[1], allSteps[2:]...)
}
