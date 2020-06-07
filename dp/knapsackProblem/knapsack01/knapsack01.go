package knapsack01

/*
记忆化搜索方法解决01背包问题
*/

func knapSack(w []int, v []int, c int) int {
	n := len(v)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, c+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return bestValue(w, v, n-1, c, memo)
}

func bestValue(w []int, v []int, index, c int, memo [][]int) int {
	if c <= 0 || index < 0 {
		return 0
	}
	if memo[index][c] != -1 {
		return memo[index][c]
	}
	res := bestValue(w, v, index-1, c, memo)
	if w[index] <= c {
		res = Max(res, v[index]+bestValue(w, v, index-1, c-w[index], memo))
	}
	memo[index][c] = res
	return res
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
