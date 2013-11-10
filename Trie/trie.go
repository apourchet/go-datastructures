package trie

import (
	"fmt"
	"unsafe"
)

type ContainsFunction func(n1, n2 uintptr) bool
type EqualsFunction func(n1, n2 uintptr) bool

type Trie struct {
	Contains ContainsFunction
	Equals   EqualsFunction
	Head     *TrieNode
	Count    int
}

type TrieNode struct {
	Value    uintptr // This can be a pointer to a string if neeed be
	Children []*TrieNode
}

func Construct(superSet uintptr, contains ContainsFunction, equals EqualsFunction) *Trie {
	trie := Trie{}
	trie.Count = 1
	trie.Contains = contains
	trie.Equals = equals
	trie.Head = &TrieNode{superSet, []*TrieNode{}}
	return &trie
}

func (trie *Trie) Insert(insertion uintptr) {
	trie.Count++
	trie.Head.insert(trie.Contains, insertion)
}

func (trieNode *TrieNode) insert(contains ContainsFunction, insertion uintptr) {
	for _, child := range trieNode.Children {
		if contains(insertion, child.Value) {
			child.insert(contains, insertion)
			return
		}
	}
	newNode := TrieNode{insertion, []*TrieNode{}}
	trieNode.Children = append(trieNode.Children, &newNode)
}

func (trie *Trie) Remove(deletion uintptr) {
	trie.Count--
	trie.Head.remove(trie.Contains, trie.Equals, deletion)
}

func (trieNode *TrieNode) remove(contains ContainsFunction, equals EqualsFunction, deletion uintptr) {
	if len(trieNode.Children) == 0 {
		fmt.Println("No such node to remove")
		return
	}
	newChildren := []*TrieNode{}
	for _, child := range trieNode.Children {
		if contains(deletion, child.Value) {
			if equals(deletion, child.Value) {
				for _, subChild := range child.Children {
					newChildren = append(newChildren, subChild)
				}
			} else {
				child.remove(contains, equals, deletion)
				return
			}
		} else {
			newChildren = append(newChildren, child)
		}
	}
	trieNode.Children = newChildren
}

func (trie *Trie) Depth() int {
	return trie.Head.depth()
}

func (trieNode *TrieNode) depth() int {
	if len(trieNode.Children) == 0 {
		return 0
	}
	maxDepth := 0
	for _, child := range trieNode.Children {
		thisDepth := 1 + child.depth()
		if thisDepth > maxDepth {
			maxDepth = thisDepth
		}
	}
	return maxDepth
}

func (trie *Trie) Find(query uintptr) *TrieNode {
	return trie.Head.find(trie.Contains, trie.Equals, query)
}

func (trieNode *TrieNode) find(contains ContainsFunction, equals EqualsFunction, query uintptr) *TrieNode {
	if equals(query, trieNode.Value) {
		return trieNode
	}
	if !contains(query, trieNode.Value) {
		return nil
	}
	for _, child := range trieNode.Children {
		r := child.find(contains, equals, query)
		if r != nil {
			return r
		}
	}
	return nil
}

func StringToUintptr(input string) uintptr {
	return uintptr(unsafe.Pointer(&input))
}

func UintptrToString(address uintptr) string {
	return *((*string)(unsafe.Pointer(address)))
}
