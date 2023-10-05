package main

import (
	"fmt"
	"strconv"
)

func ToStringSlice[T fmt.Stringer](inputSlice []T) []string {

	result := []string{}

	for _, element := range inputSlice {
		result = append(result, element.String())
	}

	return result
}

type MyInt int

func (i MyInt) String() string {

	return strconv.Itoa(int(i))

}

func main() {

	fmt.Println(ToStringSlice([]MyInt{1, 2, 3, 4, 5}))

}
