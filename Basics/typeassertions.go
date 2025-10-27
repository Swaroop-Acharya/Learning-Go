package main

import (
	"fmt"
)

func printAny(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Println(s)
	} else if n, ok := i.(int); ok {
		fmt.Println(n)
	} else {
		fmt.Println("Unknown type")
	}
}

func assertionExample() {
	//Json parsing or fmt logic
	printAny("Hello")
	printAny(12)
	printAny(12.121)
}

func main() {
	var i interface{} = "hello"

	//not safe assertion
	//t := i.(string)
	// fmt.Println(t)

	//safe assertions
	t, ok := i.(int)

	fmt.Println(t, ok)

	assertionExample()
}
