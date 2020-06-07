package leetcode_174

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
	col := len(dungeon[0])
	if row == 0 || col == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		placeHolder := make([]int, col)
		dp[i] = placeHolder
	}
	// 初始化dp数组
	dp[row-1][col-1] = int(math.Max(1, float64(1-dungeon[row-1][col-1])))
	//if dungeon[row-1][col-1] >= 1 {
	//	dp[row-1][col-1] = 1
	//} else {
	//	dp[row-1][col-1] = 1 - dungeon[row-1][col-1]
	//}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i == row-1 && j < col-1 {
				dp[i][j] = int(math.Max(1, float64(dp[i][j+1] - dungeon[i][j])))
				//if dungeon[i][j] >= dp[i][j+1] {
				//	dp[i][j] = 1
				//} else {
				//	dp[i][j] = dp[i][j+1] - dungeon[i][j]
				//}
			} else if j == col-1 && i < row-1 {
				dp[i][j] = int(math.Max(1, float64(dp[i+1][j] - dungeon[i][j])))
				//if dungeon[i][j] >= dp[i+1][j] {
				//	dp[i][j] = 1
				//} else {
				//	dp[i][j] = dp[i+1][j] - dungeon[i][j]
				//}
			} else if i < row-1 && j < col-1 {
				minNeed := int(math.Min(float64(dp[i][j+1]), float64(dp[i+1][j])))
				dp[i][j] = int(math.Max(1, float64(minNeed - dungeon[i][j])))
				//if dungeon[i][j] >= minNeed {
				//	dp[i][j] = 1
				//} else {
				//	dp[i][j] = minNeed - dungeon[i][j]
				//}
			}
		}
	}
	return dp[0][0]
}
