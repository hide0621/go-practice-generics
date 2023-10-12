package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {

	s := make(Set[T])
	for _, v := range xs {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func main() {

	a := NewSet(1, 2, 3)
	fmt.Println(a)

	a.Remove(1)
	fmt.Println(a)
}
