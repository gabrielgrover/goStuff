package main

import "fmt"

type stack struct {
	length int
	top    *node
}

func (s *stack) push(elem int) {
	n := s.top
	top := node{
		data: elem,
		next: n,
	}
	s.length++
	s.top = &top
}

func (s *stack) pop() int {
	data := s.top.data
	s.top = s.top.next
	s.length--

	return data
}

func printStack(s stack) {
	p := s.top

	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
