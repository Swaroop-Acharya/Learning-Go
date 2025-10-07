package main

import (
	"fmt"
	"runtime"
	"time"
)

// declaring variable
var python, java bool

var start, end int = 1, 3

var (
	num    int     = 1
	pi     float32 = 3.1
	piFull float64 = 3.14159
	name   string  = "swaroop"
	flag   bool    = true
	ch     rune    = '😊'
)

func add(x int, y int) int {
	return x + y
}

// return multiple return results
func swap(x, y string) (string, string) {
	return y, x
}

// naked return
func addNsub(x, y int) (add, sub int) {
	add = x + y
	sub = x - y
	return
}

func variables() {
	k := 3
	fmt.Println(k)
	fmt.Println(pi, piFull)
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)
}

func flowControl() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	//	for ; ; {
	//	 	fmt.Println("Hi")
	//	}

	// while loop
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)

	// forever
	// for{}

	// if statment
	if sum == 100 {
		fmt.Println(sum)
	}

	// if with a short statement
	if n := add(1, 2); n == 3 {
		fmt.Println("Corrrect", n)
	} else {
		fmt.Println("incorrect", n)
	}

	// switch statements
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")

	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s/n", os)
	}

	fmt.Print("When is sunday?")
	today := time.Now().Weekday()
	fmt.Printf("type is : %T/n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("In one day")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func deferStmt() {
	fmt.Println("Count start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Count end")
}

func main() {
	fmt.Println(add(1, 2))
	a, b := swap("hellow", "world")
	fmt.Println(a, b)
	fmt.Println(addNsub(2, 2))
	fmt.Println(python, java)
	fmt.Println(start, end)
	variables()
	flowControl()
	deferStmt()
}
