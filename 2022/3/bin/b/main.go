package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bitfield/script"
)

func getval(b rune) int {
	var val int
	if int(b) >= int('a') && int(b) <= int('z') {
		val = int(b) - int('a') + 1
	} else {
		val = int(b) - int('A') + 27
	}

	return val
}

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	line := 0
	var mymap map[rune]int
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		if line%3 == 0 {
			mymap = make(map[rune]int)
		}
		occ := make(map[rune]bool)
		for _, b := range v {
			if _, exists := occ[b]; !exists {
				occ[b] = true
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
