package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/cbebe/aoc"
)

const partA bool = true

type RangeGrid struct {
	grid aoc.Grid[int]
	maxx int
	minx int
	maxy int
	miny int
}

func NewRangeGrid(maxx, minx, maxy, miny int) RangeGrid {
	g := RangeGrid{make(aoc.Grid[int], 0, maxy-miny+1), maxx, minx, maxy, miny}
	for i := 0; i < maxy-miny+1; i++ {
		g.grid = append(g.grid, make([]int, maxx-minx+1))
	}
	return g
}

func (g *RangeGrid) GetMutRangeCell(x, y int) *int {
	return g.grid.GetMutCell(x-g.minx, y-g.miny)
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	Run(lines)
}

func PrintGrid(g aoc.Grid[int]) {
	for _, r := range g {
		for _, c := range r {
			var b rune
			switch c {
			case -1:
				b = '#'
			case 0:
				b = '.'
			case 1:
				b = 'B'
			case 2:
				b = 'S'
			}
			fmt.Print(string(b))
		}
		fmt.Println()
	}
}

func Run(lines []string) {
	g, bs := ParseGrid(lines)
	for _, v := range bs {
		v.MarkEmpty(&g)
	}
	y := 10
	total := 0
	for x := g.minx; x < g.maxx; x++ {
		if c := g.GetMutRangeCell(x, y); c != nil && *c == -1 {
			total++
		}
	}
	PrintGrid(g.grid)
	fmt.Println(total)
}

type BS struct {
	b aoc.Point
	s aoc.Point
}

func MarkCell(g *RangeGrid, x, y int) {
	if c := g.GetMutRangeCell(x, y); c != nil && *c == 0 {
		*c = -1
	}
}

func MarkNeighbours(g *RangeGrid, x, y int) {
	MarkCell(g, x+1, y)
	MarkCell(g, x-1, y)
	MarkCell(g, x, y+1)
	MarkCell(g, x, y-1)
}

func (bs BS) MarkEmpty(g *RangeGrid) {
	x, y := bs.s.X, bs.s.Y
	max := aoc.Abs(bs.b.X-x) + aoc.Abs(bs.b.Y-y)
	for i := 1; i <= max; i++ {
		a, b, c, d := aoc.NewPoint(x+i, y), aoc.NewPoint(x, y+i), aoc.NewPoint(x-i, y), aoc.NewPoint(x, y-i)
		m, n, o, p := aoc.NewPoint(x+i, y), aoc.NewPoint(x, y-i), aoc.NewPoint(x-i, y), aoc.NewPoint(x, y+i)
		l1, l2, l3, l4 := aoc.NewLine(a, b).Points(), aoc.NewLine(c, d).Points(), aoc.NewLine(m, n).Points(), aoc.NewLine(o, p).Points()
		for _, c := range append(append(append(l1, l2...), l3...), l4...) {
			MarkCell(g, c.X, c.Y)
		}
	}
}

func ParseGrid(lines []string) (RangeGrid, []BS) {
	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0
	arr := make([]BS, 0, len(lines))
	for _, v := range lines {
		if v == "" {
			continue
		}
		var sx, sy, bx, by int
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		miny, _ = aoc.MinMaxInt(sy, by, miny)
		minx, _ = aoc.MinMaxInt(sx, bx, minx)
		_, maxx = aoc.MinMaxInt(sx, bx, maxx)
		_, maxy = aoc.MinMaxInt(sy, by, maxy)

		b := aoc.NewPoint(bx, by)
		s := aoc.NewPoint(sx, sy)
		arr = append(arr, BS{b, s})
	}
	g := NewRangeGrid(maxx, minx, maxy, miny)
	fmt.Println("make grid ok")
	for _, bs := range arr {
		*g.GetMutRangeCell(bs.b.X, bs.b.Y) = 1
		*g.GetMutRangeCell(bs.s.X, bs.s.Y) = 2
	}

	fmt.Println("make grid ok")
	return g, arr
}
