package main

import "fmt"

func add(x int, y int) int {
  return x+y;
}

// return multiple return results
func swap(x, y string) (string, string){
  return y,x;
}

func main()  {
	fmt.Println(add(1,2))
	a,b:= swap("hellow","world");
	fmt.Println(a,b);
}

