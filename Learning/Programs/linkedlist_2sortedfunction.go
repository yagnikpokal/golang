package main

import "fmt"

// Listnode represents a node in a linked list
type Listnode struct {
	Val  int
	Next *Listnode
}

// linksort sorts two linked lists and returns a new sorted linked list
func linksort(l1 *Listnode, l2 *Listnode) *Listnode {
	var dummy = new(Listnode)
	var p = dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return dummy.Next
}

func main() {
	// Create two lists to sort
	l1 := &Listnode{
		Val: 1,
		Next: &Listnode{
			Val: 3,
			Next: &Listnode{
				Val: 5,
			},
		},
	}
	l2 := &Listnode{
		Val: 2,
		Next: &Listnode{
			Val: 4,
			Next: &Listnode{
				Val: 6,
			},
		},
	}

	// Sort the lists
	sortedList := linksort(l1, l2)

	// Print the sorted list
	for curr := sortedList; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}
}

/*
go run linkedlist_2sortedfunction.go
1
2
3
4
5
6
*/
