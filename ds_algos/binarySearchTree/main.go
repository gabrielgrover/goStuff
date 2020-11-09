package main

import "fmt"

func main() {
	bs := newBinarySearchTree(1)

	bs.insert(2)
	bs.insert(7)
	bs.insert(4)

	fmt.Println(bs)
	fmt.Println(bs.right)
	fmt.Println(bs.right.right)

	fmt.Println("*************")

	b := bs.search(2)

	fmt.Println(b)
}
