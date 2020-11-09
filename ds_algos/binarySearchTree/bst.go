package main

// BinarySearchTree data structure
type BinarySearchTree struct {
	data  int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func newBinarySearchTree(data int) BinarySearchTree {
	return BinarySearchTree{
		data: data,
	}
}

func (bs *BinarySearchTree) insert(data int) {
	p0 := bs
	p1 := p0
	right := false

	for p1 != nil {
		p0 = p1
		if data < p1.data {
			p1 = p1.left
			right = false
		} else {
			p1 = p1.right
			right = true
		}
	}

	newNode := BinarySearchTree{
		data: data,
	}

	if right {
		p0.right = &newNode
	} else {
		p0.left = &newNode
	}
}

func (bs *BinarySearchTree) search(data int) *BinarySearchTree {
	n := bs

	for n != nil {
		if data < n.data {
			n = n.left
		} else if data > n.data {
			n = n.right
		} else {
			break
		}
	}

	return n
}
