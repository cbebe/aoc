package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type Sub struct {
	up       func(*Sub, int)
	down     func(*Sub, int)
	forward  func(*Sub, int)
	distance int
	depth    int
	aim      int
}

func NewSub(up, down, forward func(*Sub, int)) *Sub {
	return &Sub{up, down, forward, 0, 0, 0}
}

func UpA(s *Sub, d int)      { s.depth -= d }
func DownA(s *Sub, d int)    { s.depth += d }
func ForwardA(s *Sub, d int) { s.distance += d }

func UpB(s *Sub, d int)   { s.aim -= d }
func DownB(s *Sub, d int) { s.aim += d }
func ForwardB(s *Sub, d int) {
	s.distance += d
	s.depth += d * s.aim
}

func main() {
	f := "input.txt"
	// f := "test.txt"

	// sub := NewSub(UpA, DownA, ForwardA)
	sub := NewSub(UpB, DownB, ForwardB)

	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " ")
		num := aoc.ParseInt(arr[1])
		if arr[0][0] == 'f' {
			sub.forward(sub, num)
		} else if arr[0][0] == 'd' {
			sub.down(sub, num)
		} else if arr[0][0] == 'u' {
			sub.up(sub, num)
		}
	}

	fmt.Println(sub.distance * sub.depth)
}
