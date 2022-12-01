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

	totals := map[int]int{}
	for i, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		totals[i] += num
		totals[i - 1] += num
		totals[i - 2] += num
	}
	times := 0
	for t := 1; t < len(totals) - 2; t++ {
		if totals[t] > totals[t - 1] {
			times++
		}
	}
	fmt.Println(times)
}
