// Advent of Code 2023, Day 11
package main

import (
	"testing"

	"github.com/ghonzo/advent2023/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 374},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolution(t *testing.T) {
	type args struct {
		g               common.Grid
		expansionFactor int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"10x", args{common.ArraysGridFromLines(common.ReadStringsFromFile("testdata/example.txt")), 10}, 1030},
		{"100x", args{common.ArraysGridFromLines(common.ReadStringsFromFile("testdata/example.txt")), 100}, 8410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSolution(tt.args.g, tt.args.expansionFactor); got != tt.want {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
