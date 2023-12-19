// Advent of Code 2023, Day 19
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 19:
// Part 1 answer: 449531
// Part 2 answer:
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

func part2(lines []string) int {
	return 0
}
