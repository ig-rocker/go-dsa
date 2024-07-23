## Stack
Stack is a data structure which store the elements in LIFO(Last In First Out) order, where the last element to be stored is removed first. Stack always push, pop, and top operation to add, remove and get data respectively.

One way to visualise is a pile of books where you put the book on top and take out the book from top only.

For Stack data structure, there is mainly four methods, i.e., Push, Pop, Top and Size.

# Implementation Using Slice:
```
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

```

# Complexity
Time Complexity: As we only adding the data according to the value of top, so all the operations have O(1) size.

Memory Complexity: As we dont have using extra memory, so it will O(1) complexity.
The issue with Stack using Array is it needs to have the fixed memory size regardless of whether we are using them or not.