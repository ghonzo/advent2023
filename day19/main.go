// Advent of Code 2023, Day 19
package main

import (
	"fmt"
	"maps"
	"regexp"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 19: Aplenty
// Part 1 answer: 449531
// Part 2 answer: 122756210763577
func main() {
	fmt.Println("Advent of Code 2023, Day 19")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type rule struct {
	category string
	lessThan bool
	value    int
	action   string
}

type workflow struct {
	name        string
	rules       []rule
	finalAction string
}

type part map[string]int

var ruleRegex = regexp.MustCompile(`(.)(.)(\d+):(.+)`)

func part1(lines []string) int {
	workflows := make(map[string]workflow)
	var lineNum int
	var line string
	for lineNum, line = range lines {
		if len(line) == 0 {
			break
		}
		leftIndex := strings.Index(line, "{")
		wf := workflow{name: line[:leftIndex]}
		ruleParts := strings.Split(line[leftIndex+1:len(line)-1], ",")
		for _, ruleStr := range ruleParts[:len(ruleParts)-1] {
			group := ruleRegex.FindStringSubmatch(ruleStr)
			wf.rules = append(wf.rules, rule{group[1], group[2] == "<", common.Atoi(group[3]), group[4]})
		}
		wf.finalAction = ruleParts[len(ruleParts)-1]
		workflows[wf.name] = wf
	}
	var total int
	for _, line = range lines[lineNum+1:] {
		p := make(part)
		for _, rating := range strings.Split(line[1:len(line)-1], ",") {
			p[rating[:1]] = common.Atoi(rating[2:])
		}
		total += processPart(p, workflows)
	}
	return total
}

func processPart(p part, workflows map[string]workflow) int {
	var total int
	if doWorkflow(p, workflows["in"], workflows) {
		for _, v := range p {
			total += v
		}
	}
	return total
}

func doWorkflow(p part, wf workflow, workflows map[string]workflow) bool {
	for _, r := range wf.rules {
		pv := p[r.category]
		if (r.lessThan && pv < r.value) || (!r.lessThan && pv > r.value) {
			switch r.action {
			case "A":
				return true
			case "R":
				return false
			default:
				return doWorkflow(p, workflows[r.action], workflows)
			}
		}
	}
	switch wf.finalAction {
	case "A":
		return true
	case "R":
		return false
	default:
		return doWorkflow(p, workflows[wf.finalAction], workflows)
	}
}

type ratingRange struct {
	min, max int
}

type partRatings map[string]ratingRange

func (pr partRatings) totalCombos() uint64 {
	// Through trial and error I figured out this fits ... I'm sure there's a real name for it!
	// avgOverallSum * permutations
	var permutations uint64 = 1
	for _, rr := range pr {
		permutations *= uint64(rr.max - rr.min + 1)
	}
	return permutations
}

func part2(lines []string) uint64 {
	workflows := make(map[string]workflow)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		leftIndex := strings.Index(line, "{")
		wf := workflow{name: line[:leftIndex]}
		ruleParts := strings.Split(line[leftIndex+1:len(line)-1], ",")
		for _, ruleStr := range ruleParts[:len(ruleParts)-1] {
			group := ruleRegex.FindStringSubmatch(ruleStr)
			wf.rules = append(wf.rules, rule{group[1], group[2] == "<", common.Atoi(group[3]), group[4]})
		}
		wf.finalAction = ruleParts[len(ruleParts)-1]
		workflows[wf.name] = wf
	}
	fullRange := ratingRange{1, 4000}
	allParts := partRatings{"x": fullRange, "m": fullRange, "a": fullRange, "s": fullRange}
	return doWorkflow2(allParts, workflows["in"], workflows)
}

func doWorkflow2(p partRatings, wf workflow, workflows map[string]workflow) uint64 {
	var total uint64
	for _, r := range wf.rules {
		pCopy := make(partRatings)
		maps.Copy(pCopy, p)
		if r.lessThan {
			rr := p[r.category]
			rr.max = r.value - 1
			pCopy[r.category] = rr
			rr = p[r.category]
			rr.min = r.value
			p[r.category] = rr
		} else {
			rr := p[r.category]
			rr.min = r.value + 1
			pCopy[r.category] = rr
			rr = p[r.category]
			rr.max = r.value
			p[r.category] = rr
		}
		switch r.action {
		case "A":
			total += pCopy.totalCombos()
		case "R":
			total += 0
		default:
			total += doWorkflow2(pCopy, workflows[r.action], workflows)
		}
	}
	switch wf.finalAction {
	case "A":
		total += p.totalCombos()
	case "R":
		total += 0
	default:
		total += doWorkflow2(p, workflows[wf.finalAction], workflows)
	}
	return total
}
