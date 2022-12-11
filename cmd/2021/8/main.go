package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"sort"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	// f := "input.txt"
	f := "test.txt"

	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	PartB(lines)
}

func toSeg(str string) Seg {
	set := aoc.Set[rune]{}
	for _, v := range str {
		set.Add(v)
	}
	return Seg{set}
}

type Seg struct {
	aoc.Set[rune]
}

func (s Seg) AsInt() int {
	num := 0
	for k := range s.Set {
		switch k {
		case 'a':
			num += (1 << 0)
		case 'b':
			num += (1 << 1)
		case 'c':
			num += (1 << 2)
		case 'd':
			num += (1 << 3)
		case 'e':
			num += (1 << 4)
		case 'f':
			num += (1 << 5)
		case 'g':
			num += (1 << 6)
		}
	}

	return num
}

type Segment int

const (
	A Segment = 1 << 0
	B         = 1 << 1
	C         = 1 << 2
	D         = 1 << 3
	E         = 1 << 4
	F         = 1 << 5
	G         = 1 << 6
)

var segmap map[rune]Segment

type Segs struct {
	ints map[int]int
	segs map[int]Seg
}

func horizontal(w io.Writer, s rune) {
	fmt.Fprintf(w, " %s \n", strings.Repeat(string(s), 4))
}
func vertical(w io.Writer, a, b rune) {
	for i := 0; i < 2; i++ {
		fmt.Fprintf(w, "%s    %s\n", string(a), string(b))
	}
}

func (s Seg) String() string {
	lit := aoc.Set[Segment]{}
	var sb strings.Builder
	for k := range s.Set {
		lit.Add(segmap[k])
	}
	a, b, c, d, e, f, g := '.', '.', '.', '.', '.', '.', '.'
	if lit.Has(A) {
		a = 'a'
	}
	if lit.Has(B) {
		b = 'b'
	}
	if lit.Has(C) {
		c = 'c'
	}
	if lit.Has(D) {
		d = 'd'
	}
	if lit.Has(E) {
		e = 'e'
	}
	if lit.Has(F) {
		f = 'f'
	}
	if lit.Has(G) {
		g = 'g'
	}
	horizontal(&sb, a)
	vertical(&sb, b, c)
	horizontal(&sb, d)
	vertical(&sb, e, f)
	horizontal(&sb, g)
	return sb.String()
}

func (s *Segs) Add(num int, segs string) {
	s.ints[toSeg(segs).AsInt()] = num
	s.segs[num] = toSeg(segs)
}

func (s Segs) Get(num int) Seg {
	return s.segs[num]
}

func (s Seg) Diff(other Seg, by int) bool {
	return len(s.Difference(other.Set)) == by
}

func createNums(hints []string) map[int]int {
	segmap = make(map[rune]Segment)
	sort.Slice(hints, func(i, j int) bool {
		return len(hints[i]) < len(hints[j])
	})

	fmt.Println(hints)
	s := Segs{map[int]int{}, map[int]Seg{}}
	//  ....
	// .    c
	// .    c
	//  ....
	// .    f
	// .    f
	//  ....
	s.Add(1, hints[0])
	s.Add(7, hints[1])
	s.Add(4, hints[2])
	s.Add(8, hints[9])

	segmap[s.Get(7).Difference(s.Get(1).Set).First()] = A

	//  aaaa
	// b    c
	// b    c
	//  dddd
	// .    f
	// .    f
	//  ....
	abcdf := s.Get(1).Union(s.Get(7).Set).Union(s.Get(4).Set)

	s.Add(9, findNine(hints, Seg{abcdf}))
	segmap[s.Get(9).Difference(abcdf).First()] = G
	s.Add(2, findTwo(hints, s.Get(9)))
	s.Add(3, findThree(hints, s.Get(2)))

	//  ....
	// .    .
	// .    .
	//  ....
	// e    .
	// e    .
	//  ....
	e := s.Get(2).Difference(s.Get(3).Set)
	segmap[e.First()] = E
	//  ....
	// .    .
	// .    .
	//  ....
	// .    f
	// .    f
	//  ....
	f := s.Get(3).Difference(s.Get(2).Set)
	segmap[f.First()] = F
	fmt.Println(s.Get(1).Set)
	fmt.Println(e, f)
	fmt.Println(1)
	fmt.Println(s.Get(1))
	fmt.Println(2)
	fmt.Println(s.Get(2))
	fmt.Println(3)
	fmt.Println(s.Get(3))
	fmt.Println(4)
	fmt.Println(s.Get(4))
	fmt.Println(7)
	fmt.Println(s.Get(7))
	fmt.Println(8)
	fmt.Println(s.Get(8))
	fmt.Println(9)
	fmt.Println(s.Get(9))
	s.Add(5, findFive(hints, s.Get(2)))
	s.Add(6, findSix(hints, Seg{e.Union(s.Get(5).Set)}))

	return s.ints
}

func findNum(hints []string, pred func(v string) bool, start int, name string) string {
	for _, v := range hints[start : start+3] {
		if pred(v) {
			return v
		}
	}
	log.Fatalf("can't find %s\n", name)
	return ""
}

func findTwo(hints []string, nine Seg) string {
	return findNum(hints, func(v string) bool { return !nine.Subset(toSeg(v).Set) }, 3, "two")
}

func findThree(hints []string, two Seg) string {
	return findNum(hints, func(v string) bool { return two.Diff(toSeg(v), 1) }, 3, "three")
}

func findFive(hints []string, two Seg) string {
	return findNum(hints, func(v string) bool { return two.Diff(toSeg(v), 2) }, 3, "five")
}

func findSix(hints []string, six Seg) string {
	return findNum(hints, func(v string) bool { return six.Diff(toSeg(v), 0) }, 6, "six")
}

func findNine(hints []string, abcdf Seg) string {
	return findNum(hints, func(v string) bool { return abcdf.Diff(toSeg(v), 1) }, 6, "nine")
}

func PartB(lines []string) {
	total := 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " | ")
		hints := strings.Split(arr[0], " ")
		numMap := createNums(hints)
		nums := strings.Split(arr[1], " ")
		n := 0
		for _, num := range nums {
			n *= 10
			n += numMap[toSeg(num).AsInt()]
		}
		total += n
	}
	fmt.Println(total)
}

func PartA(lines []string) {
	total := 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		nums := strings.Split(strings.Split(v, " | ")[1], " ")
		for _, num := range nums {
			n := len(num)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				total++
			}

		}
	}
	fmt.Println(total)
}
