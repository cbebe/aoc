package main

import (
	"fmt"
	"runtime"
	"sort"
	"strings"

	"github.com/cbebe/aoc"
)

type Monkey struct {
	items     *aoc.Queue[int]
	op        func(int, *Monkey) int
	t         int
	f         int
	inspected int
	div       int
	partA     bool
}

func (m *Monkey) test(i int) bool {
	return i%m.div == 0
}

func ParseMonkeys(monke []string, partA bool) []*Monkey {
	monkeys := []*Monkey{}
	for _, v := range monke {
		if v == "" {
			continue
		}
		arr := aoc.TrimSplit(v, "\n")
		items := aoc.ParseSplit(strings.Split(arr[1], "Starting items: ")[1], ", ")
		op := parseOp(strings.Split(arr[2], "Operation: ")[1])
		div := aoc.ParseInt(strings.Split(arr[3], "Test: divisible by ")[1])
		to := aoc.ParseInt(strings.Split(arr[4], "If true: throw to monkey ")[1])
		from := aoc.ParseInt(strings.Split(arr[5], "If false: throw to monkey ")[1])

		q := aoc.Queue[int](items)
		monkeys = append(monkeys, &Monkey{&q, op, to, from, 0, div, partA})
	}
	return monkeys
}

func (m *Monkey) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "true: %d, false: %d\n", m.t, m.f)
	fmt.Fprintf(&sb, "items: %v\n", m.items)
	fmt.Fprintf(&sb, "inspected: %v\n", m.inspected)
	return sb.String()
}

func PrintMonkeys(monkeys []*Monkey, i int, partA bool) {
	fmt.Printf("After round %d:\n", i+1)
	if partA {
		for j, m := range monkeys {
			fmt.Println(j, m.items)
		}
	} else {
		for j, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %d times.\n", j, m.inspected)
		}
	}
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	partA := true
	// partA := false
	rounds := 20
	if !partA {
		rounds = 10000
	}

	_, filename, _, _ := runtime.Caller(0)
	file := aoc.ReadFile(f, filename)
	monkeys := ParseMonkeys(strings.Split(file, "\n\n"), partA)
	PrintMonkeys(monkeys, -1, partA)
	toPrint := aoc.Set[int]{}
	toPrint.AddSlice([]int{1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000})
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			if partA {
				fmt.Printf("starting items for monkey %d: %v\n", j, monkeys[j].items)
			}
			for len(*monkeys[j].items) > 0 {
				if item, ok := monkeys[j].items.Pop(); ok {
					if partA {
						fmt.Printf("Monkey %d:\n", j)
						fmt.Printf("  Monkey inspects an item with a worry level of %d:\n", item)
					}
					monkeys[j].inspected++
					worry := monkeys[j].op(item, monkeys[j])
					if partA {
						fmt.Printf("  Worry level is transformed to %d.\n", worry)
						worry /= 3
						fmt.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %d.\n", worry)
					}
					to := monkeys[j].f
					if monkeys[j].test(worry) {
						to = monkeys[j].t
					}
					if partA {
						fmt.Printf("  Item with worry level %d is thrown to Monkey %d.\n\n", worry, to)
					}
					monkeys[to].items.Push(worry)
				}
			}
		}
		if partA || toPrint.Has(i+1) {
			PrintMonkeys(monkeys, i, partA)
		}
	}
	scores := []int{}
	for i, m := range monkeys {
		scores = append(scores, m.inspected)
		fmt.Println(i, m)
	}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[len(scores)-1] * scores[len(scores)-2])
}

func parseOp(exp string) func(int, *Monkey) int {
	toks := strings.Split(exp, " ")
	a, op, b := toks[2], toks[3], toks[4]
	fmt.Printf("operation: %s, a: %s, b: %s, op: %s\n", exp, a, b, op)
	return func(i int, m *Monkey) int {
		x := i
		if a != "old" {
			x = aoc.ParseInt(a)
		}
		y := i
		if b != "old" {
			y = aoc.ParseInt(b)
		}
		switch op {
		case "+":
			if m.partA {
				return x + y
			} else {
				return x%m.div + y%m.div
			}
		case "*":
			if m.partA {
				return x * y
			} else {
				return x%m.div + y%m.div
			}
		}
		return 0
	}
}
