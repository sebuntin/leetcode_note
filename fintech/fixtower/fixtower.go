package main

import (
	"fmt"
	"math"
	"sort"
)

func fixTower(nums []int, k int) int {
	// 对nums进行排序
	sort.Ints(nums)
	res1, res2 := 0, 0
	for i := 0; i < k; i++ {
		// 将所有高度较小的塔堆到正序第k大的塔高，所需要的次数为
		res1 += nums[k-1] - nums[k-1-i]
		// 将所有高度较大的塔拆到倒序序第k大的塔高，所需要的次数为
		res2 += nums[len(nums)-1-i] - nums[len(nums)-k]
	}
	for i:=k;i<len(nums);i++ {
		if nums[i] == nums[k-1] {
			res1 --
		}
		if nums[len(nums)-i-1] == nums[len(nums)-k] {
			res2 --
		}
	}
	return int(math.Max(0, math.Min(float64(res1), float64(res2))))
}

func main() {
    nums := []int{1, 2, 2, 4, 3, 2}
    fmt.Println(fixTower(nums, 5))
}
