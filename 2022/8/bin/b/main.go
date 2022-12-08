package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Grid [][]int

func (g Grid) Scene(i, j int) int {
	return g.Peek(i, j, Up) * g.Peek(i, j, Down) *
		g.Peek(i, j, Left) * g.Peek(i, j, Right)
}

func (g Grid) Peek(x, y int, dir Direction) int {
	visible := 0
	switch dir {
	case Up:
		for i := y - 1; i >= 0; i-- {
			visible++
			if g[x][y] <= g[x][i] {
				break
			}
		}
	case Down:
		for i := y + 1; i < len(g); i++ {
			visible++
			if g[x][y] <= g[x][i] {
				break
			}
		}
	case Left:
		for i := x + 1; i < len(g[0]); i++ {
			visible++
			if g[x][y] <= g[i][y] {
				break
			}
		}
	case Right:
		for i := x - 1; i >= 0; i-- {
			visible++
			if g[x][y] <= g[i][y] {
				break
			}
		}
	}
	return visible
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, _ := script.File(f).String()
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

	scene := map[string]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			key := fmt.Sprintf("%d-%d", j, i)
			scene[key] = grid.Scene(i, j)
		}
	}

	max := 0
	for _, v := range scene {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
