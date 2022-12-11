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
