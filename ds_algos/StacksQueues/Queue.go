package main

import (
	"errors"
	"fmt"
)

type queue struct {
	length int
	first  *node
}

func (q *queue) push(data int) {
	last := node{
		data: data,
	}

	n := q.first

	q.length++

	if n == nil {
		q.first = &last
		return
	}

	for n.next != nil {
		n = n.next
	}

	n.next = &last
}

func (q *queue) pop() (int, error) {
	n := q.first

	if n == nil {
		return 0, errors.New("queue empty")
	}

	q.first = n.next
	q.length--

	return q.first.data, nil
	// return 1
}

func printQueue(q queue) {
	n := q.first

	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}
