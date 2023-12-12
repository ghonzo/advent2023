// Advent of Code 2023, Day 12
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 12: Hot Springs
// Part 1 answer: 7922
// Part 2 answer: 18093821750095
func main() {
	fmt.Println("Advent of Code 2023, Day 12")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type gameState struct {
	// Position in record
	pos int
	// Which group we're currently matching
	groupIndex int
	// The number of broken springs we've seen so far in this group
	numInGroup int
	// If we expect (require) the next record to be a working spring
	expectWorking bool
}

func part1(lines []string) int {
	var total int
	for _, line := range lines {
		c := countAllArrangements(line)
		total += c
	}
	return total
}

func countAllArrangements(line string) int {
	fields := strings.Fields(line)
	return count(fields[0], convertToInts(fields[1]))
}

func count(conditionRecord string, damagedGroups []int) int {
	var total int
	// Value is number of times we've seen this state in our problem space
	currentStates := map[gameState]int{{}: 1}
	nextStates := map[gameState]int{}
	for len(currentStates) > 0 {
		for state, num := range currentStates {
			// Did we reach the end?
			if state.pos == len(conditionRecord) {
				// And did we match all the groups?
				if state.groupIndex == len(damagedGroups) {
					// Yes, add the number of matched states to the total
					total += num
				}
				continue
			}
			recordAtPos := conditionRecord[state.pos]
			if state.groupIndex < len(damagedGroups) && !state.expectWorking && recordAtPos != '.' {
				// We're in a group
				// Move up a space
				state.pos++
				if recordAtPos == '?' && state.numInGroup == 0 {
					// This could be working as well, so record that state
					nextStates[state] += num
				}
				state.numInGroup++
				// Are we at the end of the group?
				if state.numInGroup == damagedGroups[state.groupIndex] {
					// Advance to the next group
					state.groupIndex++
					// Reset the number
					state.numInGroup = 0
					// We need a working one next
					state.expectWorking = true
				}
				nextStates[state] += num
			} else if recordAtPos != '#' && state.numInGroup == 0 {
				// In between groups
				state.expectWorking = false
				state.pos++
				nextStates[state] += num
			}
		}
		// More states to look at
		currentStates = nextStates
		nextStates = map[gameState]int{}
	}
	return total
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
	}
	return total
}

func countAllArrangements2(line string) int {
	fields := strings.Fields(line)
	return count(strings.Join([]string{fields[0], fields[0], fields[0], fields[0], fields[0]}, "?"),
		convertToInts(strings.Join([]string{fields[1], fields[1], fields[1], fields[1], fields[1]}, ",")))
}
