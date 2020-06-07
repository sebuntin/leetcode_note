package main

import "fmt"

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
*/

func partition(s string) [][]string {
	result := make([][]string, 0)
	if len(s) == 0 {
		return result
	}
	partition := make([]string, 0)
	findPartition([]byte(s), partition, &result)
	return result
}

func findPartition(s []byte, partition []string, result *[][]string) {
	if len(s) == 0 {
		*result = append(*result, partition)
		return
	}
	part := []byte{}
	for i := range s {
		part = append(part, s[i])
		if !isPalindrome(part) {
			continue
		}
		findPartition(s[i+1:], append(partition, string(part)), result)
	}
}

// 判断字符串是否为回文串
func isPalindrome(s []byte) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func testStr(s []byte) {
	s[0] = 'f'
}

func main() {
	fmt.Println(partition("aabcadabaacbb"))
}
