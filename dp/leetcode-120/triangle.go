package leetcode_120

import "math"

func minimumTotal(triangle [][]int) int {
	return bottomUp(triangle)
}

func bottomUp(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	// 初始化dp数组
	for i := 0; i < len(triangle); i++ {
		level := make([]int, len(triangle[i]))
		dp[i] = level
	}
	copy(dp[len(triangle)-1], triangle[len(triangle)-1])
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = int(math.Min(float64(dp[i+1][j]), float64(dp[i+1][j+1]))) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func topDown(triangle [][]int) (minsteps int) {
	// 边界条件
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle))
	// 初始化dp数组
	for i := 0; i < len(triangle); i++ {
		level := make([]int, len(triangle[i]))
		dp[i] = level
	}
	copy(dp[:1], triangle[:1])
	for i := 1; i < len(triangle)-1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if j > 0 && j < len(triangle[i])-1 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1]))) + triangle[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}
	// 找到dp最后一行的最小值
	minsteps = dp[len(dp)-1][0]
	for _, steps := range dp[len(dp)-1] {
		if minsteps > steps {
			minsteps = steps
		}
	}
	return
}
