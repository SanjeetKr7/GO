package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	list []int
}

func (s *Stack) Push(x int) {
	s.list = append(s.list, x)
}

func (s *Stack) Pop() (int, error) {
	l := len(s.list)
	if l == 0 {
		return -1, errors.New("no value is present")
	}
	top := s.list[l-1]
	s.list = s.list[:l-1]
	return top, nil

}

func main() {
	s := Stack{}
	fmt.Println("push data ", s.list)
	s.Push(1)
	fmt.Println("push data ", s.list)
	top, err := s.Pop()
	if err != nil {
		fmt.Println("No data is present")
	} else {
		fmt.Println(" Top values is ", top)
	}
	fmt.Println("push data ", s.list)

	top, err = s.Pop()
	if err != nil {
		fmt.Println("No data is present")
	} else {
		fmt.Println(" Top values is ", top)
	}

	fmt.Println("push data ", s.list)

	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println("push data ", s.list)

	top, err = s.Pop()
	if err != nil {
		fmt.Println("No data is present")
	} else {
		fmt.Println(" Top values is ", top)
	}
	fmt.Println("push data ", s.list)
}
