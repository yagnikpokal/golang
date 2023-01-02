package main

import "fmt"

//Declaring trie_Node  for creating node in a trie
type trie_Node struct {
	//assigning limit of 26 for child nodes
	childrens [26]*trie_Node
	//declaring a bool variable to check the word end.
	wordEnds bool
}

//Initializing the root of the trie
type trie struct {
	root *trie_Node
}

//inititlaizing a new trie
func trieData() *trie {
	t := new(trie)
	t.root = new(trie_Node)
	return t
}

//Passing words to trie
func (t *trie) insert(word string) {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(trie_Node)
		}
		current = current.childrens[index]
	}
	current.wordEnds = true
}

//Initializing the search for word in node
func (t *trie) search(word string) int {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return 0
		}
		current = current.childrens[index]
	}
	if current.wordEnds {
		return 1
	}
	return 0
}

//initializing the main function
func main() {
	trie := trieData()
	//Passing the words in the trie
	word := []string{"and", "ant", "dad", "do"}
	for _, wr := range word {
		trie.insert(wr)
	}
	//initializing search for the words
	words_Search := []string{"and", "ant", "dad", "do", "cat", "dog", "can"}
	for _, wr := range words_Search {
		found := trie.search(wr)
		if found == 1 {
			fmt.Printf("\"%s\"Word found in trie\n", wr)
		} else {
			fmt.Printf(" \"%s\" Word not found in trie\n", wr)
		}
	}
}

/*
go run trieinsertesearch.go
"and"Word found in trie
"ant"Word found in trie
"dad"Word found in trie
"do"Word found in trie
 "cat" Word not found in trie
 "dog" Word not found in trie
 "can" Word not found in trie


*/
