// Advent of Code 2023, Day 21
package main

import (
	"testing"

	"github.com/ghonzo/advent2023/common"
)

func Test_numberOfPlots(t *testing.T) {
	type args struct {
		lines []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 6}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPlots(tt.args.lines, tt.args.steps); got != tt.want {
				t.Errorf("numberOfPlots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfPlotsInfinite(t *testing.T) {
	type args struct {
		lines []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 6}, 16},
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 10}, 50},
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 50}, 1594},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPlotsInfinite(tt.args.lines, tt.args.steps); got != tt.want {
				t.Errorf("numberOfPlotsInfinite() = %v, want %v", got, tt.want)
			}
		})
	}
}
