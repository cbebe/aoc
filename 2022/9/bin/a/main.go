package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

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

type Step struct {
	steps int
	dir   Direction
}

func Parse(v string) Direction {
	switch v {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		log.Fatalln("unrecognized direction")
	}
	return -1
}

type Steps []Step

func (s Step) String() string {
	return fmt.Sprintf("(%s %d)", s.dir.String(), s.steps)
}

type Cell struct {
	x int
	y int
}

func (c Cell) String() string {
	return fmt.Sprintf("(%d %d)", c.x, c.y)
}

func (c *Cell) Move(dir Direction) {
	switch dir {
	case Up:
		c.y++
	case Down:
		c.y--
	case Left:
		c.x--
	case Right:
		c.x++
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Visited map[int]bool

func (v Visited) String() string {
	var sb strings.Builder
	for k := range v {
		x := k >> 16
		y := 0xffff & k
		fmt.Fprintf(&sb, "%s, ", Cell{x, y})
	}
	return strings.Trim(sb.String(), ", ")
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, err := script.File(f).String()
	if err != nil {
		log.Fatal(err)
	}

	head := Cell{0, 0}
	tail := Cell{0, 0}
	visited := Visited{}
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " ")
		dir := arr[0]
		numsteps, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatalln(err)
		}
		step := Step{numsteps, Parse(dir)}
		fmt.Printf("%s %d\n", step.dir, step.steps)
		for i := 0; i < step.steps; i++ {
			fmt.Printf("h %s -> ", head)
			head.Move(step.dir)
			fmt.Printf("%s", head)
			x := tail.x
			y := tail.y
			if head.x-tail.x == -2 {
				tail.y = head.y
				tail.x--
			} else if head.x-tail.x == 2 {
				tail.y = head.y
				tail.x++
			} else if head.y-tail.y == -2 {
				tail.x = head.x
				tail.y--
			} else if head.y-tail.y == 2 {
				tail.x = head.x
				tail.y++
			}
			if tail.x != x || tail.y != y {
				fmt.Printf(" | t %s -> %s\n", Cell{x, y}, tail)
			} else {
				fmt.Println()
			}
			visited[(tail.x<<16)+tail.y] = true
		}
	}
	fmt.Println(len(visited))
	// fmt.Println(visited)
}
