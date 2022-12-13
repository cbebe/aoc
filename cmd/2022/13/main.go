package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sort"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"
	// f := "test.txt"
	_, filename, _, _ := runtime.Caller(0)
	file := aoc.ReadFile(f, filename)
	PartA(aoc.TrimSplit(file, "\n\n"))
	PartB(aoc.Lines(f, filename))
}

func PartB(lines []string) {
	b := append(lines, "[[2]]", "[[6]]")
	items := []Item{}
	for _, v := range b {
		if v == "" {
			continue
		}
		items = append(items, parseItem(v))
	}
	sort.Slice(items, func(i, j int) bool { return items[i].compare(items[j]) == Right })
	p := 1
	for i, v := range items {
		if isPacket(v) {
			p *= i + 1
		}
	}
	fmt.Println(p)
}

func isPacket(i Item) bool {
	if j, ok := i.(List); ok && len(j.items) == 1 {
		if k, ok := j.items[0].(List); ok && len(k.items) == 1 {
			n := k.items[0]
			if n == Number(2) || n == Number(6) {
				return true
			}
		}
	}
	return false
}

func PartA(lines []string) {
	total := 0
	for i, v := range lines {
		if v == "" {
			continue
		}
		arr := aoc.TrimSplit(v, "\n")
		left, right := parseItem(arr[0]), parseItem(arr[1])
		res := left.compare(right)
		if res == Next {
			log.Fatalf("something went wrong: result was Next")
		} else if res == Right {
			total += i + 1
		}
	}
	fmt.Println(total)
}

type Item interface{ compare(other Item) Result }
type List struct{ items aoc.Queue[Item] }
type Number int

type Result int

const (
	Right Result = iota
	Wrong
	Next
	Never
)

func (n Number) toList() List { return List{[]Item{n}} }

func (n Number) compare(other Item) Result {
	if olist, ok := other.(List); ok {
		return n.toList().compare(olist)
	} else if onum, ok := other.(Number); ok {
		a, b := n, onum
		if a > b {
			return Wrong
		} else if a < b {
			return Right
		} else {
			return Next
		}
	}
	switch v := other.(type) {
	default:
		fmt.Printf("unexpected type %T", v)
	}
	log.Fatalf("failed to cast %v to a list or number", other)
	return Never
}

func (l List) compare(other Item) Result {
	if olist, ok := other.(List); ok {
		for {
			a, aok := l.items.Pop()
			b, bok := olist.items.Pop()
			if !aok && !bok {
				return Next
			}
			if !aok {
				return Right
			}
			if !bok {
				return Wrong
			}
			if cmp := a.compare(b); cmp != Next {
				return cmp
			}
		}
	} else if onum, ok := other.(Number); ok {
		return l.compare(onum.toList())
	}
	log.Fatalf("failed to cast %v to a list or number", other)
	return Never
}

func collect(item any) Item {
	if n, ok := item.(float64); ok {
		return Number(int(n))
	}
	if list, ok := item.([]any); ok {
		items := []Item{}
		for _, w := range list {
			items = append(items, collect(w))
		}
		return List{items}
	}
	switch v := item.(type) {
	default:
		log.Fatalf("unexpected type %T", v)
	}
	return Number(-1)
}

func parseItem(item string) Item {
	var v any
	json.Unmarshal([]byte(item), &v)
	return collect(v)
}
