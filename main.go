package main

import (
	"fmt"
)

type Vector[T any] []T

type IntVector = Vector[int]

func main() {

	var v Vector[int] = []int{1, 2, 3}
	fmt.Println(v)

	var v2 Vector[string] = []string{"a", "b", "c"}
	fmt.Println(v2)

	v3 := IntVector{10, 20, 30}
	fmt.Println(v3)

}
