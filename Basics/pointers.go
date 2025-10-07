package main

import (
	"fmt"
)


func addOne(x *int) {
	*x = *x + 1;
}

func main() {
	x := 10
	p := &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	fmt.Println(x)

	addOne(&x);

	fmt.Println(x)

}
