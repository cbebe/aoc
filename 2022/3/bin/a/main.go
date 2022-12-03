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
		mymap := make(map[rune]bool)
		half := len(v) / 2
		for _, b := range v[:half] {
			mymap[b] = true
		}
		for _, b := range v[half:] {
			if _, exists := mymap[b]; exists {
				var val int
				if int(b) >= int('a') && int(b) <= int('z') {
					val = int(b) - int('a') + 1
				} else {
					val = int(b) - int('A') + 27
				}
				score += val
				goto end
			}
		}
	end:
	}

	fmt.Println(score)
}
