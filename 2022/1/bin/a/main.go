package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	elves := []int{}
	elves = append(elves, 0)
	maxcal := 0
	maxidx := 0
	curridx := 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			elves = append(elves, 0)
			curridx += 1
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
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
