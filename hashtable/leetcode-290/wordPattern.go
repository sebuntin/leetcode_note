package main

/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	wordSlice := strings.Split(str, " ")
	if len(pattern) != len(wordSlice) {
		return false
	}
	p2s := make(map[uint8]string, 1024)
	s2p := make(map[string]uint8, 1024)

	for i := 0; i < len(wordSlice); i++ {
		if word, ok := p2s[pattern[i]]; ok {
			if word != wordSlice[i] {
				return false
			}
		} else {
			if p, ok := s2p[wordSlice[i]]; ok {
				return false
			} else {
				p2s[pattern[i]] = wordSlice[i]
				s2p[wordSlice[i]] = p
			}
		}
	}
	return true
}

func main() {
	for {
		pattern, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		str, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		fmt.Println(wordPattern(string(pattern), string(str)))
	}
}
