package aoc

import (
	// "fmt"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func TrimSplit(f, sep string) []string {
	trimmed := []string{}
	for _, v := range strings.Split(f, sep) {
		trimmed = append(trimmed, strings.TrimSpace(v))
	}
	return trimmed
}

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

// Oof at Golang generics
func MinMap[T constraints.Ordered, K comparable](m map[K]T, max T) (K, T) {
	var min = max
	var key K
	for k, v := range m {
		if v < min {
			min = v
			key = k
		}
	}
	return key, min
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinMax[T constraints.Ordered](a, b T) (T, T) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func MinMaxInt(a, b int, s ...int) (int, int) {
	if len(s) == 0 {
		return MinMax(a, b)
	}

	arr := append(s, a, b)
	min, max := math.MaxInt, 0
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
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

type Point struct {
	X int
	Y int
}

type Line struct {
	A Point
	B Point
}

func NewLine(a, b Point) Line {
	return Line{a, b}
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func X(l Line) int { return Abs(l.A.X-l.B.X) / (l.B.X - l.A.X) }
func Y(l Line) int { return Abs(l.B.Y-l.A.Y) / (l.B.Y - l.A.Y) }

func (l Line) Points() []Point {
	s := []Point{}
	if l.A.X == l.B.X { // Vertical
		min, max := MinMax(l.A.Y, l.B.Y)
		for i := min; i <= max; i++ {
			s = append(s, Point{l.A.X, i})
		}
	} else if l.A.Y == l.B.Y { // Horizontal
		min, max := MinMax(l.A.X, l.B.X)
		for i := min; i <= max; i++ {
			s = append(s, Point{i, l.A.Y})
		}
	} else { // Diagonal
		for x, y := l.A.X, l.A.Y; x != l.B.X || y != l.B.Y; x, y = x+X(l), y+Y(l) {
			s = append(s, Point{x, y})
		}
		s = append(s, Point{l.B.X, l.B.Y})
	}
	return s
}
