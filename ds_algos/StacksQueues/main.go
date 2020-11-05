package main

import "fmt"

func main() {
	// var s stack

	// s.push(1)
	// s.push(2)
	// s.push(888)

	// printStack(s)

	// s.pop()

	// fmt.Println("********************")

	// printStack(s)

	var q queue

	q.push(1)
	q.push(2)
	q.push(3)

	// fmt.Println(q.first.data)
	printQueue(q)

	q.pop()

	fmt.Println("********************")

	printQueue(q)

}
