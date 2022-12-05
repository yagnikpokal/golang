package main

import (
	"fmt"
)

//Create prototype

// LL container which going to store list
type LL struct {
	list *linklist
}

// linklist for value and next pointer details
type linklist struct {
	val  int
	next *linklist
}

// createNode use for create node for list
func createNode(value int) *linklist {
	return &linklist{
		val:  value,
		next: nil,
	}
}

func (lstVal *LL) insertAtBeginning(data int) {
	if nil == lstVal.list {
		lstVal.list = createNode(data)
		return
	}

	tempNode := createNode(data)
	head := lstVal.list
	tempNode.next = head
	lstVal.list = tempNode
}

func (lstVal *LL) printList() {
	if nil != lstVal && nil != lstVal.list {
		head := lstVal.list
		for nil != head {
			fmt.Printf(" %d", head.val)
			head = head.next
		}
	}
	fmt.Println()
}

func (lstVal *LL) deleteFromBeginning() {
	if nil != lstVal && nil != lstVal.list {
		head := lstVal.list
		lstVal.list = head.next
		head = nil
	}
}

func sort(ll *linklist, insertedNode *linklist) *linklist {
	head := ll
	if ll.val > insertedNode.val {
		insertedNode.next = ll
		ll = insertedNode
	} else {
		for head.next != nil && head.next.val < insertedNode.val {
			head = head.next
		}
		insertedNode.next = head.next
		head.next = insertedNode
	}
	return ll
}

func sortTheLinkedList(ll *linklist) *linklist {
	head := ll
	sortedList := new(linklist)
	var firstTime bool = true
	for head != nil {
		nextNode := head.next
		if firstTime {
			sortedList = head
			sortedList.next = nil
			firstTime = false
		} else {
			sortedList = sort(sortedList, head)
		}
		head = nextNode
	}
	return sortedList
}

func main() {
	staticList := []int{5, 7, 1, 3, 4, 9}
	linklst := new(LL)
	for _, value := range staticList {
		linklst.insertAtBeginning(value)
	}
	fmt.Println("PrintList")
	linklst.printList()
	linklst.list = sortTheLinkedList(linklst.list)
	fmt.Println("After Sorting PrintList")
	linklst.printList()
}
