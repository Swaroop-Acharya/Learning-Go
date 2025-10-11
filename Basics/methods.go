package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type MyInt int;

func (p Person) greet() {
	fmt.Println(p.name)
}

func methodsOnStruct() {
	p:= Person{"Swaroop",23};
	p.greet();

}


func (n MyInt) double() int {
	return int(n*2);
}

func methodsOnBasicType(){
	var i MyInt = 10;
	fmt.Println(i.double())

	i2:= MyInt(30); // convert this into custom type so that double methods can be called on that
	fmt.Printf("%T\n",i2); // type is main.MyInt 
	fmt.Println(i2.double())
}


// Functions with a receiver == Method
func main() {
	methodsOnStruct();
	methodsOnBasicType();
}
