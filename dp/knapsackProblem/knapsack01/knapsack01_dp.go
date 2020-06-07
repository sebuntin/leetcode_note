package knapsack01

/*
动态规划方法解决01背包问题
*/
func knapSack01Dp(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("weight array must have the same size with value array.")
	}
	if len(w) == 0 {
		return 0
	}
	dp := make([]int, c+1)
	for index := 0; index < len(w); index++ {
		for j := c; j >= w[index]; j-- {
			 dp[j] = Max(dp[j-w[index]]+v[index], dp[j])
		}
	}
	return dp[c]
}
