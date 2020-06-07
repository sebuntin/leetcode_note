package leetcode_1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if text1[i] == text2[0] {
			dp[i][0] = 1
		}
	}
	for j := range dp[0] {
		if text2[j] == text1[0] {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
