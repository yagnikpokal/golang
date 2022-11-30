package main

import "fmt"

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

// prepend function is used to add the data of single node in linked list
func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// We can not print all the linked list without the function
// We can print the single adress of the linkedlist
//prepend function will print all the linked list
func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	// Creating the linked list
	myList := LinkedList{}

	// Add the data to linked list
	//node1 := &Node{data: 48}
	//node2 := &Node{data: 18}
	node3 := &Node{data: 16}

	// Print the linked list
	myList.prepend(&Node{data: 48})
	myList.prepend(&Node{data: 18})
	myList.prepend(node3)

	myList.printListData()

}
