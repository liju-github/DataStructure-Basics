package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

func (t *Trie) findNode(str string) *TrieNode {
	node := t.root
	for _, char := range str {
		if _, exists := node.children[char]; !exists {
			return nil
		}
		node = node.children[char]
	}
	return node
}

func (t *Trie) AutoComplete(prefix string) []string {
	node := t.findNode(prefix)
	if node == nil {
		return []string{}
	}

	var results []string
	t.dfs(node, prefix, &results)
	return results
}

func (t *Trie) dfs(node *TrieNode, currentPrefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, currentPrefix)
	}

	for char, childNode := range node.children {
		t.dfs(childNode, currentPrefix+string(char), results)
	}
}

func main() {
	trie := NewTrie()

	words := []string{"apple", "app", "apricot", "banana", "bat", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Test Search
	fmt.Println("Search 'apple':", trie.Search("apple"))
	fmt.Println("Search 'app':", trie.Search("app"))
	fmt.Println("Search 'appl':", trie.Search("appl"))

	// Test StartsWith
	fmt.Println("StartsWith 'app':", trie.StartsWith("app"))
	fmt.Println("StartsWith 'ban':", trie.StartsWith("ban"))
	fmt.Println("StartsWith 'cat':", trie.StartsWith("cat"))

	// Test AutoComplete
	fmt.Println("AutoComplete 'ap':", trie.AutoComplete("ap"))
	fmt.Println("AutoComplete 'ba':", trie.AutoComplete("ba"))
	fmt.Println("AutoComplete 'c':", trie.AutoComplete("c"))
}
