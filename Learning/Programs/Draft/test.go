package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(1)
	first.AddFront(2)
	first.AddFront(3)
	first.AddFront(4)
	second := initList()
	second.AddFront(1)
	second.AddFront(2)
	second.AddFront(3)
	second.AddFront(4)

	first.Head.Traverse()
	fmt.Println("\n")
	first.Head.Traverse()
	//myList := LinkedList{}
	myList := linksort(&first, &second)
	//myList.Head.Traverse()

	//first.Reverse()
	//	fmt.Println("")
	//	first.Head.Traverse()

}

func initList() *SingleList {
	return &SingleList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) Reverse() {

	curr := s.Head
	var prev *ListNode
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	s.Head = prev
}
func (s *SingleList) AddFront(num int) {
	ele := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}
	s.Len++
}

// Contains the data and adress of next node
type Node struct {
	data int
	next *Node
}

// LinkedList contains head adress and length of the LinkedList however length is not necessory every time
type LinkedList struct {
	head   *Node
	length int
}

func linksort(first *ListNode, second *ListNode) *ListNode {
	//var dummy = new(Listnode)
	dummy := &ListNode{}
	var p = dummy
	for first != nil && second != nil {
		if first.Val < second.Val {
			p.Next = first
			first = first.Next
		} else {
			p.Next = second
			second = second.Next

		}
		p = p.Next

	}
	if first != nil {
		p.Next = first

	} else {
		p.Next = second
	}
	//dummy.Head.Traverse()

	return dummy.Next
}
