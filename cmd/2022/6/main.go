package main

import (
	"fmt"
	"runtime"

	"github.com/cbebe/aoc"
)

func ProcessLine(v string, n int) int {
	s := aoc.Set[rune]{}
	m := make(aoc.Queue[rune], 0)
	for i, c := range v {
		m.Push(c)
		if s.Has(c) {
			for len(m) > 1 && m[0] != c {
				if popped, hasItem := m.Pop(); hasItem {
					delete(s, popped)
				}
			}
			m.Pop()
		} else {
			s.Add(c)
		}
		// fmt.Printf("%v\n", m)
		if len(m) >= n {
			return i + 1
		}
	}

	return len(v)
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		p := v
		if f == "input.txt" {
			p = v[:20] + "..."
		}
		fmt.Printf("4: %d -- %s\n", ProcessLine(v, 4), p)
		fmt.Printf("14: %d -- %s\n", ProcessLine(v, 14), p)
	}
}
