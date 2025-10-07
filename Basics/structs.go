package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v:= Vertex{10,20};
	fmt.Println(v.x)
	fmt.Println(v.y)

	v.x=v.x+20;
	fmt.Println(v.x)

}
