package main

import "fmt"

type TrieNode struct {
	char   byte
	isWord bool
	adj    []*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) New() {
	t.root = &TrieNode{'0', false, make([]*TrieNode, 26)}
}

func (t *Trie) Insert(str string) {
	curr := t.root
	for i := 0; i < len(str); i++ {
		if curr.adj[str[i]-'a'] == nil {
			curr.adj[str[i]-'a'] = &TrieNode{str[i], false, make([]*TrieNode, 26)}
		}
		curr = curr.adj[str[i]-'a']
	}
	curr.isWord = true
}

func (t *Trie) GetNode(str string) *TrieNode {
	curr := t.root
	for i := 0; i < len(str); i++ {
		if curr.adj[str[i]-'a'] == nil {
			curr.adj[str[i]-'a'] = &TrieNode{str[i], false, make([]*TrieNode, 26)}
		}
		curr = curr.adj[str[i]-'a']
	}
	return curr
}

func (t *Trie) Search(str string) bool {
	node := t.GetNode(str)
	return node.isWord
}

func main() {
	var trie Trie
	trie.New()

	trie.Insert("hamza")
	trie.Insert("hammer")

	fmt.Println(trie.Search("hamza"))
	fmt.Println(trie.Search("hamzaa"))
	fmt.Println(trie.Search("hammer"))

}
