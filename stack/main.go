package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	list []int
	top int
}

func (s *Stack)Push(val int){
	s.list[s.top+1]=val
	s.top++
}

func (s *Stack) Pop() (bool, error) {
	if s.top== -1 {
		return false, errors.New("Stack is Empty")
	}
	s.top--
	return true, nil
}
func (s *Stack)Size()int{
	return s.top+1
}
func (s *Stack) Print() {
	fmt.Printf("Stack data:-> \n")
	for i:=0;i<=s.top;i++{
		fmt.Println(s.list[i])
	}
}

func main() {
	s := Stack{
		list: make([]int,10),
		top:-1,
	}
	s.Push(11)
	s.Push(55)
	n:=s.Size()
	fmt.Printf("size of stack %d \n",n)
	s.Print()
	s.Pop()
	s.Print()
}
