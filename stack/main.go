package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	list []int
}

func (s *Stack)Push(val int){
	s.list=append(s.list, val)
}

func (s *Stack) Pop() (bool, error) {
	n := len(s.list)
	if n == 0 {
		return false, errors.New("Stack is Empty")
	}
	s.list = s.list[:n-1]
	return true, nil
}
func (s *Stack)Size()int{
	return len(s.list)
}
func (s *Stack) Print() {
	fmt.Printf("Stack data:-> \n")
	for _, val := range s.list {
		fmt.Println(val)
	}
}

func main() {
	s := Stack{
		list: []int{1,2,4,5,6},
	}
	s.Push(11)
	s.Push(55)
	n:=s.Size()
	fmt.Printf("size of stack %d \n",n)
	s.Pop()
	s.Print()
}
