package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d\n", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[0] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, "\n")
	}
	return maze
}

type searchNode struct {
	x, y int // 表示坐标
	pre  int // 表示前驱在队列中的位置
	step int // 从起点到达该点的最短距离
}

func minPathForMaze(grid [][]int) [][][]int {
	// 获取迷宫的行和列
	ret := make([][][]int, 0)
	row := len(grid)
	col := len(grid[0])
	// 初始化标记矩阵
	visit := make([][]int, row)
	for i := 0; i < row; i++ {
		placeHolder := make([]int, col)
		visit[i] = placeHolder
	}
	// 0代表未被探索, -1代表不能到达
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				visit[i][j] = -1
			}
		}
	}
	pos := make([]int, 0)
	Q := make([]searchNode, 0)
	Q = append(Q, searchNode{0, 0, -1, 1})
	visit[0][0] = 1
	BFS(grid, visit, &Q, &pos)
	for i := 0; i < len(pos); i++ {
		path := make([][]int, 0)
		for node := Q[pos[i]]; node.pre != -1; node = Q[node.pre] {
			path = append(path, []int{node.x, node.y})
		}
		path = append(path, []int{0, 0})
		ret = append(ret, path)
	}
	return ret
}

func BFS(grid, visit [][]int, Q *[]searchNode, pos *[]int) {
	// 获取迷宫形状
	row := len(grid)
	col := len(grid[0])
	// 创建搜索队列
	visit[0][0] = 1
	loc := 0
	// 定义方向数组
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	// 记录最短路径长度
	var minpath int
	for loc < len(*Q) {
		cur_x, cur_y := (*Q)[loc].x, (*Q)[loc].y
		cur_step := (*Q)[loc].step
		if minpath != 0 && minpath < cur_step+1 {
			break
		}
		for i := 0; i < 4; i++ {
			next_x := cur_x + dx[i]
			next_y := cur_y + dy[i]
			if next_x < 0 || next_y < 0 || next_x >= row ||
				next_y >= col || visit[next_x][next_y] != 0 {
				continue
			}
			if next_x == row-1 && next_y == col-1 {
				if minpath == 0 {
					minpath = cur_step + 1
					*pos = append(*pos, len(*Q))
				}
			}
			visit[next_x][next_y] = cur_step + 1
			// 将该节点与前其前驱所在位置一起入队列
			*Q = append(*Q, searchNode{next_x, next_y, loc, cur_step + 1})
		}
		loc++
	}
}

func main() {
	maze := readMaze("./maze.in")
	for i:=0; i<len(maze); i++ {
		for j:=0; j<len(maze[0]); j++ {
			fmt.Print(maze[i][j], " ")
		}
		fmt.Println()
	}
	//grid := [][]int{{0, 1, 0, 0, 0},
	//				{0, 0, 0, 1, 0},
	//				{0, 1, 0, 1, 0},
	//				{1, 1, 1, 0, 0},
	//				{0, 1, 0, 0, 1},
	//				{0, 1, 0, 0, 0}}
	//
	ret := minPathForMaze(maze)
	fmt.Println(ret)
}
