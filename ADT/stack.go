package main

import "errors"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	poppedItem := s.items[len(s.items)-1]
	s.items = s.items[:poppedItem]
	return poppedItem, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func main() {}
