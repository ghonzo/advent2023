// Advent of Code 2023, Day 20
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 20: Pulse Propagation
// Part 1 answer: 818723272
// Part 2 answer: 243902373381257
func main() {
	fmt.Println("Advent of Code 2023, Day 20")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type moduleDef struct {
	name string
	// True if conjunction (&) otherwise flip-flop (%)
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
	moduleDefMap := initializeModuleDefs(lines)
	moduleStateMap := initializeModuleStates(moduleDefMap)
	var totalLow, totalHigh int
	for i := 0; i < 1000; i++ {
		low, high := pressButton(moduleDefMap, moduleStateMap)
		totalLow += low
		totalHigh += high
	}
	return totalLow * totalHigh
}

// Returns name pointing to module definition
func initializeModuleDefs(lines []string) map[string]moduleDef {
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
	return moduleDefMap
}

// Returns name pointing to mutable module state
func initializeModuleStates(moduleDefMap map[string]moduleDef) map[string]*moduleState {
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
	return moduleStateMap
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

/*
 * So at least in my input, the goal is the output of a conjunction module ... it looks like:
 *
 * &cn -> rx
 *
 * So I need to find the cycle times (when it sends a high pulse) for each of the inputs to that
 * conjunction module and then the LCM of those cycle times will be when rx receives a low pulse
 */
func part2(lines []string) int {
	moduleDefMap := initializeModuleDefs(lines)
	moduleStateMap := initializeModuleStates(moduleDefMap)
	// Find the conjection module that outputs to "rx"
	var m moduleDef
	for _, m = range moduleDefMap {
		if m.outputs[0] == "rx" {
			break
		}
	}
	var cycleTimes []int
	// Now find all the modules that feed into *that* one and find cycle times
	for name := range moduleStateMap[m.name].conInputs {
		cycleTimes = append(cycleTimes, findCycleTime(name, m.name, moduleDefMap, initializeModuleStates(moduleDefMap)))
	}
	return common.LCM(cycleTimes[0], cycleTimes[1], cycleTimes[2:]...)
}

func findCycleTime(origin, dest string, moduleDefMap map[string]moduleDef, moduleStateMap map[string]*moduleState) int {
	for cycle := 1; ; cycle++ {
		// Button sends low to broadcaster
		var pulses []pulse
		// Do the broadcaster first
		for _, output := range moduleDefMap["broadcaster"].outputs {
			pulses = append(pulses, pulse{origin: "broadcaster", dest: output})
		}
		for len(pulses) > 0 {
			p := pulses[0]
			if p.origin == origin && p.dest == dest && p.high {
				return cycle
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
	}
}
