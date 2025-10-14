package main

import ("fmt")

type Shape interface{
	Area() float64;
}

type Square struct{
	a float64;
}

type Rectangle struct{
	h, w float64;
}


func (s Square) Area() float64{
	return s.a * s.a;
}

func (r Rectangle) Area() float64{
	return r.h * r.w;
}

func printArea(s Shape){
	fmt.Println("s:",s)
	fmt.Println(s.Area())
}

func main(){
	s:= Square{a:10}
	printArea(s)

	r:= Rectangle{h:5,w:10}
	printArea(r)
}
