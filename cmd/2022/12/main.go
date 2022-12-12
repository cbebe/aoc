package main

import (
	"fmt"
	"log"
	"math"
	"runtime"

	"github.com/cbebe/aoc"
)

type Grid [][]Point

type Point struct {
	Value rune
	X     int
	Y     int
}

var visited aoc.Set[Point]
var distance map[Point]int
var grid Grid

var pq aoc.PriorityQueue[Point]

var start Point
var end Point

// const partA = true

const partA = false

func main() {
	f := "input.txt"
	// f := "test.txt"
	visited = make(aoc.Set[Point])
	distance = make(map[Point]int)
	pq = make(aoc.PriorityQueue[Point], 0)
	grid = make(Grid, 0)

	_, filename, _, _ := runtime.Caller(0)
	for i, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		row := make([]Point, 0, len(v))
		for j, w := range v {
			point := Point{Value: w, X: j, Y: i}
			distance[point] = math.MaxInt
			if partA {
				if w == 'S' {
					start = point
					distance[point] = 0
				} else if w == 'E' {
					end = point
				}
			} else {
				if w == 'E' {
					start = point
					distance[point] = 0
				}
			}
			row = append(row, point)
		}
		grid = append(grid, row)
	}
	// fmt.Println(grid)
	dijkstra()
}

func dijkstra() {
	pq.Push(&aoc.PQItem[Point]{Value: start, Priority: 0})
	for len(pq) != 0 {
		node := pq.Min()
		point, dist := node.Value, node.Priority
		fmt.Println(point)
		if (partA && point == end) || (!partA && point.Value == 'a') {
			fmt.Println(dist)
			return
		}
		if visited.Has(point) {
			continue
		}
		visited.Add(point)
		for _, p := range point.getChildren() {
			d, _ := aoc.MinMax(dist+1, distance[p])
			distance[p] = d
			pq.Push(&aoc.PQItem[Point]{Value: p, Priority: d})
		}
	}
	log.Fatalf("goal not found")
}

func (g Grid) getCell(x, y int) *Point {
	if len(g) > 0 && y >= 0 && y < len(g) && x >= 0 && x < len(g[0]) {
		return &g[y][x]
	}
	return nil
}

func val(r rune) int {
	if r == 'S' {
		if partA {
			return 96
		} else {
			return val('a')
		}
	} else if r == 'E' {
		return 123
	}
	return int(r)
}

func (p Point) isNext(other Point) bool {
	if partA {
		return val(other.Value) <= val(p.Value) || val(other.Value) == val(p.Value)+1
	} else {
		return val(other.Value) >= val(p.Value) || val(other.Value) == val(p.Value)-1
	}
}

func (p Point) getChildren() []Point {
	children := []Point{}
	for _, other := range []*Point{grid.getCell(p.X-1, p.Y), grid.getCell(p.X+1, p.Y), grid.getCell(p.X, p.Y-1), grid.getCell(p.X, p.Y+1)} {
		if other != nil && p.isNext(*other) {
			children = append(children, *other)
		}
	}
	return children
}

func (p Point) String() string {
	return fmt.Sprintf("%s (%d-%d)", string(p.Value), p.X, p.Y)
}
