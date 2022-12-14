package main

import (
	"fmt"
	"log"
	"math"
	"runtime"

	"github.com/cbebe/aoc"
)

func main() {
	// f := "input.txt"
	f := "test.txt"
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	PartA(lines)
	// PartB(lines)
}

func PrintGrid(g Grid[int], start int) {
	for y, r := range g {
		for x, c := range r {
			if y == 0 && x == start {
				fmt.Print("+")
			} else if c == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func MoveToCell(g *aoc.Grid[int], current *aoc.Point, x, y int) bool {
	return true
}

func PartA(lines []string) {
	g, start := ParseGrid(lines)
	g[start][0] = 2
	for {
		p := aoc.NewPoint(start, 0)
		for {
			x, y := p.X, p.Y
			if MoveToCell(&g, &p, x, y+1) ||
				MoveToCell(&g, &p, x-1, y+1) || MoveToCell(&g, &p, x+1, y+1) {
				goto done
			}
		}
	done:
	}
}

// Returns the grid and the x-index of the starting point
func ParseGrid(lines []string) (aoc.Grid[int], int) {
	l := []aoc.Point{}
	minx, maxx, maxy := math.MaxInt, 0, 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		arr := aoc.TrimSplit(v, " -> ")
		for i := 0; i < len(arr)-1; i++ {
			nums := aoc.ParseSplit(arr[i], ",")
			a := aoc.NewPoint(nums[0], nums[1])
			nums = aoc.ParseSplit(arr[i+1], ",")
			b := aoc.NewPoint(nums[0], nums[1])
			lminx, lmaxx := aoc.MinMax(a.X, b.X)
			_, lmaxy := aoc.MinMax(a.Y, b.Y)
			minx, _ = aoc.MinMax(minx, lminx)
			_, maxx = aoc.MinMax(maxx, lmaxx)
			_, maxy = aoc.MinMax(maxy, lmaxy)
			l = append(l, aoc.NewLine(a, b).Points()...)
		}
	}
	s := aoc.Set[aoc.Point]{}
	s.AddSlice(l)
	g := make(aoc.Grid[int], 0, maxy)
	for i := 0; i < maxy; i++ {
		g = append(g, make([]int, maxx-minx+1))
	}
	for _, p := range l {
		x, y := p.X-minx, p.Y-1
		if c := g.GetMutCell(x, y); c != nil {
			if *c == 0 {
				*c = 1
			}
		} else {
			log.Fatalf("out of bounds")
		}
	}
	return g, 500 - minx
}
