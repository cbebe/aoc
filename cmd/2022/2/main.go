package main

import (
	"fmt"
	"runtime"

	"github.com/cbebe/aoc"
)

func GetScoreB(v string) int {
	// 1 - R
	// 2 - P
	// 3 - S
	// 0 - L
	// 3 - D
	// 6 - W
	// X - L
	// Y - D
	// Z - W
	switch v {
	case "A X": // R S
		return 3 + 0
	case "B X": // P R
		return 1 + 0
	case "C X": // S P
		return 2 + 0
	case "A Y": // R R
		return 1 + 3
	case "B Y": // P P
		return 2 + 3
	case "C Y": // S S
		return 3 + 3
	case "A Z": // R P
		return 2 + 6
	case "B Z": // P S
		return 3 + 6
	case "C Z": // S R
		return 1 + 6
	}
	return 0
}

func GetScoreA(v string) int {
	switch v {
	// 1 - R
	// 2 - P
	// 3 - S
	// 0 - L
	// 3 - D
	// 6 - W
	case "A X": // R R - draw R
		return 1 + 3
	case "B X": // P R
		return 1 + 0
	case "C X": // S R
		return 1 + 6
	case "A Y": // R P
		return 2 + 6
	case "B Y": // P P
		return 2 + 3
	case "C Y": // S P
		return 2 + 0
	case "A Z": // R S
		return 3 + 0
	case "B Z": // P S
		return 3 + 6
	case "C Z": // S S
		return 3 + 3
	}
	return 0
}

func main() {
	f := "input.txt"

	score := 0

	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		// score += GetScoreA(v)
		score += GetScoreB(v)
	}

	fmt.Println(score)
}
