package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

type Address struct {
	city, street string
}

type Employee struct {
	name    string
	address Address
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{}
	v3 = Vertex{x: 10}
	p  = &Vertex{20, 20}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{10, 20}
	fmt.Println(v.x)
	fmt.Println(v.y)

	v.x = v.x + 20
	fmt.Println(v.x)

	p := &v

	fmt.Println(p)

	(*p).x = 100
	fmt.Println(v)
	// or
	p.x = 200
	fmt.Println(v)

	// Structs inside struct

	e1 := Employee{"Swaroop", Address{"Mangalore", "Belthangdy"}}
	e2 := Employee{name: "Sam", address: Address{city: "Bangalore", street: "Whitefield"}}

	fmt.Println(e1, e2)
}
