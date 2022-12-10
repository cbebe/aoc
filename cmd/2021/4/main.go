package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func parseSplit(v, sep string) []int {
	arr := strings.Split(v, sep)
	ints := make([]int, 0)
	for _, a := range arr {
		if a == "" {
			continue
		}
		ints = append(ints, aoc.ParseInt(a))
	}
	return ints
}

type Board [][]int

func (b Board) match() {
	fmt.Printf("%v\n", b)
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	var calls []int
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	first := true
	boards := make([]Board, 0)
	n := 0
	var b Board
	for _, v := range lines {
		if v == "" {
			continue
		}
		if first {
			calls = parseSplit(v, ",")
			first = false
			continue
		} else {
			b = append(b, parseSplit(v, " "))
			n++
			if n%5 == 0 {
				boards = append(boards, b)
				b = make([][]int, 0, 5)
				n = 0
			}
		}
	}
	for _, v := range boards {
		v.match()
	}
	fmt.Println(len(boards))
	fmt.Println(calls)
}
