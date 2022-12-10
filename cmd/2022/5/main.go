package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

func main() {
	f := "input.txt"

	numStacks := 0
	var stacks []aoc.Stack[byte]
	var rows aoc.Stack[[]byte]
	shouldStack := false
	doneStacks := false
	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		if numStacks == 0 {
			numStacks = (len(v) + 1) / 4
			stacks = make([]aoc.Stack[byte], numStacks)
		} else if shouldStack {
			for i := len(rows) - 1; i >= 0; i-- {
				for j, b := range rows[i] {
					if b != 0 {
						stacks[j].Push(b)
					}
				}
			}
			shouldStack = false
			doneStacks = true
		}

		if !doneStacks {
			var res string
			var boxes = make([]byte, numStacks)
			for i, r := range v + "0" {
				if i > 0 && (i+1)%4 == 0 {
					s := strings.TrimSpace(res)
					if s != "" {
						if len(s) == 1 {
							// time to stack
							shouldStack = true
						} else {
							boxes[((i+1)/4)-1] = s[1]
						}
					}
					res = ""
				}
				res = res + string(r)
			}
			rows.Push(boxes)
		} else {
			commands := make([]string, 0, 6)
			arr := strings.Split(v, " ")
			for _, r := range arr {
				if r != "" {
					commands = append(commands, r)
				}
			}
			if len(commands) == 6 {
				count := aoc.ParseInt(commands[1])
				from := aoc.ParseInt(commands[3]) - 1
				to := aoc.ParseInt(commands[5]) - 1
				// PartA(count, from, to, stacks)
				PartB(count, from, to, stacks)
			}
		}
	}

	var message []byte
	for _, v := range stacks {
		item, exists := v.Pop()
		if exists {
			message = append(message, item)
		}
	}

	fmt.Printf("%v\n", string(message))
}

func PartA(count, from, to int, stacks []aoc.Stack[byte]) {
	for i := 0; i < count; i++ {
		item, exists := stacks[from].Pop()
		if exists {
			stacks[to].Push(item)
		}
	}
}

func PartB(count, from, to int, stacks []aoc.Stack[byte]) {
	stack := make(aoc.Stack[byte], 0)
	for i := 0; i < count; i++ {
		item, exists := stacks[from].Pop()
		if exists {
			stack.Push(item)
		}
	}
	for i := 0; i < count; i++ {
		item, exists := stack.Pop()
		if exists {
			stacks[to].Push(item)
		}
	}
}
