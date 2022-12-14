package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type Grid aoc.Grid[int]

func (g *Grid) apply(line aoc.Line) {
	for k := range line.Points() {
		(*g)[k.X][k.Y]++
	}
}

func (g *Grid) count() int {
	total := 0
	for _, row := range *g {
		for _, c := range row {
			if c >= 2 {
				total++
			}
		}
	}
	return total
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	// partA := true
	partA := false

	lines := []aoc.Line{}
	max := 0
	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " ")
		a := aoc.ParseSplit(arr[0], ",")
		b := aoc.ParseSplit(arr[2], ",")

		if partA && a[0] != b[0] && a[1] != b[1] {
			continue
		}

		_, m := aoc.MinMax(aoc.MaxSlice(a), aoc.MaxSlice(b))
		if m > max {
			max = m
		}
		lines = append(lines, aoc.Line{A: aoc.NewPoint(a[0], a[1]), B: aoc.NewPoint(b[0], b[1])})
	}
	size := max + 1
	grid := make(Grid, 0, size)
	for i := 0; i < size; i++ {
		grid = append(grid, make([]int, size))
	}
	for _, l := range lines {
		grid.apply(l)
	}
	fmt.Println(grid.count())
}
