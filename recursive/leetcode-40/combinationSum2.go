package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	findCombination(candidates, 0, target, []int{},
		&result)
	return result
}

func findCombination(nums []int, start int, target int,
	combination []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
	}
	//used := make(map[int]int)
	for i := start; i < len(nums) && target > 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		findCombination(nums, i+1, target-nums[i],
			append(combination, nums[i]), result)
	}
}

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(nums, 8))

}
