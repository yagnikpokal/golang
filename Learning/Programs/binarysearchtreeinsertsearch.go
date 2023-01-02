package main

import (
	"fmt"
)

// Node represents the compinents of the binary search tree
type Node struct {
	key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

// Insert will add a node to the tree
// The key to add shuld not be allready
func (n *Node) Insert(k int) {
	if n.key < k {
		// Move right
		if n.Right == nil {
			n.Right = &Node{key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.key > k {
		// Move left
		if n.Left == nil {
			n.Left = &Node{key: k}

		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true is there is a node with that value
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.key < k {
		// Move right
		return n.Right.Search(k)
	} else if n.key > k {
		// Move left
		return n.Left.Search(k)
	}
	return true

}

func main() {
	tree := &Node{key: 100}
	//D := tree

	tree.Insert(-20)
	tree.Insert(-50)
	tree.Insert(-15)
	tree.Insert(-60)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(55)
	tree.Insert(85)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(-10)

	//tree.printPreOrder(key)
	//	fmt.Println(tree)
	fmt.Println(tree.Search(71))

}
