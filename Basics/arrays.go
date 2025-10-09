package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	//Declaring
	var a [2]string
	a[0] = "Swaroop"
	a[1] = "Acharya"
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	//Slice
	evens := [6]int{2, 4, 6, 8, 10, 12}

	var s []int = evens[1:4]
	fmt.Println(s)

	//Slices are references to arrays

	var names = [4]string{"S", "W", "T", "K"}

	s1 := names[0:2]
	s2 := names[1:3]

	s2[0] = "XXX"

	fmt.Println(s1, s2)
	fmt.Println(names)

	//Slice literals
	//Ex: nothing but Dynamic array

	b := []bool{false, true, false}

	fmt.Println(b)

	st := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
	}
	fmt.Println(st)

	// Empty slice
	emp := []int{}
	fmt.Println(emp)

	//Slice defaults

	sa := []int{2, 3, 5, 7, 11, 13}
	sa = sa[1:4]
	fmt.Println(s)

	sa = sa[:2]
	fmt.Println(s)

	sa = sa[1:]
	fmt.Println(s)

	sl := []int{1, 2, 3, 4, 5}
	printSlice(sl)

	sl = sl[:0] //Slice to give it zero length
	printSlice(sl)

	sl = sl[:4] // Extend the length
	printSlice(sl)

	sl = sl[2:]
	printSlice(sl)

}
