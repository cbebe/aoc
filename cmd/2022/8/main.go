package main

import (
	"fmt"
	"runtime"

	"github.com/cbebe/aoc"
)

func MakeCell(pos, i int, dir aoc.Direction) string {
	if dir == aoc.Up || dir == aoc.Down {
		return fmt.Sprintf("%d-%d", pos, i)
	} else {
		return fmt.Sprintf("%d-%d", i, pos)
	}
}

func (g Grid) PrintStep(pos, max int, dir aoc.Direction) {
	fmt.Printf("pos: %d, direction: %s -- ", pos, dir)
	fmt.Print("looking at: ")
	switch dir {
	case aoc.Up:
		for i := max - 1; i >= 0; i-- {
			fmt.Printf("%d ", g[i][pos])
		}
	case aoc.Down:
		for i := 0; i < max; i++ {
			fmt.Printf("%d ", g[i][pos])
		}
	case aoc.Left:
		for i := max - 1; i >= 0; i-- {
			fmt.Printf("%d ", g[pos][i])
		}
	case aoc.Right:
		for i := 0; i < max; i++ {
			fmt.Printf("%d ", g[pos][i])
		}
	}
	fmt.Println()
}

type Grid [][]int

func (g Grid) C(pos, i int, dir aoc.Direction) int {
	if dir == aoc.Up || dir == aoc.Down {
		return g[i][pos]
	} else {
		return g[pos][i]
	}
}

func (g Grid) Peek(pos int, dir aoc.Direction, visible aoc.Set[string]) {
	max := len(g[0])
	if dir == aoc.Up || dir == aoc.Down {
		max = len(g)
	}
	g.PrintStep(pos, max, dir)
	if dir == aoc.Left || dir == aoc.Up {
		visible.Add(MakeCell(pos, max-1, dir))
		h := g.C(pos, max-1, dir)
		for i := max - 2; i >= 0; i-- {
			current := g.C(pos, i, dir)
			if current > g.C(pos, i+1, dir) && current > h {
				visible.Add(MakeCell(pos, i, dir))
				h = current
			}
		}
	} else {
		visible.Add(MakeCell(pos, 0, dir))
		h := g.C(pos, 0, dir)
		for i := 1; i < max; i++ {
			current := g.C(pos, i, dir)
			if current > g.C(pos, i-1, dir) && current > h {
				visible.Add(MakeCell(pos, i, dir))
				h = current
			}
		}
	}
}

func MakeGrid(lines []string) Grid {
	grid := Grid{}
	for _, v := range lines {
		if v == "" {
			continue
		}
		row := []int{}
		for _, c := range v {
			height := aoc.ParseInt(string(c))
			row = append(row, height)
		}
		grid = append(grid, row)
	}

	return grid
}

func PartA(grid Grid) {
	visible := aoc.Set[string]{}

	for i := 0; i < len(grid); i++ {
		grid.Peek(i, aoc.Right, visible)
		grid.Peek(i, aoc.Left, visible)
	}
	for j := 0; j < len(grid[0]); j++ {
		grid.Peek(j, aoc.Down, visible)
		grid.Peek(j, aoc.Up, visible)
	}

	rows := [][]string{}
	for i, v := range grid {
		row := []string{}
		for j, c := range v {
			if visible.Has(fmt.Sprintf("%d-%d", j, i)) {
				row = append(row, fmt.Sprint(c))
			} else {
				row = append(row, "X")
			}
		}
		rows = append(rows, row)
		fmt.Println(v)
	}
	fmt.Println()
	for _, r := range rows {
		fmt.Println(r)
	}
	fmt.Println(len(visible))
}

func PartB(grid Grid) {
	scene := map[string]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			key := fmt.Sprintf("%d-%d", j, i)
			scene[key] = grid.Scene(i, j)
		}
	}

	fmt.Println(aoc.MaxMap(scene))
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	grid := MakeGrid(aoc.Lines(f, filename))
	// PartA(grid)
	PartB(grid)
}

func (g Grid) Scene(i, j int) int {
	return g.LookStraight(i, j, aoc.Up) * g.LookStraight(i, j, aoc.Down) *
		g.LookStraight(i, j, aoc.Left) * g.LookStraight(i, j, aoc.Right)
}

func (g Grid) LookStraight(x, y int, dir aoc.Direction) int {
	visible := 0
	switch dir {
	case aoc.Up:
		for i := y - 1; i >= 0; i-- {
			visible++
			if g[x][y] <= g[x][i] {
				break
			}
		}
	case aoc.Down:
		for i := y + 1; i < len(g); i++ {
			visible++
			if g[x][y] <= g[x][i] {
				break
			}
		}
	case aoc.Left:
		for i := x + 1; i < len(g[0]); i++ {
			visible++
			if g[x][y] <= g[i][y] {
				break
			}
		}
	case aoc.Right:
		for i := x - 1; i >= 0; i-- {
			visible++
			if g[x][y] <= g[i][y] {
				break
			}
		}
	}
	return visible
}
