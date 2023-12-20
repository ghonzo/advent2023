// Advent of Code 2023, Day 20
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 20: Pulse Propagation
// Part 1 answer: 818723272
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2023, Day 20")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type moduleDef struct {
	name string
	// True if conjunction (%) otherwise false (%)
	conjunction bool
	outputs     []string
}

type moduleState struct {
	flipFlop  bool
	conInputs map[string]bool
}

func (ms *moduleState) allHigh() bool {
	for _, v := range ms.conInputs {
		if !v {
			return false
		}
	}
	return true
}

func part1(lines []string) int {
	moduleDefMap := make(map[string]moduleDef)
	for _, line := range lines {
		var md moduleDef
		parts := strings.Split(line, " -> ")
		if parts[0] == "broadcaster" {
			md.name = parts[0]
		} else {
			md.name = parts[0][1:]
			if parts[0][0] == '&' {
				md.conjunction = true
			}
		}
		md.outputs = strings.Split(parts[1], ", ")
		moduleDefMap[md.name] = md
	}
	// Now create all the moduleStates
	moduleStateMap := make(map[string]*moduleState)
	for _, md := range moduleDefMap {
		moduleStateMap[md.name] = &moduleState{conInputs: make(map[string]bool)}
	}
	// Now link all the conjunction outputs
	for _, md := range moduleDefMap {
		for _, otherMd := range md.outputs {
			if ms, ok := moduleStateMap[otherMd]; ok {
				ms.conInputs[md.name] = false
			}
		}
	}
	var totalLow, totalHigh int
	for i := 0; i < 1000; i++ {
		low, high := pressButton(moduleDefMap, moduleStateMap)
		totalLow += low
		totalHigh += high
	}
	return totalLow * totalHigh
}

type pulse struct {
	origin, dest string
	high         bool
}

func pressButton(moduleDefMap map[string]moduleDef, moduleStateMap map[string]*moduleState) (int, int) {
	var low, high int
	// Button sends low to broadcaster
	low++
	var pulses []pulse
	// Do the broadcaster first
	for _, output := range moduleDefMap["broadcaster"].outputs {
		pulses = append(pulses, pulse{origin: "broadcaster", dest: output})
		low++
	}
	for len(pulses) > 0 {
		p := pulses[0]
		pulses = pulses[1:]
		md, ok := moduleDefMap[p.dest]
		if !ok {
			continue
		}
		ms := moduleStateMap[p.dest]
		var outputHigh bool
		if md.conjunction {
			// Record what we received
			ms.conInputs[p.origin] = p.high
			outputHigh = !ms.allHigh()
		} else if p.high {
			// A flip-flop module received a high pulse, which means nothing happens
			continue
		} else {
			// A flip-flop module received a low pulse
			ms.flipFlop = !ms.flipFlop
			outputHigh = ms.flipFlop
		}
		// Send it to all outputs
		for _, output := range md.outputs {
			pulses = append(pulses, pulse{origin: md.name, dest: output, high: outputHigh})
			if outputHigh {
				high++
			} else {
				low++
			}
		}
	}
	return low, high
}

func part2(lines []string) int {
	moduleDefMap := make(map[string]moduleDef)
	for _, line := range lines {
		var md moduleDef
		parts := strings.Split(line, " -> ")
		if parts[0] == "broadcaster" {
			md.name = parts[0]
		} else {
			md.name = parts[0][1:]
			if parts[0][0] == '&' {
				md.conjunction = true
			}
		}
		md.outputs = strings.Split(parts[1], ", ")
		moduleDefMap[md.name] = md
	}
	// Now create all the moduleStates
	moduleStateMap := make(map[string]*moduleState)
	for _, md := range moduleDefMap {
		moduleStateMap[md.name] = &moduleState{conInputs: make(map[string]bool)}
	}
	// Now link all the conjunction outputs
	for _, md := range moduleDefMap {
		for _, otherMd := range md.outputs {
			if ms, ok := moduleStateMap[otherMd]; ok {
				ms.conInputs[md.name] = false
			}
		}
	}
	for i := 0; ; i++ {
		if pressButton2(moduleDefMap, moduleStateMap) {
			return i + 1
		}
	}
}

func pressButton2(moduleDefMap map[string]moduleDef, moduleStateMap map[string]*moduleState) bool {
	// Button sends low to broadcaster
	var pulses []pulse
	// Do the broadcaster first
	for _, output := range moduleDefMap["broadcaster"].outputs {
		pulses = append(pulses, pulse{origin: "broadcaster", dest: output})
	}
	for len(pulses) > 0 {
		p := pulses[0]
		if p.dest == "rx" && !p.high {
			return true
		}
		pulses = pulses[1:]
		md, ok := moduleDefMap[p.dest]
		if !ok {
			continue
		}
		ms := moduleStateMap[p.dest]
		var outputHigh bool
		if md.conjunction {
			// Record what we received
			ms.conInputs[p.origin] = p.high
			outputHigh = !ms.allHigh()
		} else if p.high {
			// A flip-flop module received a high pulse, which means nothing happens
			continue
		} else {
			// A flip-flop module received a low pulse
			ms.flipFlop = !ms.flipFlop
			outputHigh = ms.flipFlop
		}
		// Send it to all outputs
		for _, output := range md.outputs {
			pulses = append(pulses, pulse{origin: md.name, dest: output, high: outputHigh})
		}
	}
	return false
}
