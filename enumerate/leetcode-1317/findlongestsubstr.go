package main

import (
	"math"
)

func findLongestSubString(s string) int {
	var state, ans int
	// 定义hashtable记录每种奇偶性最早出现的位置
	table := make([]int, 1<<5)
	for i := 1; i < len(table); i++ {
		table[i] = -1
	}
	// 变量字符串
	for i := range s {
		switch s[i] {
		case 'a':
			state ^= 1 << 4
		case 'e':
			state ^= 1 << 3
		case 'i':
			state ^= 1 << 2
		case 'o':
			state ^= 1 << 1
		case 'u':
			state ^= 1 << 0
		}
		if table[state] >= 0 {
			ans = int(math.Max(float64(ans), float64(i-table[state]+1)))
			continue
		}
		table[state] = i + 1
	}
	return ans
}
