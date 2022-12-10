package aoc

import (
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Lines(f, filename string) []string {
	return strings.Split(ReadFile(f, filename), "\n")
}

func GetDir(filename string) string {
	dir := path.Dir(filename)
	return strings.Split(dir, path.Dir(path.Dir(dir)))[1]
}

func ReadFile(f, filename string) string {
	b, err := os.ReadFile("input" + GetDir(filename) + "/" + f)
	if err != nil {
		log.Fatalf("readFile: %v", err)
	}
	return string(b)
}

func ParseInt(v string) int {
	num, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("parseInt: %v", err)
	}
	return num
}

func ParseSplit(v, sep string) []int {
	arr := strings.Split(v, sep)
	ints := make([]int, 0)
	for _, a := range arr {
		if a == "" {
			continue
		}
		ints = append(ints, ParseInt(a))
	}
	return ints
}

func InRange[T constraints.Ordered](x, a, b T) bool {
	return x >= a && x <= b
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

// Oof at Golang generics
func MaxMap[T constraints.Ordered, K comparable](m map[K]T) T {
	var max T
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func MinMax[T constraints.Ordered](a, b T) (T, T) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

// Oof at Golang generics
func MaxSlice[T constraints.Ordered](s []T) T {
	var max T
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
