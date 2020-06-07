package main

func orangesRotting(grid [][]int) int {
	visit := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		placeHolder := make([]int, len(grid[0]))
		visit[i] = placeHolder
	}
	maxTime := -1
	// 定义搜索队列
	searchQ := make([][]int, 0)
	// 初始化搜索队列
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				searchQ = append(searchQ, []int{i, j, 0})
			}
		}
	}
	maxTime = BFScount(searchQ, grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return maxTime
}

func BFScount(Q [][]int, grid [][]int) int {
	// 定义方向数组
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
	maxTime := 0
	for len(Q) != 0 {
		X := Q[0][0]
		Y := Q[0][1]
		T := Q[0][2]
		Q = Q[1:]
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				newX := X + dx[i]
				newY := Y + dy[i]
				if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid[0]) ||
					grid[newX][newY] == 2 || grid[newX][newY] == 0 {
					continue
				}
				Q = append(Q, []int{newX, newY, T + 1})
				grid[newX][newY] = 2
				maxTime = T + 1
			}
		}
	}
	return maxTime
}