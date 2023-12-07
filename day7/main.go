// Advent of Code 2023, Day 7
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ghonzo/advent2023/common"
)

// Day 7: Camel Cards
// Part 1 answer: 246795406
// Part 2 answer: 249356515
func main() {
	fmt.Println("Advent of Code 2023, Day 7")
	lines := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type entry struct {
	hand string
	bid  int
}

// Hand ranks in increasing order
const (
	high = iota
	onePair
	twoPair
	three
	fullHouse
	four
	five
)

func part1(lines []string) int {
	var entries []entry
	for _, line := range lines {
		groups := strings.Fields(line)
		entries = append(entries, entry{hand: groups[0], bid: common.Atoi(groups[1])})
	}
	sort.Slice(entries, func(i, j int) bool {
		handTypeLeft := classifyHand(entries[i].hand)
		handTypeRight := classifyHand(entries[j].hand)
		if handTypeLeft != handTypeRight {
			return handTypeLeft < handTypeRight
		}
		// Same rank, so sort by first different card value
		for p := 0; p < 5; p++ {
			cardValueLeft := cardValue(entries[i].hand[p])
			cardValueRight := cardValue(entries[j].hand[p])
			if cardValueLeft != cardValueRight {
				return cardValueLeft < cardValueRight
			}
		}
		panic("equal hands")

	})
	// Now they are sorted in ascending order of strength
	var total int
	for i, e := range entries {
		total += (i + 1) * e.bid
	}
	return total
}

func classifyHand(hand string) int {
	cards := make(map[byte]int)
	for _, c := range []byte(hand) {
		cards[c]++
	}
	var hasThree bool
	var twos int
	// Five and four
	for _, v := range cards {
		if v == 5 {
			return five
		}
		if v == 4 {
			return four
		}
		if v == 3 {
			hasThree = true
		}
		if v == 2 {
			twos++
		}
	}
	if hasThree && twos == 1 {
		return fullHouse
	}
	if hasThree {
		return three
	}
	if twos == 2 {
		return twoPair
	}
	if twos == 1 {
		return onePair
	}
	return high
}

func cardValue(b byte) int {
	switch b {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(b - '0')
	}
}

func part2(lines []string) int {
	var entries []entry
	for _, line := range lines {
		groups := strings.Fields(line)
		entries = append(entries, entry{hand: groups[0], bid: common.Atoi(groups[1])})
	}
	sort.Slice(entries, func(i, j int) bool {
		handTypeLeft := classifyHandWithJokers(entries[i].hand)
		handTypeRight := classifyHandWithJokers(entries[j].hand)
		if handTypeLeft != handTypeRight {
			return handTypeLeft < handTypeRight
		}
		// Same rank, so sort by first different card value
		for p := 0; p < 5; p++ {
			cardValueLeft := cardValueWithJokers(entries[i].hand[p])
			cardValueRight := cardValueWithJokers(entries[j].hand[p])
			if cardValueLeft != cardValueRight {
				return cardValueLeft < cardValueRight
			}
		}
		panic("equal hands")

	})
	// Now they are sorted in ascending order of strength
	var total int
	for i, e := range entries {
		total += (i + 1) * e.bid
	}
	return total
}
func classifyHandWithJokers(hand string) int {
	cards := make(map[byte]int)
	for _, c := range []byte(hand) {
		cards[c]++
	}
	var hasThree bool
	var twos int
	// Five and four
	for k, v := range cards {
		if v == 5 {
			return five
		}
		if v == 4 {
			if cards['J'] == 1 || k == 'J' {
				return five
			}
			return four
		}
		if v == 3 {
			hasThree = true
		}
		if v == 2 {
			twos++
		}
	}
	switch cards['J'] {
	case 3:
		if twos == 1 {
			return five
		}
		return four
	case 2:
		if hasThree {
			return five
		}
		if twos == 2 {
			return four
		}
		return three
	case 1:
		if hasThree {
			return four
		}
		if twos == 2 {
			return fullHouse
		}
		if twos == 1 {
			return three
		}
		return onePair
	}
	// No jokers
	if hasThree && twos == 1 {
		return fullHouse
	}
	if hasThree {
		return three
	}
	if twos == 2 {
		return twoPair
	}
	if twos == 1 {
		return onePair
	}
	return high
}

func cardValueWithJokers(b byte) int {
	switch b {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		return int(b - '0')
	}
}
