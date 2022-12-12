package aoc

import (
	"container/heap"
	"fmt"
)

type Stack[T any] []T

func (s *Stack[T]) Push(item T) { *s = append(*s, item) }

func (s *Stack[T]) Pop() (T, bool) {
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

type Queue[T any] []T

func (s *Queue[T]) Push(item T) { *s = append(*s, item) }

func (s *Queue[T]) Pop() (T, bool) {
	var empty T
	if len(*s) == 0 {
		return empty, false
	} else {
		item := (*s)[0]
		*s = (*s)[1:]
		return item, true
	}
}

type void struct{}

var member void

type Set[T comparable] map[T]void

func (s *Set[T]) AddSlice(items []T) {
	for _, i := range items {
		s.Add(i)
	}
}

func (s *Set[T]) Add(item T) { (*s)[item] = member }

func (s *Set[T]) Remove(item T) { delete(*s, item) }

func (s *Set[T]) Has(item T) bool {
	_, ok := (*s)[item]
	return ok
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersect := Set[T]{}
	for k := range s {
		if other.Has(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := Set[T]{}
	for k := range s {
		union.Add(k)
	}
	for k := range other {
		union.Add(k)
	}
	return union
}

func (s Set[T]) Subset(other Set[T]) bool {
	for k := range other {
		if !s.Has(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := Set[T]{}
	for k := range s {
		if !other.Has(k) {
			diff.Add(k)
		}
	}
	return diff
}

func (s Set[T]) First() T {
	var val T
	for v := range s {
		return v
	}

	return val
}

type PQItem[T any] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T any] []*PQItem[T]

func (pq PriorityQueue[T]) Len() int           { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool { return (*pq[i]).Priority < (*pq[j]).Priority }
func (pq *PriorityQueue[T]) Min() *PQItem[T]   { return heap.Remove(pq, 0).(*PQItem[T]) }
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n, item := len(*pq), x.(*PQItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Grid[T any] [][]T

func (g Grid[T]) GetCell(x, y int) T {
	if len(g) > 0 && y >= 0 && y < len(g) && x >= 0 && x < len(g[0]) {
		return g[y][x]
	}
	var empty T
	return empty
}

// Like GetCell but will crash on unsafe access.
// If err is not nil, the returned value is invalid
func (g Grid[T]) SafeGetCell(x, y int) (T, error) {
	var empty T
	if len(g) == 0 {
		return empty, fmt.Errorf("grid has no rows")
	} else if len(g[0]) == 0 {
		return empty, fmt.Errorf("grid has no columns")
	} else if y < 0 || y >= len(g) || x < 0 || x >= len(g[0]) {
		return empty, fmt.Errorf("index out of bounds for grid with %d rows and %d cols: (x: %d, y: %d)", len(g), len(g[0]), x, y)
	}

	return g[y][x], nil
}
