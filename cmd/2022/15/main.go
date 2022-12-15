package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/cbebe/aoc"
)

var theY int

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

func LL() {
	for i := 0; i < 80; i++ {
		fmt.Print("#")
	}
	fmt.Println()
}

func main() {
	pts = make(aoc.Set[aoc.Point])
	var f, filename string
	var lines []string
	f, theY = "input.txt", 2000000
	_, filename, _, _ = runtime.Caller(0)
	lines = aoc.Lines(f, filename)
	Run(lines)
	// LL()
	// f, theY = "test.txt", 10
	// _, filename, _, _ = runtime.Caller(0)
	// lines = aoc.Lines(f, filename)
	// Run(lines)
	// LL()
	f, theY = "test-2.txt", 5
	_, filename, _, _ = runtime.Caller(0)
	lines = aoc.Lines(f, filename)
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

var pts aoc.Set[aoc.Point]

func Run(lines []string) {
	g, bs := ParseGrid(lines)
	for _, v := range bs {
		v.MarkEmpty(&g)
	}
	total := 0
	for x := g.minx; x <= g.maxx; x++ {
		if c := g.GetMutRangeCell(x, theY); c != nil && *c == -1 {
			total++
		}
	}
	// PrintGrid(g.grid)
	fmt.Println(total)
	fmt.Println(len(pts))
	fmt.Println(total + len(pts))
}

type BS struct {
	b aoc.Point
	s aoc.Point
}

func MarkCell(g *RangeGrid, x, y int) {
	c := g.GetMutRangeCell(x, y)
	if c != nil && *c == 0 {
		*c = -1
	} else if c == nil && y == theY {
		pts.Add(aoc.NewPoint(x, y))
	}
}

func (bs BS) MarkEmpty(g *RangeGrid) {
	x, y := bs.s.X, bs.s.Y
	max := aoc.Abs(bs.b.X-x) + aoc.Abs(bs.b.Y-y)
	if !aoc.InRange(theY, y-max, y+max) {
		fmt.Println("skipped", bs)
		return
	}
	fmt.Println("---------")
	fmt.Println("bs:", bs)
	fmt.Println("max:", max)
	d := max - (aoc.Abs(y - theY))
	var a, b aoc.Point
	a = aoc.NewPoint(x-d, theY)
	b = aoc.NewPoint(x+d, theY)
	fmt.Println("a,b:", a, b)
	for _, c := range aoc.NewLine(a, b).Points() {
		MarkCell(g, c.X, c.Y)
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
	fmt.Println(arr)
	// fmt.Println(theY, theY)
	g := NewRangeGrid(maxx, minx, theY, theY)
	// g := NewRangeGrid(maxx, minx, maxy, miny)
	for _, bs := range arr {
		if b := g.GetMutRangeCell(bs.b.X, bs.b.Y); b != nil {
			*b = 1
		}
		if s := g.GetMutRangeCell(bs.s.X, bs.s.Y); s != nil {
			*s = 2
		}
	}

	return g, arr
}
