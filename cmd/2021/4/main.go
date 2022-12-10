package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type Board struct {
	grid    [][]int
	inboard aoc.Set[int]
	marked  aoc.Set[int]
}

func NewBoard() Board {
	return Board{make([][]int, 0, 5), aoc.Set[int]{}, aoc.Set[int]{}}
}

func ParseBoards(file string) ([]int, []Board) {
	arr := strings.Split(file, "\n\n")
	calls := aoc.ParseSplit(arr[0], ",")
	boards := make([]Board, 0)
	for _, v := range arr[1:] {
		rows := strings.Split(v, "\n")
		b := NewBoard()
		for _, row := range rows {
			r := aoc.ParseSplit(row, " ")
			b.grid = append(b.grid, r)
			for _, i := range r {
				b.inboard.Add(i)
			}
		}
		boards = append(boards, b)
	}
	return calls, boards
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	file := aoc.ReadFile(f, filename)
	calls, boards := ParseBoards(file)

	// PartA(calls, boards)
	PartB(calls, boards)
}

func PartB(calls []int, boards []Board) {
	winSet := aoc.Set[int]{}
	winArr := []int{}
	for _, c := range calls {
		for i, b := range boards {
			val := b.match(c)
			if val != 0 && !winSet.Has(i) {
				winSet.Add(i)
				winArr = append(winArr, val)
			}
		}
	}
	fmt.Println(winArr[len(winArr)-1])
}

func PartA(calls []int, boards []Board) {
	for _, c := range calls {
		for _, b := range boards {
			val := b.match(c)
			if val != 0 {
				fmt.Println(val)
				return
			}
		}
	}
}

func (b *Board) horizontal(x int) bool {
	for i := 0; i < 5; i++ {
		if !b.marked.Has(b.grid[x][i]) {
			return false
		}
	}
	return true
}

func (b *Board) vertical(x int) bool {
	for i := 0; i < 5; i++ {
		if !b.marked.Has(b.grid[i][x]) {
			return false
		}
	}
	return true
}

func (b *Board) diagonal(dir aoc.Direction) bool {
	for x := 0; x < 5; x++ {
		y := x
		if dir == aoc.Down {
			y = 4 - x
		}
		if !b.marked.Has(b.grid[x][y]) {
			return false
		}
	}
	return true
}

func (b *Board) value() int {
	total := 0
	for _, i := range b.grid {
		for _, j := range i {
			if !b.marked.Has(j) {
				total += j
			}
		}
	}
	return total
}

func (b *Board) match(c int) int {
	if b.inboard.Has(c) {
		b.marked.Add(c)
	}
	if len(b.marked) >= 5 {
		if b.diagonal(aoc.Up) || b.diagonal(aoc.Down) {

		}
		for i := 0; i < 5; i++ {
			if b.horizontal(i) || b.vertical(i) {
				return c * b.value()
			}
		}

	}
	return 0
}
