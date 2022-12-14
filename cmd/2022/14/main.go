package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"

	"github.com/cbebe/aoc"
)

const partA bool = true

func main() {
	f := "input.txt"
	// f := "test.txt"
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	Run(lines)
}

func PrintGrid(g aoc.Grid[int], start int) {
	for y, r := range g {
		for x, c := range r {
			if y == 0 && x == start && c != 2 {
				fmt.Print("+")
			} else if c == 1 {
				fmt.Print("#")
			} else if c == 2 {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func CountRest(g aoc.Grid[int], start int) {
	total := 0
	for _, r := range g {
		for _, c := range r {
			if c == 2 {
				total++
			}

		}
	}
	PrintGrid(g, start)
	if partA {
		fmt.Println(total - 1)
	} else {
		fmt.Println(total)
	}
	os.Exit(0)
}

func MoveToCell(g *aoc.Grid[int], current *aoc.Point, off, start int) bool {
	x, y := current.X+off, current.Y+1
	c, err := g.SafeGetCell(x, y)
	if err != nil {
		if partA {
			CountRest(*g, start)
		} else {
			log.Fatalf("falling sand went out of bounds")
		}
	}
	if c == 0 {
		*g.GetMutCell(current.X, current.Y) = 0
		current.X = x
		current.Y = y
		*g.GetMutCell(x, y) = 2
		return false
	}
	return true
}

func Run(lines []string) {
	g, start := ParseGrid(lines)
	g[0][start] = 2
	for {
		p := aoc.NewPoint(start, 0)
		for {
			if MoveToCell(&g, &p, 0, start) &&
				MoveToCell(&g, &p, -1, start) &&
				MoveToCell(&g, &p, 1, start) {
				if !partA && p.X == start && p.Y == 0 {
					*g.GetMutCell(p.X, p.Y) = 2
					CountRest(g, start)
				}
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
	if partA {
		return MakePartAGrid(l, minx, maxx, maxy)
	} else {
		return MakePartBGrid(l, minx, maxx, maxy)
	}
}

func MakePartAGrid(l []aoc.Point, minx, maxx, maxy int) (aoc.Grid[int], int) {
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

func MakePartBGrid(l []aoc.Point, minx, maxx, maxy int) (aoc.Grid[int], int) {
	g := make(aoc.Grid[int], 0, maxy+4)
	offset := maxy + 3
	for i := 0; i < maxy+4; i++ {
		g = append(g, make([]int, (2*(offset+1))))
	}
	l = append(l, aoc.NewLine(aoc.NewPoint(500-offset, maxy+2), aoc.NewPoint(500+offset, maxy+2)).Points()...)
	for _, p := range l {
		x, y := p.X-(500-offset), p.Y
		if c := g.GetMutCell(x, y); c != nil {
			if *c == 0 {
				*c = 1
			}
		} else {
			log.Fatalf("out of bounds: %d, %d", x, y)
		}
	}
	return g, offset
}
