package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"
	// f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	crabs := map[int]int{}
	for _, c := range aoc.ParseSplit(strings.TrimSpace(aoc.ReadFile(f, filename)), ",") {
		crabs[c]++
	}

	fuel := map[int]int{}
	for i := 0; i < len(crabs); i++ {
		for k, v := range crabs {
			// fuel[k] += PartA(k, l, w)
			fuel[i] += PartB(i, k, v)
		}
	}
	_, min := aoc.MinMap(fuel, math.MaxInt)
	fmt.Println(min)
}

func PartA(i, k, v int) int {
	return aoc.Abs(i-k) * v
}

// Sum of natural numbers
// https://www.cuemath.com/numbers/natural-numbers-from-1-to-100/
func Nat(n int) int {
	return n * (2 + (n - 1)) / 2
}

func PartB(i, k, v int) int {
	return Nat(aoc.Abs(i-k)) * v
}
