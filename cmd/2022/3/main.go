package main

import (
	"fmt"
	"runtime"

	"github.com/cbebe/aoc"
)

func getval(b rune) int {
	var val int
	if aoc.InRange(int(b), int('a'), int('z')) {
		val = int(b) - int('a') + 1
	} else {
		val = int(b) - int('A') + 27
	}

	return val
}

func main() {
	f := "input.txt"
	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	// PartA(lines)
	PartB(lines)
}

func PartA(lines []string) {
	score := 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		mymap := aoc.Set[rune]{}
		half := len(v) / 2
		for _, b := range v[:half] {
			mymap.Add(b)
		}
		for _, b := range v[half:] {
			if mymap.Has(b) {
				score += getval(b)
				goto end
			}
		}
	end:
	}

	fmt.Println(score)
}

func PartB(lines []string) {
	score := 0
	line := 0
	var mymap map[rune]int
	for _, v := range lines {
		if v == "" {
			continue
		}
		if line%3 == 0 {
			mymap = make(map[rune]int)
		}
		occ := aoc.Set[rune]{}
		for _, b := range v {
			if !occ.Has(b) {
				occ.Add(b)
				mymap[b] += 1
				if mymap[b] == 3 {
					score += getval(b)
					line = 0
					goto end
				}
			}
		}
		line++
	end:
	}

	fmt.Println(score)
}
