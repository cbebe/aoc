package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"
	// f := "test.txt"

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
			num += int(A)
		case 'b':
			num += int(B)
		case 'c':
			num += int(C)
		case 'd':
			num += int(D)
		case 'e':
			num += int(E)
		case 'f':
			num += int(F)
		case 'g':
			num += int(G)
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
	ints  map[int]int
	segs  map[int]Seg
	hints []string
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
	oop := func(i Segment, r rune) rune {
		if lit.Has(i) {
			return r
		} else {
			return '.'
		}
	}
	for k := range s.Set {
		lit.Add(segmap[k])
	}
	a, b, c, d := oop(A, 'a'), oop(B, 'b'), oop(C, 'c'), oop(D, 'd')
	e, f, g := oop(E, 'e'), oop(F, 'f'), oop(G, 'g')
	var sb strings.Builder
	horizontal(&sb, a)
	vertical(&sb, b, c)
	horizontal(&sb, d)
	vertical(&sb, e, f)
	horizontal(&sb, g)
	return sb.String()
}

func (s *Segs) Add(num int, segs string) {
	seg := toSeg(segs)
	s.ints[seg.AsInt()] = num
	s.segs[num] = seg
}

func (s Segs) Get(num int) Seg {
	return s.segs[num]
}

func (s Seg) Match(other Seg) bool {
	in := s.Intersection(other.Set)
	return len(in) == len(s.Set)
}

func createNums(hints []string) map[int]int {
	fives := []string{}
	sixes := []string{}
	s := Segs{map[int]int{}, map[int]Seg{}, hints}
	for _, v := range hints {
		switch len(v) {
		case 2:
			s.Add(1, v)
		case 3:
			s.Add(7, v)
		case 4:
			s.Add(4, v)
		case 5:
			fives = append(fives, v)
		case 6:
			sixes = append(sixes, v)
		case 7:
			s.Add(8, v)
		default:
			log.Fatalf("invalid len: %s", v)
		}
	}
	one := s.Get(1).Set
	four := s.Get(4).Set
	eight := s.Get(8).Set
	a := s.Get(7).Difference(one)
	segmap[a.First()] = A
	s.Add(6, Find(sixes, func(o string) bool { return !toSeg(o).Subset(one) }))
	c := one.Difference(s.Get(6).Set)
	segmap[c.First()] = C
	f := one.Difference(c)
	segmap[f.First()] = F
	fivesChars := toSeg(fives[0]).Set
	for _, v := range fives[1:] {
		fivesChars = fivesChars.Intersection(toSeg(v).Set)
	}
	g := fivesChars.Difference(four).Difference(a)
	segmap[g.First()] = G
	nine := g.Union(a).Union(four)
	s.Add(9, MatchSeg(sixes, nine))
	e := eight.Difference(nine)
	segmap[e.First()] = E
	s.Add(2, Find(fives, func(o string) bool { return toSeg(o).Subset(e) }))
	two := s.Get(2).Set
	d := two.Difference(a).Difference(c).Difference(e).Difference(g)
	segmap[d.First()] = D
	three := one.Union(a).Union(d).Union(g)
	s.Add(3, MatchSeg(fives, three))
	b := four.Difference(three)
	segmap[b.First()] = B
	five := a.Union(b).Union(d).Union(f).Union(g)
	s.Add(5, MatchSeg(fives, five))
	zero := a.Union(b).Union(c).Union(e).Union(f).Union(g)
	s.Add(0, MatchSeg(sixes, zero))

	return s.ints
}

func MatchSeg(hints []string, s aoc.Set[rune]) string {
	return Find(hints, func(o string) bool { return Seg{s}.Match(toSeg(o)) })
}

func Find[T any](arr []T, where func(T) bool) T {
	for _, v := range arr {
		if where(v) {
			return v
		}
	}
	log.Fatalf("could not find")
	var t T
	return t
}

func PartB(lines []string) {
	segmap = map[rune]Segment{}
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
			n = (n * 10) + numMap[toSeg(num).AsInt()]
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
		for _, num := range strings.Split(strings.Split(v, " | ")[1], " ") {
			n := len(num)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				total++
			}
		}
	}
	fmt.Println(total)
}
