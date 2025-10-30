package main

import (
	"fmt"
)

func Filter[T comparable](s []T, elem T) int {
	for i, v := range s {
		if v == elem {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Filter(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Filter(ss, "bar"))
}
