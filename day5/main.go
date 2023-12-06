// Advent of Code 2023, Day 5
package main

import (
	"fmt"
	"slices"
	"sort"
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
	// First line is Seeds
	colonIndex := strings.Index(entries[0], ":")
	var seeds []int
	for _, seedStr := range strings.Fields(entries[0][colonIndex+1:]) {
		seeds = append(seeds, common.Atoi(seedStr))
	}
	var allMaps [][]seedMap
	var sms []seedMap
	for i := 3; i < len(entries); i++ {
		line := entries[i]
		if len(line) == 0 {
			allMaps = append(allMaps, sms)
			sms = make([]seedMap, 0)
			i++
		} else {
			var sm seedMap
			strs := strings.Fields(line)
			sm.dest = common.Atoi(strs[0])
			sm.source = common.Atoi(strs[1])
			sm.length = common.Atoi(strs[2])
			sms = append(sms, sm)
		}
	}
	allMaps = append(allMaps, sms)
	// Now find the destinations
	mm := new(common.MaxMin[int])
	for _, v := range seeds {
		for _, sms := range allMaps {
			v = doMap(v, sms)
		}
		mm.Accept(v)
	}
	return mm.Min
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
	// First line is Seeds
	colonIndex := strings.Index(entries[0], ":")
	var seeds []int
	for _, seedStr := range strings.Fields(entries[0][colonIndex+1:]) {
		seeds = append(seeds, common.Atoi(seedStr))
	}
	var allMaps [][]seedMap
	var sms []seedMap
	for i := 3; i < len(entries); i++ {
		line := entries[i]
		if len(line) == 0 {
			allMaps = append(allMaps, sms)
			sms = make([]seedMap, 0)
			i++
		} else {
			var sm seedMap
			strs := strings.Fields(line)
			sm.dest = common.Atoi(strs[0])
			sm.source = common.Atoi(strs[1])
			sm.length = common.Atoi(strs[2])
			sms = append(sms, sm)
		}
	}
	allMaps = append(allMaps, sms)
	// Now reverse the maps
	slices.Reverse(allMaps)
	// Now start with the destinations and go backwards until we find seeds
	locationMap := allMaps[0]
	sort.Slice(locationMap, func(i, j int) bool {
		return locationMap[i].dest < locationMap[j].dest
	})
	locationMap = append([]seedMap{{dest: 0, source: 0, length: locationMap[0].dest}}, locationMap...)
	for _, sm := range locationMap {
		// Do the backwards map until we find a valid seed
		for location := sm.dest; location < sm.dest+sm.length; location++ {
			v := location
			for _, sms := range allMaps[1:] {
				v = doReverseMap(v, sms)
			}
			if validSeed(v, seeds) {
				return location
			}

		}
	}
	panic("derp")
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
