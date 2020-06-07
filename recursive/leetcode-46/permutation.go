package main

import "fmt"

func permute(nums []int) [][]int {
	used := make([]int, len(nums))
	result := make([][]int, 0)
	permutation := make([]int, 0, len(nums))
	findPermutation(nums, used, permutation, &result)
	return result
}

func findPermutation(nums []int, used, permutation []int, result *[][]int) {
	if len(permutation) == len(nums) {
		temp := make([]int, len(permutation))
		copy(temp, permutation)
		*result = append(*result, temp)
		return
	}
	for i := range nums {
		if used[i] == 0 {
			// 以nums[i]为第一个元素的排列
			used[i] = 1
			findPermutation(nums, used, append(permutation, nums[i]), result)
			used[i] = 0
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))

}
