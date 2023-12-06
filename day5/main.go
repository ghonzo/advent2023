// Advent of Code 2023, Day 5
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 5: If You Give A Seed A Fertilizer
// Part 1 answer: 84470622
// Part 2 answer: 26714516
func main() {
	fmt.Println("Advent of Code 2023, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type seedMap struct {
	dest, source, length int
}

func part1(entries []string) int {
	seeds, almanac := parseAlmanac(entries)
	// Now find the destinations
	mm := new(common.MaxMin[int])
	for _, v := range seeds {
		for _, sms := range almanac {
			v = doMap(v, sms)
		}
		mm.Accept(v)
	}
	return mm.Min
}

// Return the seeds and all of the x-to-y maps
func parseAlmanac(entries []string) ([]int, [][]seedMap) {
	// First line is Seeds
	colonIndex := strings.Index(entries[0], ":")
	var seeds []int
	for _, seedStr := range strings.Fields(entries[0][colonIndex+1:]) {
		seeds = append(seeds, common.Atoi(seedStr))
	}
	// The almanac is a collection of all the maps
	var almanac [][]seedMap
	// This is just one of the x-to-y maps
	var sms []seedMap
	for i := 3; i < len(entries); i++ {
		line := entries[i]
		if len(line) == 0 {
			almanac = append(almanac, sms)
			sms = make([]seedMap, 0)
			i++
		} else {
			strs := strings.Fields(line)
			sms = append(sms, seedMap{dest: common.Atoi(strs[0]), source: common.Atoi(strs[1]), length: common.Atoi(strs[2])})
		}
	}
	almanac = append(almanac, sms)
	return seeds, almanac
}

func doMap(v int, sms []seedMap) int {
	for _, m := range sms {
		if v >= m.source && v < m.source+m.length {
			offset := v - m.source
			return m.dest + offset
		}
	}
	// None found
	return v
}

func part2(entries []string) int {
	seeds, almanac := parseAlmanac(entries)
	// Now reverse the maps
	slices.Reverse(almanac)
	// Just start counting up from 0 and do reverse map until we get to a valid seed
	for location := 0; ; location++ {
		v := location
		for _, sms := range almanac[1:] {
			v = doReverseMap(v, sms)
		}
		if validSeed(v, seeds) {
			return location
		}
	}
}

func doReverseMap(v int, sms []seedMap) int {
	for _, m := range sms {
		if v >= m.dest && v < m.dest+m.length {
			offset := v - m.dest
			return m.source + offset
		}
	}
	// None found
	return v
}

func validSeed(v int, seeds []int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if v >= seeds[i] && v < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}
