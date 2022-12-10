package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type Step struct {
	steps int
	dir   aoc.Direction
}

func Parse(v string) aoc.Direction {
	switch v {
	case "U":
		return aoc.Up
	case "D":
		return aoc.Down
	case "L":
		return aoc.Left
	case "R":
		return aoc.Right
	default:
		log.Fatalln("unrecognized direction")
	}
	return -1
}

type Steps []Step

func (s Step) String() string {
	return fmt.Sprintf("(%s %d)", s.dir.String(), s.steps)
}

type Rope []Cell

func (r Rope) Tail() *Cell {
	return &r[len(r)-1]
}

func (r Rope) Head() *Cell {
	return &r[0]
}

type Cell struct {
	x int
	y int
}

func (c Cell) String() string {
	return fmt.Sprintf("(%d %d)", c.x, c.y)
}

func (c *Cell) Move(dir aoc.Direction) {
	switch dir {
	case aoc.Up:
		c.y++
	case aoc.Down:
		c.y--
	case aoc.Left:
		c.x--
	case aoc.Right:
		c.x++
	}
}

func (c *Cell) Follow(other Cell) {
	if aoc.Abs(other.x-c.x) == 2 && aoc.Abs(other.y-c.y) == 2 {
		c.x += (other.x - c.x) / 2
		c.y += (other.y - c.y) / 2
	} else if other.x-c.x == -2 {
		c.y = other.y
		c.x--
	} else if other.x-c.x == 2 {
		c.y = other.y
		c.x++
	} else if other.y-c.y == -2 {
		c.x = other.x
		c.y--
	} else if other.y-c.y == 2 {
		c.x = other.x
		c.y++
	}
}

type Visited struct {
	aoc.Set[int]
}

func (v Visited) String() string {
	var sb strings.Builder
	for k := range v.Set {
		x := k >> 16
		y := 0xffff & k
		fmt.Fprintf(&sb, "%s, ", Cell{x, y})
	}
	return strings.Trim(sb.String(), ", ")
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	// f := "test-2.txt"

	// rope := make(Rope, 2)
	rope := make(Rope, 10)
	visited := Visited{aoc.Set[int]{}}
	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		arr := strings.Split(v, " ")
		dir := arr[0]
		numsteps := aoc.ParseInt(arr[1])
		step := Step{numsteps, Parse(dir)}
		fmt.Printf("%s %d\n", step.dir, step.steps)
		for i := 0; i < step.steps; i++ {
			head := rope.Head()
			fmt.Printf("h %s -> ", head)
			rope.Head().Move(step.dir)
			fmt.Printf("%s", head)
			tail := rope.Tail()
			x := tail.x
			y := tail.y
			for i := 1; i < len(rope); i++ {
				rope[i].Follow(rope[i-1])
			}
			if tail.x != x || tail.y != y {
				fmt.Printf(" | t %s -> %s\n", Cell{x, y}, tail)
			} else {
				fmt.Println()
			}
			visited.Set.Add((tail.x << 16) + tail.y)
		}
	}
	// 1996 too low
	fmt.Println(len(visited.Set))
}
