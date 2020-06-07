package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindrome
*/

func longestPalindrome(s string) int {
	// 初始化一个字符hash表
	char_map := make([]int, 128)
	for _, char := range []byte(s) {
		char_map[int(char)]++
	}
	// 定义flag表示是否有中心字符
	centerFlag := 0
	longestLen := 0
	for _, count := range char_map {
		if count%2 == 0 {
			longestLen += count
		} else {
			longestLen += count - 1
			centerFlag = 1
		}
	}
	longestLen += centerFlag
	return longestLen
}

func main() {
	for {
		readLine, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		fmt.Println(longestPalindrome(string(readLine)))
	}
}
