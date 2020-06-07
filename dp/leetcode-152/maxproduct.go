package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxdp := make([]int, len(nums))
	mindp := make([]int, len(nums))
	maxProd := nums[0]
	maxdp[0] = nums[0]
	mindp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mintemp := math.Min(float64(mindp[i-1]*nums[i]), float64(nums[i]))
		maxtemp := math.Max(float64(maxdp[i-1]*nums[i]), float64(nums[i]))
		mindp[i] = int(math.Min(mintemp, float64(maxdp[i-1]*nums[i])))
		maxdp[i] = int(math.Max(maxtemp, float64(mindp[i-1]*nums[i])))
		if maxProd < maxdp[i] {
			maxProd = maxdp[i]
		}
	}
	return maxProd
}

func main() {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums))
}
