package main

import (
	"fmt"
)

type Position struct {
	x int
	y int
}

func mapBasics() {
	var m map[string]int = make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)
	fmt.Println(m["one"])

	m1 := make(map[string]Position)

	fmt.Println(m1)

	m1["first"] = Position{x: 10, y: 20}
	fmt.Println(m1)
}

func mapLiterals() {
	m := map[string]Position{
		"first":  Position{x: 10, y: 20},
		"second": Position{x: 15, y: 25},
	}
	//insert
	m["third"] = Position{ x: 30, y:20}
	fmt.Println(m)

	//retrieve
	elem:= m["first"]
	fmt.Println(elem)

	//delete
	delete(m,"second")
	fmt.Println(m)

	//update
	m["first"]=Position{ x: 300, y:200} 
	fmt.Println(m)
}

func main() {
	mapBasics()
	mapLiterals()
}
