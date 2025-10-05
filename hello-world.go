package main

import "fmt"


//declaring variable
var python, java bool;

var start, end int = 1,3;

var (
 num int = 1
 pi float32 = 3.1
piFull float64 = 3.14159
	name string = "swaroop"
	flag bool = true
	ch rune = '😊'
)



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


func variables(){
	k:=3;
	fmt.Println(k)
	fmt.Println(pi,piFull);
	fmt.Println(ch);
	fmt.Printf("%c\n",ch);
}



func main()  {
	fmt.Println(add(1,2))
	a,b:= swap("hellow","world");
	fmt.Println(a,b);
	fmt.Println(addNsub(2,2));
	fmt.Println(python,java);
	fmt.Println(start,end);
	variables();
}

