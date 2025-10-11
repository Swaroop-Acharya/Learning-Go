package main

import (
	"fmt"
)

func compute(fn func(int, int) int) int {
	return fn(100, 200)
}

func getCount() func() int {
	count := 0
	return func() int {
		count++;
		return count;
	}
}

func functionValues() {
	foo := func() {
		fmt.Println("Hi")
	}
	foo()

	food := func(name string) string {
		return "I like " + name
	}

	fmt.Println(food("Pizza"))

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(10, 20))
	fmt.Println(compute(add))
}

func functionClosures() {
	counter := getCount();
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func main() {
	functionValues()
	functionClosures()
}
