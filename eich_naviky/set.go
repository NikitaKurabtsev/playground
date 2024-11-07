package main

import (
	"fmt"
	"log"
)

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

func (s *Set) Add(value string) {
	if isEmptyValue(value) {
		return
	}

	s.elements[value] = struct{}{}
}

func (s *Set) Remove(value string) {
	if isEmptyValue(value) {
		return
	}

	delete(s.elements, value)
}

func (s *Set) Contains(value string) bool {
	if _, ok := s.elements[value]; ok {
		return true
	}
	return false
}

func (s *Set) List() []string {
	if len(s.elements) == 0 {
		return nil
	}

	list := make([]string, 0, len(s.elements))

	for key := range s.elements {
		list = append(list, key)
	}
	return list
}

func isEmptyValue(value string) bool {
	if len(value) == 0 {
		log.Println("value cannot be empty")
		return true
	}
	return false
}

func main() {
	set := NewSet()
	set.Add("apple")
	set.Add("banana")
	fmt.Println("Contains 'apple':", set.Contains("apple"))
	fmt.Println("Set elements:", set.List())
	set.Remove("banana")
	fmt.Println("Set after removal:", set.List())
}
