package main

import (
	"fmt"
)

const TRIE_MAX_CHAR_NUM = 26

type TrieNode struct {
	Child []*TrieNode
	IsEnd bool
}

func Constructor() TrieNode {
	return TrieNode{make([]*TrieNode, TRIE_MAX_CHAR_NUM), false}
}

func (this *TrieNode) Insert(word string) {
	ptr := this
	for i := 0; i < len(word); i++ {
		if ptr.Child[int(word[i]-'a')] == nil {
			ptr.Child[int(word[i]-'a')] = &TrieNode{make([]*TrieNode, TRIE_MAX_CHAR_NUM), false}
		}
		ptr = ptr.Child[int(word[i]-'a')]
		if i == len(word)-1 {
			ptr.IsEnd = true
		}
	}
}

func (this *TrieNode) Search(word string) bool {
	var pos int
	ptr := this
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			pos = int(word[i] - 'a')
			if ptr.Child[pos] == nil {
				return false
			}
			ptr = ptr.Child[pos]
		} else if word[i] == '.' && i == len(word)-1 {
			for _, p := range ptr.Child {
				if p != nil && p.IsEnd {
					return true
				}
			}
			return false
		} else if word[i] == '.' && i < len(word)-1 {
			for _, p := range ptr.Child {
				if p != nil && p.Search(word[i+1:]) {
						return true
				}
			}
			return false
		}
	}
	return ptr.IsEnd
}

func (this *TrieNode) StartWith(word string) bool {
	ptr := this
	for i := 0; i < len(word); i++ {
		if ptr.Child[int(word[i]-'a')] == nil {
			return false
		}
		ptr = ptr.Child[int(word[i]-'a')]
	}
	return true
}

func getAllWord(node *TrieNode, byteStack *[]byte, result *[]string) {
	for i := 0; i < len(node.Child); i++ {
		if node.Child[i] != nil {
			//for j := 0; j < layer; j++ {
			//	fmt.Print("---")
			//}
			*byteStack = append(*byteStack, byte(i)+'a')
			if node.Child[i].IsEnd {
				//fmt.Println(string(byte(i) + 'a') + "(end)")
				*result = append(*result, string(*byteStack))
			}
			getAllWord(node.Child[i], byteStack, result)
			*byteStack = (*byteStack)[1:]
		}
	}
}

func main() {
	root := Constructor()
	n1 := Constructor()
	n2 := Constructor()
	n3 := Constructor()
	root.Child[int('a'-'a')] = &n1
	root.Child[int('b'-'a')] = &n2
	root.Child[int('e'-'a')] = &n3
	n2.IsEnd = true

	n4 := Constructor()
	n5 := Constructor()
	n6 := Constructor()
	n1.Child[int('b'-'a')] = &n4
	n2.Child[int('c'-'a')] = &n5
	n3.Child[int('f'-'a')] = &n6
	n5.IsEnd = true

	byteStack := make([]byte, 0)
	result := make([]string, 0)
	root.Insert("hello")
	getAllWord(&root, &byteStack, &result)
	fmt.Println(result)
	fmt.Println(root.Search("hell."))
	fmt.Println(root.StartWith("hell"))
}
