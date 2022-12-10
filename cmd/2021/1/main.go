package main

import (
	"fmt"
	"runtime"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	// PartA(lines)
	PartB(lines)
}

func PartA(lines []string) {
	times := 0
	max := 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		num := aoc.ParseInt(v)
		if max > 0 && num > max {
			times++
		}
		max = num
	}

	fmt.Println(times)
}

func PartB(lines []string) {
	totals := map[int]int{}
	for i, v := range lines {
		if v == "" {
			continue
		}
		num := aoc.ParseInt(v)
		totals[i] += num
		totals[i-1] += num
		totals[i-2] += num
	}
	times := 0
	for t := 1; t < len(totals)-2; t++ {
		if totals[t] > totals[t-1] {
			times++
		}
	}
	fmt.Println(times)
}
