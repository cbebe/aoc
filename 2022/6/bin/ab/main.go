package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bitfield/script"
)

type Queue[T any] []T

func Push[T any](s *Queue[T], item T) {
	*s = append(*s, item)
}

func Pop[T any](s *Queue[T]) (T, bool) {
	var empty T
	if len(*s) == 0 {
		return empty, false
	} else {
		item := (*s)[0]
		*s = (*s)[1:]
		return item, true
	}
}

func ProcessLine(v string, n int) int {
	s := make(map[rune]bool, 0)
	m := make(Queue[rune], 0)
	for i, c := range v {
		Push(&m, c)
		if _, exists := s[c]; exists {
			for len(m) > 1 && m[0] != c {
				if popped, hasItem := Pop(&m); hasItem {
					delete(s, popped)
				}
			}
			Pop(&m)
		} else {
			s[c] = true
		}
		// fmt.Printf("%v\n", m)
		if len(m) >= n {
			return i + 1
		}
	}

	return len(v)
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		p := v
		if f == "input.txt" {
			p = v[:20] + "..."
		}
		fmt.Printf("4: %d -- %s\n", ProcessLine(v, 4), p)
		fmt.Printf("14: %d -- %s\n", ProcessLine(v, 14), p)
	}
}
