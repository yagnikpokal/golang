// Go Implementation of a Thread Safe Trie Data Structure and (some of) Trie Operations
package main

import (
	"fmt"
	"sync"
)

type Node struct {
	// list of all the children
	Children      []*Node
	m             *sync.RWMutex
	childIndexMap map[rune]int
	Val           rune

	// isEndOfWord is true if the node
	// represents end of a word
	IsEndOfWord bool
}

// CreateNode returns an initialized trie node
func CreateNode(v rune) *Node {
	return &Node{
		Children:      make([]*Node, 0),
		childIndexMap: make(map[rune]int),
		Val:           v,
		m:             new(sync.RWMutex),
	}
}

// AddChildNode add child node to current node with value v
func (n *Node) AddChildNode(v rune) *Node {
	node, exist := n.GetChildNode(v)
	if exist {
		return node
	}
	n.m.Lock()
	n.childIndexMap[v] = len(n.Children)
	n.m.Unlock()
	node = CreateNode(v)
	n.Children = append(n.Children, node)
	return node
}

// Len returns the number of children of node
func (n *Node) Len() int {
	n.m.RLock()
	l := len(n.childIndexMap)
	n.m.RUnlock()
	return l
}

// IsLeafNode returns true if current node is a leaf node in Trie
func (n *Node) IsLeafNode() bool {
	return n.Len() == 0
}

// GetChildNode retrieve child node with value v.
func (n *Node) GetChildNode(v rune) (node *Node, exist bool) {
	n.m.RLock()
	defer n.m.RUnlock()
	if i, ok := n.childIndexMap[v]; !ok {
		return nil, false
	} else {
		return n.Children[i], true
	}
}

// DeleteChildNode Deletes the child node if it exist
func (n *Node) DeleteChildNode(v rune) {
	n.m.Lock()
	defer n.m.Unlock()
	n.Children[n.childIndexMap[v]] = nil
	delete(n.childIndexMap, v)
}

// Trie Represent a node of trie
// Not advised to create the node directly. Use helper function CreateNode() instead
type Trie struct {
	Node
}

// New Creates an initialized trie data structure
func New() *Trie {
	return &Trie{
		*CreateNode(0),
	}
}

// Insert allow one or more keyword to be inserted in trie
// keyword can be any unicode string
func (t *Trie) Insert(keywords ...string) *Trie {
	for _, v := range keywords {
		t.insert(v)
	}

	return t
}

func (t *Trie) insert(keyword string) {
	node := &t.Node
	for _, v := range []rune(keyword) {
		node = node.AddChildNode(v)
	}
	node.IsEndOfWord = true
}

// PrefixSearch checks if keyword exist in trie as a keyword or prefix to a keyword
func (t *Trie) PrefixSearch(key string) (found bool) {
	node := &t.Node
	for _, v := range []rune(key) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			found = ok
			continue
		}
		return false
	}

	return found
}

// Search checks if keyword exist in trie as a fully qualified keyword.
func (t *Trie) Search(keyword string) (found bool) {
	node := &t.Node
	for _, v := range []rune(keyword) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			found = ok
			continue
		}
		return false
	}

	return found && node.IsEndOfWord
}

// Delete deletes a keyword from a trie if keyword exist in trie
func (t *Trie) Delete(keyword string) {
	node := &t.Node
	var breakNode *Node
	var breakRune rune
	for _, v := range []rune(keyword) {
		if n, ok := node.GetChildNode(v); ok {
			if node.IsEndOfWord {
				breakNode = node
				breakRune = v
			}
			node = n
			continue
		}
		return
	}

	if !node.IsEndOfWord {
		return
	}

	if !node.IsLeafNode() {
		node.IsEndOfWord = false
		return
	}

	if breakNode == nil {
		breakNode = &t.Node
		breakRune = []rune(keyword)[0]
	}

	breakNode.DeleteChildNode(breakRune)
}

// DeleteBranch deletes all child after last letter of key if key exists in trie
// If key is found, key will be treated as a keyword after this operation
func (t *Trie) DeleteBranch(key string) {
	node := &t.Node
	for _, v := range []rune(key) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			continue
		}
		return
	}

	node.childIndexMap = make(map[rune]int)
	node.IsEndOfWord = true
	node.Children = make([]*Node, 0)
}

func main() {
	trie := New().Insert("foo", "bar", "baz")
	trie.insert("food")
	fmt.Println(trie.PrefixSearch("food")) // returns true since we added food

	fmt.Println(trie.PrefixSearch("fo")) // returns true
	fmt.Println(trie.PrefixSearch("fb")) // returns false
	fmt.Println(trie.Search("fo"))       // returns false
	fmt.Println(trie.Search("foo"))      // returns true
	trie.Delete("foo")
	fmt.Println(trie.Search("foo")) // returns false because we delete the foo

	trie.DeleteBranch("ba")         // Delete bar and baz since we added ba then ba and after all items deleted
	fmt.Println(trie.Search("bar")) // returns false because we delete the bar
	fmt.Println(trie.Search("baz")) // returns false because we delete the baz
	fmt.Println(trie.Search("ba"))  // returns true because ba is present

}

/*
go run trieinsertsearchdeletebranchchildnodeprefixsearch.go
true
true
false
false
true
false
false
false
true
*/
