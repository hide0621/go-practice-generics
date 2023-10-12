package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {

	s := make(Set[T])
	for _, v := range xs {
		s[v] = struct{}{}
	}
	return s
}

func main() {

	a := NewSet(1, 2, 3)
	fmt.Println(a)

}
