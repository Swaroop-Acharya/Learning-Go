package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	a float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	h, w float64
}

func (s Square) Area() float64 {
	return s.a * s.a
}

func (c *Circle) Area() float64 {
	if c == nil {
		return 0.0
	}
	return math.Pi * c.r * c.r
}

func (r Rectangle) Area() float64 {
	return r.h * r.w
}

func printArea(s Shape) {
	fmt.Println("s:", s)
	fmt.Println(s.Area())
}

func interfaceBasics() {
	s := Square{a: 10}
	printArea(s)

	r := Rectangle{h: 5, w: 10}
	printArea(r)
}

func describe(s Shape) {
	fmt.Printf("(%v, %T)\n", s, s)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var s Shape
	s = Square{10}
	fmt.Println(s)
	describe(s)

	s = Rectangle{10, 15}
	fmt.Println(s)
	describe(s)
}

func interfaceNilValueHandler() {
	var s Shape
	var c *Circle
	s = c
	describe(s)
	fmt.Println(s.Area())
}

func nilInterfaceValues() {
	var s Shape

	describe(s)
	fmt.Println(s.Area()) // runtime -error nill interface
}

func emptyInterface() {
	var i interface{} // empty interface / generic interface
	describeEmptyInterface(i)
	i = 31
	describeEmptyInterface(i)
	i = "hello"
	describeEmptyInterface(i)
}

func main() {
	interfaceBasics()
	interfaceValues()
	interfaceNilValueHandler()
	nilInterfaceValues()
	emptyInterface()
}
