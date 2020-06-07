package main

import "sort"

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func groupAnagrams(strs []string) [][]string {
	wordMap := make(map[string]int, 1024)
	result := make([][]string, 0, 1024)
	x := 0
	for _, word := range strs {
		if index, ok := wordMap[hashFunc(word)]; ok {
			result[index] = append(result[index], word)
		} else {
			wordMap[hashFunc(word)] = x
			x += 1
		}
	}
	return result
}

type Word []byte

func (this Word) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Word) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Word) Len() int {
	return len(this)
}

func hashFunc(word string) string {
	var wordByte Word
	wordByte = []byte(word)
	sort.Sort(wordByte)
	return string(wordByte)
}
