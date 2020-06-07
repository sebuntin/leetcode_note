package main

import "fmt"

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	if n <= 0 {
		return result
	}
	findCombination(n, k, 1, []int{}, &result)
	return result
}

func findCombination(n int, k int, start int, combination []int, result *[][]int) {
	if len(combination) == k {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}
	for i := start; i <= n-(k-len(combination))+1; i++ {
		findCombination(n, k, i+1, append(combination, i), result)
	}
}

func main() {
	fmt.Println(combine(5, 3))
}
