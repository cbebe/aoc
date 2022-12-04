package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func parseInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func split(v, sep string) (string, string) {
	arr := strings.Split(v, sep)
	return arr[0], arr[1]
}

func parseSplit(v string) (int, int) {
	a, b := split(v, "-")
	return parseInt(a), parseInt(b)
}

func r(a, b, x int) bool {
	if a > b {
		log.Fatalf("%d > %d\n", a, b)
	}
	return a <= x && b >= x
}

func main() {
	input, err := script.File("input.txt").String()
	// input, err := script.File("test.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	pairs := 0
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		a, b := split(v, ",")
		a1, a2 := parseSplit(a)
		b1, b2 := parseSplit(b)
		if r(a1, a2, b1) || r(a1, a2, b2) || r(b1, b2, a1) || r(b1, b2, a2) {
			pairs++
		}
	}

	fmt.Println(pairs)
}
