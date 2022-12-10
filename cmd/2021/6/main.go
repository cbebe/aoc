package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	// f := "input.txt"
	f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	fish := aoc.ParseSplit(strings.TrimSpace(aoc.ReadFile(f, filename)), ",")
	// lmfao
	for days := 0; days < 80; days++ {
		newFish := 0
		for i := 0; i < len(fish); i++ {
			fish[i]--
			if fish[i] == -1 {
				fish[i] = 6
				newFish++
			}
		}
		fmt.Println(len(fish), newFish)
		for i := 0; i < newFish; i++ {
			fish = append(fish, 8)
		}
		fmt.Println(len(fish))
		fmt.Println("day:", days)
		fmt.Println(fish)
	}
	fmt.Println(len(fish))
}
