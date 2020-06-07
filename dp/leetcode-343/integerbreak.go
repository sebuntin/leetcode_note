package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], dp[i-j]*j, (i-j)*j)
		}
	}
	return dp[n]
}

func Max(i, j, k int) int {
	tempMax := i
	if i < j {
		tempMax = j
	}
	if tempMax < k {
		return k
	}
	return tempMax
}

func main() {
	fmt.Println(integerBreak(10))
}
