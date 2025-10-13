package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type MyInt int

func (p Person) greet() {
	fmt.Println(p.name)
}

func (p Person) changeName(name string) {
	p.name = name
}

func (p *Person) changeNameV2(name string) {
	p.name = name
}

func methodsOnStruct() {
	p := Person{"Swaroop", 23}
	p.greet()
}

func (n MyInt) double() int {
	return int(n * 2)
}

func methodsOnBasicType() {
	var i MyInt = 10
	fmt.Println(i.double())

	i2 := MyInt(30)        // convert this into custom type so that double methods can be called on that
	fmt.Printf("%T\n", i2) // type is main.MyInt
	fmt.Println(i2.double())
}

func pointerRecievers() {
	p := Person{"Kam", 22}
	p.changeName("Sam")
	fmt.Println(p) // o/p Kam

	(&p).changeNameV2("Sam")
	fmt.Println(p) // o/p Sam

	// or

	p.changeNameV2("Tam")
	fmt.Println(p) // o/p Tam
}

func methodsAndPointerIndirection() {
	p:= Person{"Prime",21};
	p.changeNameV2("Valter")
	fmt.Println(p) // Valter
}

// Functions with a receiver == Method
func main() {
	methodsOnStruct()
	methodsOnBasicType()
	pointerRecievers()
	methodsAndPointerIndirection()
}
