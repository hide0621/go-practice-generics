package main

import (
	"fmt"
)

type Number interface {
	~int | int32 | int64 | float32 | float64 | string
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type MyInt int

func main() {

	fmt.Println(Max(1, 4))
	fmt.Println(Max(3.3, 5.9))

	fmt.Println(Max[MyInt](1, 9))

}
