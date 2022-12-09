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

func (c *Cell) Follow(other Cell) {
	if abs(other.x-c.x) == 2 && abs(other.y-c.y) == 2 {
		c.x += (other.x - c.x) / 2
		c.y += (other.y - c.y) / 2
	}
	if other.x-c.x == -2 {
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
	// f := "test-2.txt"
	input, err := script.File(f).String()
	if err != nil {
		log.Fatal(err)
	}

	rope := make([]Cell, 10)
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
			fmt.Printf("h %s -> ", rope[0])
			rope[0].Move(step.dir)
			fmt.Printf("%s", rope[0])
			x := rope[9].x
			y := rope[9].y
			for i := 1; i < len(rope); i++ {
				rope[i].Follow(rope[i-1])
			}
			if rope[9].x != x || rope[9].y != y {
				fmt.Printf(" | t %s -> %s\n", Cell{x, y}, rope[9])
			} else {
				fmt.Println()
			}
			visited[(rope[9].x<<16)+rope[9].y] = true
		}
	}
	// 1996 too low
	fmt.Println(len(visited))
	// fmt.Println(visited)
}
