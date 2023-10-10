package main

import "fmt"

type Pair[T any] struct {
	First  T
	Second T
}

func (p Pair[T]) PrintField() {
	fmt.Println(p.First)
	fmt.Println(p.Second)
}

func main() {

	intResult := Pair[int]{10, 20}
	intResult.PrintField()

	intString := Pair[string]{"Hello", "World"}
	intString.PrintField()
}
