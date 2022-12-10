package main

import (
	"fmt"
	"runtime"
	"sort"

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
	elves := []int{}
	elves = append(elves, 0)
	maxcal := 0
	maxidx := 0
	curridx := 0
	for _, v := range lines {
		if v == "" {
			elves = append(elves, 0)
			curridx += 1
		} else {
			num := aoc.ParseInt(v)
			elves[curridx] += num
			if curridx == maxidx {
				maxcal = elves[curridx]
			} else if elves[curridx] > maxcal {
				maxidx = curridx
				maxcal = elves[curridx]
			}
		}
	}

	fmt.Println(maxcal)
}

func PartB(lines []string) {
	set := make(map[int]int)
	maxes := make(map[int]int)
	curridx := 0
	set[curridx] = 0
	for _, v := range lines {
		if v == "" {
			curridx += 1
			set[curridx] = 0
		} else {
			num := aoc.ParseInt(v)
			set[curridx] += num
			if _, ismax := maxes[curridx]; ismax || len(maxes) < 3 {
				maxes[curridx] = set[curridx]
			} else {
				for k, v := range maxes {
					if set[curridx] > v {
						delete(maxes, k)
						maxes[curridx] = set[curridx]
					}
				}
			}
		}
	}

	// total := 0
	// fmt.Println(len(maxes))
	// for _, v := range maxes {
	// 	total += v
	// }
	// fmt.Println(total)

	// fmt.Printf("%v\n", set)

	// RIP
	s := []int{}
	for _, v := range set {
		s = append(s, v)
	}
	sort.Ints(s)
	fmt.Printf("%d\n", s[len(s)-1]+s[len(s)-2]+s[len(s)-3])
}
