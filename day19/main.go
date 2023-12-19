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

func (p part) sum() int {
	var total int
	for _, v := range p {
		total += v
	}
	return total
}

func part1(lines []string) int {
	workflows := make(map[string]workflow)
	var lineNum int
	var line string
	for lineNum, line = range lines {
		if len(line) == 0 {
			break
		}
		wf := parseWorkflow(line)
		workflows[wf.name] = wf
	}
	var total int
	for _, line = range lines[lineNum+1:] {
		p := parsePart(line)
		total += processPart(p, workflows)
	}
	return total
}

var ruleRegex = regexp.MustCompile(`(.)(.)(\d+):(.+)`)

func parseWorkflow(s string) workflow {
	leftBraceIndex := strings.Index(s, "{")
	wf := workflow{name: s[:leftBraceIndex]}
	ruleParts := strings.Split(s[leftBraceIndex+1:len(s)-1], ",")
	for _, ruleStr := range ruleParts[:len(ruleParts)-1] {
		group := ruleRegex.FindStringSubmatch(ruleStr)
		wf.rules = append(wf.rules, rule{group[1], group[2] == "<", common.Atoi(group[3]), group[4]})
	}
	wf.finalAction = ruleParts[len(ruleParts)-1]
	return wf
}

func parsePart(s string) part {
	p := make(part)
	for _, rating := range strings.Split(s[1:len(s)-1], ",") {
		p[rating[:1]] = common.Atoi(rating[2:])
	}
	return p
}

func processPart(p part, workflows map[string]workflow) int {
	if doWorkflow(p, workflows["in"], workflows) {
		return p.sum()
	} else {
		return 0
	}
}

// Returns true if accepted, false if rejected
func doWorkflow(p part, wf workflow, workflows map[string]workflow) bool {
	for _, r := range wf.rules {
		pv := p[r.category]
		if (r.lessThan && pv < r.value) || (!r.lessThan && pv > r.value) {
			return processAction(p, r.action, workflows)
		}
	}
	// Else...
	return processAction(p, wf.finalAction, workflows)
}

func processAction(p part, action string, workflows map[string]workflow) bool {
	switch action {
	case "A":
		return true
	case "R":
		return false
	default:
		return doWorkflow(p, workflows[action], workflows)
	}
}

// PART 2 STUFF BELOW

type ratingRange struct {
	min, max int
}

type partRatings map[string]ratingRange

func (pr partRatings) permutations() uint64 {
	var p uint64 = 1
	for _, rr := range pr {
		p *= uint64(rr.max - rr.min + 1)
	}
	return p
}

func part2(lines []string) uint64 {
	workflows := make(map[string]workflow)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		wf := parseWorkflow(line)
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
		rr := p[r.category]
		if r.lessThan {
			pCopy[r.category] = ratingRange{rr.min, r.value - 1}
			p[r.category] = ratingRange{r.value, rr.max}
		} else {
			pCopy[r.category] = ratingRange{r.value + 1, rr.max}
			p[r.category] = ratingRange{rr.min, r.value}
		}
		total += processAction2(pCopy, r.action, workflows)
	}
	total += processAction2(p, wf.finalAction, workflows)
	return total
}

func processAction2(p partRatings, action string, workflows map[string]workflow) uint64 {
	switch action {
	case "A":
		return p.permutations()
	case "R":
		return 0
	default:
		return doWorkflow2(p, workflows[action], workflows)
	}
}
