// Advent of Code 2023, Day 6
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 6: Wait For It
// Part 1 answer: 160816
// Part 2 answer: 46561107
func main() {
	fmt.Println("Advent of Code 2023, Day 6")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	total := 1
	timeStrs := strings.Fields(entries[0][10:])
	distanceStrs := strings.Fields(entries[1][10:])
	for i := 0; i < len(timeStrs); i++ {
		time := common.Atoi(timeStrs[i])
		distanceToBeat := common.Atoi(distanceStrs[i])
		var beat int
		for hold := 1; hold < time; hold++ {
			d := (time - hold) * hold
			if d > distanceToBeat {
				beat++
			}
		}
		total *= beat
	}
	return total
}

func part2(entries []string) int {
	time := common.Atoi(strings.Join(strings.Fields(entries[0][10:]), ""))
	distanceToBeat := common.Atoi(strings.Join(strings.Fields(entries[1][10:]), ""))
	var beat int
	for hold := 1; hold < time; hold++ {
		d := (time - hold) * hold
		if d > distanceToBeat {
			beat++
		}
	}
	return beat
}
