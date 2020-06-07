/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1:

输入:
11110
11010
11000
00000
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
*/

package main

func numIslands(grid [][]byte) (count int){
	// 得到grid形状
	grid_x := len(grid)
	grid_y := len(grid[0])
	// 定义并初始化mark数组
	mark := make([][]byte, 0, grid_x)
	for i := 0; i < grid_x; i++ {
		row := make([]byte, 0, grid_y)
		for j := 0; j < grid_y; j++ {
			row = append(row, '0')
		}
		mark = append(mark, row)
	}

	// 遍历grid
	for i := 0; i < grid_x; i++ {
		for j := 0; j < grid_y; j++ {
			if mark[i][j] == '0' && grid[i][j] == 1 {
				count ++
				DFS(i, j, grid, mark)
			}
		}
	}
	return count
}

// 深度优先搜索
func DFS(x, y int, grid, mark [][]byte) {
	if grid[x][y] == '0' {
		return
	}
	mark[x][y] = '1'
	//定义方向数组, 上下左右
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	// 根据grid的边长进行深搜
	for j := 0; j < 4; j++ {
		new_x := x + dx[j]
		new_y := y + dy[j]
		if new_x > 0 || new_y > 0 ||
			new_x >= len(grid) || new_y >= len(grid[new_x]) {
			continue
		}
		if grid[new_x][new_y] == '1' && mark[new_x][new_y] == '0' {
			DFS(new_x, new_y, grid, mark)
		}
	}
}

// 宽度优先搜索
func BFS(x, y int, grid, mark [][]byte) {
	if grid[x][y] == '0' || mark[x][y] != '0' {
		return
	}
	// 定义方向数组, 上下左右
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	// 搜索队列
	searchQ := make([][]int, 0, 1024)
	// 将初始位置进队
	searchQ = append(searchQ, []int{x, y})
	for len(searchQ) != 0 {
		loc := searchQ[0]
		x = loc[0]
		y = loc[1]
		searchQ = searchQ[1:]
		mark[x][y] = '1'
		for j := 0; j < 4; j++ {
			new_x := x + dx[j]
			new_y := y + dy[j]
			if new_x > 0 || new_y > 0 ||
				new_x >= len(grid) || new_y >= len(grid[new_x]) {
				continue
			}
			if grid[new_x][new_y] == '1' && mark[new_x][new_y] == '0' {
				searchQ = append(searchQ, []int{new_x, new_y})
			}
		}
	}
}
