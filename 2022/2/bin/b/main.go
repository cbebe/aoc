package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
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
			score += 3 + 0
		case "B X": // P R
			score += 1 + 0
		case "C X": // S P
			score += 2 + 0
		case "A Y": // R R
			score += 1 + 3
		case "B Y": // P P
			score += 2 + 3
		case "C Y": // S S
			score += 3 + 3
		case "A Z": // R P
			score += 2 + 6
		case "B Z": // P S
			score += 3 + 6
		case "C Z": // S R
			score += 1 + 6
		}
	}

	fmt.Println(score)
}
