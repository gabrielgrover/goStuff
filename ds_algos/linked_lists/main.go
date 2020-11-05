package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func main() {
	n0 := node{
		data: 0,
	}

	n1 := node{
		data: 1,
	}

	n2 := node{
		data: 2,
	}

	ll := newLinkedList(n0)

	ll.append(n1)

	ll.append(n2)

	// fmt.Println(ll.head.next.next)
	// fmt.Println(ll.length)

	ll = newLinkedList(n0)

	ll.prepend(n1)
	ll.prepend(n2)
	fmt.Println(ll.length)
	// fmt.Println(ll.head)
	// fmt.Println(ll.head.next)
	// fmt.Println(ll.head.next.next)
	ll.printListData()

	ll.deleteWithValue(1)

	fmt.Println("")
	ll.printListData()
}

func newLinkedList(h node) linkedList {
	return linkedList{
		head:   &h,
		length: 1,
	}
}

func (ll *linkedList) append(n node) {
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &n
	ll.length++
}

func (ll *linkedList) prepend(n node) {
	n.next = ll.head
	ll.head = &n
	ll.length++
}

func (ll linkedList) printListData() {
	currentNode := ll.head

	for currentNode.next != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (ll *linkedList) deleteWithValue(value int) {
	p1 := ll.head
	p2 := p1.next

	if p1.data == value {
		ll.head = p2
		return
	}

	for p2.data != value && p2 != nil {
		p1 = p2
		p2 = p2.next
	}

	if p2.data == value {
		p1.next = p2
	}

	ll.length--
}
