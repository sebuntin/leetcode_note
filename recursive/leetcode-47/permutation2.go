package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	findPermutation(nums,  []int{}, &result)
	return result
}

func findPermutation(nums []int,  permutation []int, result *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(permutation))
		copy(temp, permutation)
		*result = append(*result, temp)
	}
	used := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		if used[nums[i]] == 1 {
			continue
		}
		nextNums := make([]int, len(nums)-1)
		copy(nextNums[:i], nums[:i])
		copy(nextNums[i:], nums[i+1:])
		findPermutation(nextNums, append(permutation, nums[i]), result)
		used[nums[i]] = 1
	}
}

func main() {
    nums := []int{1, 1, 2, 3}
	fmt.Println(permuteUnique(nums))
}

