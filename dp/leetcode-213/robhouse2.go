package leetcode_213

import "math"

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return robHelper(nums, 0, &memo)
}

func robHelper(nums []int, start int, memo *[]int) int {
	if start >= len(nums)-1 {
		return 0
	}
	sumIncome := 0
	if (*memo)[start] != -1 {
		return (*memo)[start]
	}
	for i := start; i < len(nums); i++ {
		if i == 0 {
			sumIncome = Max(sumIncome, robHelper(nums[:len(nums)-1], i+2, memo)+nums[i])
			continue
		}
		sumIncome = Max(sumIncome, robHelper(nums, i+2, memo)+nums[i])
	}
	(*memo)[start] = sumIncome
	return sumIncome
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
