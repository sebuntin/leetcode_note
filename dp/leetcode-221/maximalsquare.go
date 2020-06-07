package leetcode_221

import "math"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int, row)
	for i:=0; i<row; i++ {
		placeHolder := make([]int, col)
		dp[i] = placeHolder
	}
	// 为dp边界赋初值
	maxEdge := 0
	for i:=0; i<col; i++ {
		if matrix[0][i] == '0' {continue}
		if row < 2 {return 1}
		dp[0][i] = 1
		if maxEdge < dp[0][i] {maxEdge = dp[0][i]}
	}
	for i:=0; i<row; i++ {
		if matrix[i][0] == '0' {continue}
		if col < 2 {return 1}
		dp[i][0] = 1
		if maxEdge < dp[i][0] {maxEdge = dp[i][0]}
	}
	for i:=1; i<row; i++ {
		for j:=1; j<col; j++ {
			if matrix[i][j] == '0' {continue}
			minEdge := math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))
			minEdge = math.Min(float64(dp[i-1][j-1]), minEdge)
			dp[i][j] = int(minEdge) + 1
			if dp[i][j] > maxEdge {maxEdge = dp[i][j]}
		}
	}
	return maxEdge * maxEdge
}