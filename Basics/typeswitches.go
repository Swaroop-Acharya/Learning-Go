package main

import (
	"fmt"
)

func do(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("This is %T\n",t)
	case string:
		fmt.Printf("This is %T\n",t)
	case float64:
		fmt.Printf("This is %T\n",t)
	default:
		fmt.Printf("This is unknown type: %T\n",t)
	}
}

func main() {
	do(10)
	do("hi")
	do(12.3)
	do(true)
}
