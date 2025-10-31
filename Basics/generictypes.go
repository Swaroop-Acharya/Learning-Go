package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(elem T) {
	s.items = append(s.items, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Size()int {
	return len(s.items)
}

func main() {
	s:= Stack[int]{}
	s.Push(10)
	s.Push(20)
	fmt.Println(s.Size())
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
		
	s1:= Stack[string]{}
	s1.Push("Hi")
	s1.Push("SWAROOp")
	fmt.Println(s1.Size())
	fmt.Println(s1.Peek())
	s1.Pop()
	fmt.Println(s1.Peek())
	s1.Pop()
	fmt.Println(s1.Peek())
	
}
