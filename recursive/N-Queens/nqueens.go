package main

import (
	"fmt"
)


// 初始化标记棋盘
func iniMark(n int) (mark [][]int, locationByte [][]byte) {
	mark = make([][]int, n)
	locationByte = make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mark[i] = append(mark[i], 0)
			locationByte[i] = append(locationByte[i], '.')
		}
	}
	return mark, locationByte
}

// 打印棋盘
func printMark(n int, mark [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", mark[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// 放置皇后
func putQueen(x, y, n int, mark [][]int) {
	// 定义方向数组
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	mark[x][y] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			new_x := x + i*dx[j]
			new_y := y + i*dy[j]
			// 跳过越界的位置
			if new_x >= 0 && new_x < n && new_y >= 0 && new_y < n {
				mark[new_x][new_y] = 1
			}
		}
	}
}

func byte2string(n int, b [][]byte) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, string(b[i]))
	}
	return s
}

func solveQueens(n int) [][]string {
	mark, locationByte := iniMark(n)
	result := make([][]string, 0)
	generate(0, n, mark, locationByte, &result)
	return result
}

func generate(k, n int, mark [][]int, locationByte [][]byte, result *[][]string){
	if k == n {
		location := byte2string(n, locationByte)
		*result = append(*result, location)
		return
	}
	// 在mark的第k行上放置皇后
	for i := 0; i < n; i++ {
		// 剪枝
		// 这一行已经有皇后了
		if mark[k][i] == 1 {
			continue
		}
		// 找到位置放置皇后
		// 记录当前mark
		tempMark := make([][]int, len(mark))
		for m := range tempMark {
			tempMark[m]= make([]int, len(mark[0]))
			copy(tempMark[m], mark[m])
		}
		putQueen(k, i, n, tempMark)
		mark[k][i] = 1
		// 同时更新location
		locationByte[k][i] = 'Q'
		generate(k+1, n, tempMark, locationByte, result)
		// 回溯,在该行找到能放置皇后的下一个位置
		mark[k][i] = 0
		locationByte[k][i] = '.'
	}
}

func main() {
	// 初始化一个棋盘
	n := 4
	result := solveQueens(n)
	fmt.Println(result)
}
