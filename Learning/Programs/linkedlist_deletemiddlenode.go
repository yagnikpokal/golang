package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(5)
	first.AddFront(4)
	first.AddFront(3)
	first.AddFront(2)
	first.AddFront(1)

	first.Head.Traverse()
	deleteMiddle(first.Head)
	fmt.Println("")
	first.Head.Traverse()

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

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	size := sizeOfList(head)
	mid := size / 2
	if mid == 0 {
		return head.Next
	}
	curr := head
	for i := 0; i < mid-1; i++ {
		curr = curr.Next
	}
	prev := curr
	midNode := prev.Next
	if midNode == nil {
		return head
	}
	midNext := midNode.Next
	prev.Next = midNext
	return head
}

func sizeOfList(head *ListNode) int {
	l := 0
	for head != nil {
		l = l + 1
		head = head.Next
	}
	return l
}
