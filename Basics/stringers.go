package main

import (
	"fmt"
)

type Employee struct{
	name string
	empId int
}

func (p Employee) String() string{
	return fmt.Sprintf("Emp name: %v\n Emp Id: %d\n",p.name,p.empId)
}

func main(){
	e1:= Employee{"SWAROOP", 12}
	fmt.Println(e1)
}
