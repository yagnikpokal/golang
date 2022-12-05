package main

import (
	"container/list"
	"fmt"
)

func main() {
	nodeA := list.New()
	nodeA.PushBack(10)
	nodeA.PushBack(4)
	nodeA.PushBack(12)
	nodeA.PushBack(8)

	nodeB := list.New()
	nodeB.PushBack(5)
	nodeB.PushBack(10)
	nodeB.PushBack(15)
	nodeB.PushFront(1)
	nodeB.PushBack(13)
	//iterate over nodeB
	for i := nodeB.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	sortLinkedList(nodeA)
	sortLinkedList(nodeB)
	fmt.Println("after sort, nodeA: ")
	for i := nodeA.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("after sort, nodeB: ")
	for i := nodeB.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// passing two sorted list to merge
	mergedList := mergeSortedLinkedList(nodeA, nodeB)
	fmt.Println("after merging two sorted list, mergedList: ")
	for i := mergedList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func sortLinkedList(node *list.List) *list.List {
	current := node.Front() //pointing to first node in the list
	if node.Front() == nil {
		return nil
	} else {
		for current != nil {
			index := current.Next() //pointing to second node in the list
			for index != nil {
				if current.Value.(int) > index.Value.(int) {
					//comparing and swaping the nodes
					temp := current.Value
					current.Value = index.Value
					index.Value = temp
				}
				index = index.Next() //increasing the pointer to next node
			}
			current = current.Next() //increasing the pointer to next node
		}
	}
	return node
}

func mergeSortedLinkedList(nodeA, nodeB *list.List) *list.List {
	//here i expect two sorted list of length > 0 passed
	//you can add more validation, leaving it for you..
	node1 := nodeA.Front() //pointing to first node of nodeA, HEAD
	node2 := nodeB.Front() //pointing to first node of nodeB, HEAD
	resNode := list.New()  //we will store and return our sorted merged list here
	for node1 != nil && node2 != nil {
		if node1.Value.(int) < node2.Value.(int) {
			resNode.PushBack(node1.Value)
			node1 = node1.Next()
		} else {
			resNode.PushBack(node2.Value)
			node2 = node2.Next()
		}
	}
	//what if node1.length > node2.length ? add remaining element of node1 to the result list
	for node1 != nil {
		resNode.PushBack(node1.Value)
		node1 = node1.Next()
	}
	// similarly what if node2.length > node1.length ? add remaining element of node2 to the result list
	for node2 != nil {
		resNode.PushBack(node2.Value)
		node2 = node2.Next()
	}

	return resNode
}
