package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}
	set := make(map[int]int)
	maxes := make(map[int]int)
	curridx := 0
	set[curridx] = 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			curridx += 1
			set[curridx] = 0
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
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

// 197384 wrong answer
