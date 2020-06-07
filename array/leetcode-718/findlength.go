package leetcode_718

/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
*/

func findLength(a []int, b []int) int {
	dp := make([][]int, len(b))
	for i := range dp {
		dp[i] = make([]int, len(a))
	}
	maxLen := 0
	for i := range b {
		for j := range a {
			if a[j] == b[i] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			}
			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}
