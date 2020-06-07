package main

func exist(board [][]byte, word string) bool {
	mark := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if wordSearch(board, mark, 0, 0, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func wordSearch(board [][]byte, mark [][]int, x, y int, word string) bool {
	if len(word) == 0 {
		return true
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for k := 0; k < 4; k++ {
		newX := x + dx[k]
		newY := y + dy[k]
		if newX < 0 || newY < 0 ||
			newX >= len(board) || newY >= len(board[0]) ||
			mark[newX][newY] == 1 || board[newX][newY] != word[0] {
			continue
		}
		mark[newX][newY] = 1
		if wordSearch(board, mark, newX, newY, word[1:]) {
			mark[newX][newY] = 0
			return true
		}
		mark[newX][newY] = 0
	}
	return false
}
