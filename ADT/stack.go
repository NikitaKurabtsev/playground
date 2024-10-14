package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	poppedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return poppedItem, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func main() {
	stack := &Stack{}

	fmt.Println(stack.IsEmpty())

	stack.Push(1)
	stack.Push(3)

	stackPeek, err := stack.Peek()
	if err != nil {
		err.Error()
	}
	fmt.Printf("Peek %d\n", stackPeek)

	popItem, err := stack.Pop()
	if err != nil {
		err.Error()
	}
	fmt.Printf("Popped item: %d\n", popItem)

	popItem, err = stack.Pop()
	if err != nil {
		err.Error()
	}
	fmt.Printf("Popped item: %d\n", popItem)
}
