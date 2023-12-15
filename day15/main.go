// Advent of Code 2023, Day 15
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 15: Lens Library
// Part 1 answer: 495972
// Part 2 answer: 245223
func main() {
	fmt.Println("Advent of Code 2023, Day 15")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var total int
	for _, step := range strings.Split(lines[0], ",") {
		total += hash(step)
	}
	return total
}

func hash(s string) int {
	var h int
	for _, b := range []byte(s) {
		h += int(b)
		h *= 17
		h %= 256
	}
	return h
}

type lens struct {
	label       string
	focalLength int
}

func part2(lines []string) int {
	var boxes [][]lens = make([][]lens, 256)
	for _, step := range strings.Split(lines[0], ",") {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			box := hash(label)
			boxes[box] = removeLens(boxes[box], label)
		} else {
			label := step[:len(step)-2]
			box := hash(label)
			boxes[box] = addLens(boxes[box], label, int(step[len(step)-1]-'0'))
		}
	}
	var total int
	for boxNum, box := range boxes {
		for slot, l := range box {
			total += (boxNum + 1) * (slot + 1) * l.focalLength
		}
	}
	return total
}

func removeLens(box []lens, label string) []lens {
	return slices.DeleteFunc(box, func(l lens) bool {
		return l.label == label
	})
}

func addLens(box []lens, label string, focalLength int) []lens {
	newLens := lens{label: label, focalLength: focalLength}
	for i, l := range box {
		if l.label == label {
			box[i] = newLens
			return box
		}
	}
	// Add it to the end
	return append(box, newLens)
}
