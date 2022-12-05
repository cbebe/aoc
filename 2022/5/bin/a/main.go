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

func split(v, sep string) (string, string) {
	arr := strings.Split(v, sep)
	return arr[0], arr[1]
}

func parseSplit(v string) (int, int) {
	a, b := split(v, "-")
	return parseInt(a), parseInt(b)
}

type Stack[T any] []T

func Push[T any](s *Stack[T], item T) {
	*s = append(*s, item)
}

func Pop[T any](s *Stack[T]) (T, bool) {
	var empty T
	if len(*s) == 0 {
		return empty, false
	} else {
		idx := len(*s) - 1
		item := (*s)[idx]
		*s = (*s)[:idx]
		return item, true
	}
}

func main() {
	input, err := script.File("input.txt").String()
	// input, err := script.File("test.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	numStacks := 0
	var stacks []Stack[byte]
	var rows Stack[[]byte]
	shouldStack := false
	doneStacks := false
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		if numStacks == 0 {
			numStacks = (len(v) + 1) / 4
			stacks = make([]Stack[byte], numStacks)
		} else if shouldStack {
			for i := len(rows) - 1; i >= 0; i-- {
				for j, b := range rows[i] {
					if b != 0 {
						Push(&stacks[j], b)
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
			Push(&rows, boxes)
		} else {
			commands := make([]string, 0, 6)
			arr := strings.Split(v, " ")
			for _, r := range arr {
				if r != "" {
					commands = append(commands, r)
				}
			}
			if len(commands) == 6 {
				count := parseInt(commands[1])
				from := parseInt(commands[3]) - 1
				to := parseInt(commands[5]) - 1
				for i := 0; i < count; i++ {
					item, exists := Pop(&stacks[from])
					if exists {
						Push(&stacks[to], item)
					}
				}
			}
		}
	}

	var message []byte
	for _, v := range stacks {
		item, exists := Pop(&v)
		if exists {
			message = append(message, item)
		}
	}

	fmt.Printf("%v\n", string(message))
}
