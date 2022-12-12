package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type Grid [][]int

func (g *Grid) apply(line Line) {
	if line.a.x == line.b.x {
		min, max := aoc.MinMax(line.a.y, line.b.y)
		for i := min; i <= max; i++ {
			(*g)[line.a.x][i]++
		}
	} else if line.a.y == line.b.y {
		min, max := aoc.MinMax(line.a.x, line.b.x)
		for i := min; i <= max; i++ {
			(*g)[i][line.a.y]++
		}
	} else {
		for x, y := line.a.x, line.a.y; x != line.b.x || y != line.b.y; x, y = x+X(line), y+Y(line) {
			(*g)[x][y]++
		}
		(*g)[line.b.x][line.b.y]++
	}
}

func X(l Line) int { return aoc.Abs(l.a.x-l.b.x) / (l.b.x - l.a.x) }
func Y(l Line) int { return aoc.Abs(l.b.y-l.a.y) / (l.b.y - l.a.y) }

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

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	// partA := true
	partA := false

	lines := []Line{}
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
		lines = append(lines, Line{
			a: Point{a[0], a[1]},
			b: Point{b[0], b[1]},
		})
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
