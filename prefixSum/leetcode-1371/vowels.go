package main

import "fmt"

/*
给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：
每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。
*/

func findTheLongestSubString(s string) int {
	// 定义hash表记录每种奇偶状态出现的最早位置
	table := make([]int, 1<<5)
	for i:= 1; i<len(table); i++ {
		table[i] = -1
	} 
	ans := 0
	state := 0
	for i := range s {
		switch s[i] {
		case 'a':
			state ^= 1 << 4
		case 'e':
			state ^= 1 << 3
		case 'i':
			state ^= 1 << 2
		case 'o':
			state ^= 1 << 0
		}
		if table[state] >= 0 {
			ans = Max(ans, i - table[state] + 1) 
			continue
		}
		table[state] = i 
	}
	return ans
}

// Max a, b
func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	testcase := "leetcodeisgreat"
	fmt.Println(findTheLongestSubString(testcase))
}