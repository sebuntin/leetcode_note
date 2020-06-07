package leetcode_279

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; (j * j) <= i; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-j*j] + 1
				continue
			}
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
