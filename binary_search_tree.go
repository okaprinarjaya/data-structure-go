package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var counter int

func (n *Node) Insert(newKey int) {
	if newKey > n.Key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: newKey}
			return
		} else {
			n.Right.Insert(newKey)
		}
	} else if newKey < n.Key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: newKey}
			return
		} else {
			n.Left.Insert(newKey)
		}
	}
}

func (n *Node) Search(key int) bool {
	counter++

	if n == nil {
		return false
	}

	if key > n.Key {
		// move right
		return n.Right.Search(key)
	} else if n.Key < key {
		// move left
		return n.Left.Search(key)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(276))
	fmt.Println(counter)

	fmt.Println(tree.Search(400))
	fmt.Println(counter)
}
