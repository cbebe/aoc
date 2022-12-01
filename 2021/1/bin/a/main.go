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

	times := 0
	max := 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if max > 0 && num > max {
			times++
		}
		max = num
	}

	fmt.Println(times)
}
