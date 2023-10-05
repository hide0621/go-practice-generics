package main

import "fmt"

// func PrintSliceInts(i []int) {

// 	for _, v := range i {
// 		fmt.Println(v)
// 	}
// }

// func PrintSliceStrings(s []string) {

// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

func PrintSlice[T any](s []T) {

	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {

	// PrintSliceInts([]int{1, 2, 3, 4, 5})

	// PrintSliceStrings([]string{"a", "b", "c", "d", "e"})

	PrintSlice([]int{1, 2, 3, 4, 5})

	PrintSlice([]string{"a", "b", "c", "d", "e"})
}
