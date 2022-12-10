package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func split(v, sep string) (string, string) {
	arr := strings.Split(v, sep)
	return arr[0], arr[1]
}

func parseSplit(v string) (int, int) {
	a, b := split(v, "-")
	return aoc.ParseInt(a), aoc.ParseInt(b)
}

func main() {
	f := "input.txt"
	pairs := 0
	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		a, b := split(v, ",")
		a1, a2 := parseSplit(a)
		b1, b2 := parseSplit(b)
		// if PartA(a1, b1, a2, b2) {
		if PartB(a1, b1, a2, b2) {
			pairs++
		}
	}

	fmt.Println(pairs)
}

func PartA(a1, b1, a2, b2 int) bool {
	return (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2)
}

func PartB(a1, b1, a2, b2 int) bool {
	return r(a1, a2, b1) || r(a1, a2, b2) || r(b1, b2, a1) || r(b1, b2, a2)
}

func r(a, b, x int) bool {
	if a > b {
		log.Fatalf("%d > %d\n", a, b)
	}
	return aoc.InRange(x, a, b)
}
