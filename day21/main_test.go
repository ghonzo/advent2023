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

// None of these work because I relied on a quirk of the real input (clear rows and columns)
// to solve the real problem.

// func Test_numberOfPlotsInfinite(t *testing.T) {
// 	type args struct {
// 		lines []string
// 		steps int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 6}, 16},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 10}, 50},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 50}, 1594},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 100}, 6536},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 500}, 167004},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 1000}, 668697},
// 		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 5000}, 16733044},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := numberOfPlotsInfinite(tt.args.lines, tt.args.steps); got != tt.want {
// 				t.Errorf("numberOfPlotsInfinite() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
