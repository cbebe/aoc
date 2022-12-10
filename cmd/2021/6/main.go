package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"
	// f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	fish := map[int]int{}
	fish[6] = 0
	for _, v := range aoc.ParseSplit(strings.TrimSpace(aoc.ReadFile(f, filename)), ",") {
		fish[v]++
	}
	for days := 0; days < 256; days++ {
		newFish := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += newFish
		fish[8] = newFish
		fmt.Println(fish)
	}
	total := 0
	for _, v := range fish {
		total += v
	}
	fmt.Println(total)
}
