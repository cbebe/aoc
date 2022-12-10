package aoc

type Stack[T any] []T

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

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

func (s *Queue[T]) Push(item T) {
	*s = append(*s, item)
}

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

func (s *Set[T]) Add(item T) {
	(*s)[item] = member
}

func (s *Set[T]) Remove(item T) {
	delete(*s, item)
}

func (s *Set[T]) Has(item T) bool {
	_, ok := (*s)[item]
	return ok
}
