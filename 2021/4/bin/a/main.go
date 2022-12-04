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

func parseSplit(v, sep string) []int {
	arr := strings.Split(v, sep)
	ints := make([]int, 0)
	for _, a := range arr {
		if a == "" {
			continue
		}
		ints = append(ints, parseInt(a))
	}
	return ints
}

type Board [][]int

func (b Board) match() {
	fmt.Printf("%v\n", b)
}

func main() {
	// input, err := script.File("input.txt").String()
	input, err := script.File("test.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	var calls []int
	lines := strings.Split(input, "\n")
	first := true
	boards := make([]Board, 0)
	n := 0
	var b Board
	for _, v := range lines {
		if v == "" {
			continue
		}
		if first {
			calls = parseSplit(v, ",")
			first = false
			continue
		} else {
			b = append(b, parseSplit(v, " "))
			n++
			if n%5 == 0 {
				boards = append(boards, b)
				b = make([][]int, 0, 5)
				n = 0
			}
		}
	}
	for _, v := range boards {
		v.match()
	}
	fmt.Println(len(boards))
	fmt.Println(calls)
}
