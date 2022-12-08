package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

var visible map[string]bool

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

func MakeCell(pos, i int, dir Direction) string {
	if dir == Up || dir == Down {
		return fmt.Sprintf("%d-%d", pos, i)
	} else {
		return fmt.Sprintf("%d-%d", i, pos)
	}
}

func (g Grid) PrintStep(pos, max int, dir Direction) {
	fmt.Printf("pos: %d, direction: %s -- ", pos, dir)
	fmt.Print("looking at: ")
	switch dir {
	case Up:
		for i := max - 1; i >= 0; i-- {
			fmt.Printf("%d ", g[i][pos])
		}
	case Down:
		for i := 0; i < max; i++ {
			fmt.Printf("%d ", g[i][pos])
		}
	case Left:
		for i := max - 1; i >= 0; i-- {
			fmt.Printf("%d ", g[pos][i])
		}
	case Right:
		for i := 0; i < max; i++ {
			fmt.Printf("%d ", g[pos][i])
		}
	}
	fmt.Println()
}

type Grid [][]int

func (g Grid) C(pos, i int, dir Direction) int {
	if dir == Up || dir == Down {
		return g[i][pos]
	} else {
		return g[pos][i]
	}
}

func (g Grid) Peek(pos int, dir Direction) {
	max := len(g[0])
	if dir == Up || dir == Down {
		max = len(g)
	}
	g.PrintStep(pos, max, dir)
	if dir == Left || dir == Up {
		visible[MakeCell(pos, max-1, dir)] = true
		h := g.C(pos, max-1, dir)
		for i := max - 2; i >= 0; i-- {
			current := g.C(pos, i, dir)
			if current > g.C(pos, i+1, dir) && current > h {
				visible[MakeCell(pos, i, dir)] = true
				h = current
			}
		}
	} else {
		visible[MakeCell(pos, 0, dir)] = true
		h := g.C(pos, 0, dir)
		for i := 1; i < max; i++ {
			current := g.C(pos, i, dir)
			if current > g.C(pos, i-1, dir) && current > h {
				visible[MakeCell(pos, i, dir)] = true
				h = current
			}
		}
	}
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, err := script.File(f).String()
	if err != nil {
		log.Fatal(err)
	}

	grid := Grid{}
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		row := []int{}
		for _, c := range v {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalln(err)
			}
			row = append(row, height)
		}
		grid = append(grid, row)
	}

	visible = map[string]bool{}

	for i := 0; i < len(grid); i++ {
		grid.Peek(i, Right)
		grid.Peek(i, Left)
	}
	for j := 0; j < len(grid[0]); j++ {
		grid.Peek(j, Down)
		grid.Peek(j, Up)
	}

	rows := [][]string{}
	for i, v := range grid {
		row := []string{}
		for j, c := range v {
			if yes, ok := visible[fmt.Sprintf("%d-%d", j, i)]; ok && yes {
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
