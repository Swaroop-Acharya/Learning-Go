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

func filter( n []int , fn func(n int) bool )[]int {
	var res []int;
	for _,v:= range n {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res;
}


func filterFunctions(){
	n:= []int {1,2,3,4,5,6,7,8,9,10,11,12}
	odd:= filter(n, func(n int) bool { return n % 2 == 0} )
	even:= filter(n, func(n int) bool { return n % 2 == 1} )
	fmt.Println(odd)
	fmt.Println(even)
}

func main() {
	functionValues()
	functionClosures()
	filterFunctions()
}
