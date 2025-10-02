package main

import "fmt"

func add(x int, y int) int {
  return x+y;
}

// return multiple return results
func swap(x, y string) (string, string){
  return y,x;
}

// naked return
func addNsub(x,y int)(add,sub int){
	add= x+y;
	sub= x-y;
	return;
}

func main()  {
	fmt.Println(add(1,2))
	a,b:= swap("hellow","world");
	fmt.Println(a,b);
	fmt.Println(addNsub(2,2));
}

