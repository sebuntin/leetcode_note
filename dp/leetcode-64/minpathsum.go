package leetcode_64

import "math"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	// 求出grid的形状
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		placeHolder := make([]int, col)
		dp = append(dp, placeHolder)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
			}
		}
	}
	return dp[row-1][col-1]
}
