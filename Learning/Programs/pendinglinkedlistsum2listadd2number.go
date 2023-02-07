package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) {
	for l1 != nil {
		D := l1.Val + l2.Val

		if D >= 9 {
			//12
			sess := D / 10
			fmt.Println(sess)
			sess1 := D % 10
			fmt.Println(sess1)
			Y := l2.Val + sess
			fmt.Println(Y)

		} else {
			fmt.Println(D)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}*/

func main() {
	// create linked list for number 342
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	// create linked list for number 465
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	// call addTwoNumbers and get the result linked list
	/*result := */
	addTwoNumbers(l1, l2)
	/*
		// print the result linked list
		for result != nil {
			fmt.Print(result.Val)
			result = result.Next
		}*/
}
