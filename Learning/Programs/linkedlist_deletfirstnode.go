package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int, next *Node) *Node {
	var n Node
	n.value = value
	n.next = next
	return &n
}
func TraverseLinkedList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}
func DeleteFirstNode(head *Node) *Node {
	if head == nil {
		return head
	}
	newHead := head.next
	head.next = nil
	return newHead
}
func main() {
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = DeleteFirstNode(head)
	fmt.Printf("After deleting first node of the linked list: ")
	TraverseLinkedList(head)
}
