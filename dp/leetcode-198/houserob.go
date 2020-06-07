package leetcode_198

import "math"

/*
你是一个专业的小偷，计划偷窃沿街的房屋。
每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
*/

func rob(nums []int) (ret int) {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//dp := make([]int, 0)
	//dp = append(dp, nums[0])
	//dp = append(dp, int(math.Max(float64(dp[0]), float64(nums[1]))))
	// 空间优化
	pre2 := nums[0]
	pre1 := int(math.Max(float64(pre2), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		ret = int(math.Max(float64(pre1), float64(pre2+nums[i])))
		pre1, pre2 = ret, pre1
	}
	return ret
}

// 使用递归+记忆化搜索完成
func robHelper(nums []int) int {
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	ret := robFoo(nums, 0, &memo)
	return ret
}

func robFoo(nums []int, start int, memo *[]int) int {
	if start > len(nums)-1 {
		return 0
	}
	if (*memo)[start] != -1 {
		return (*memo)[start]
	}
	sumIncome := 0
	for i := start; i < len(nums); i++ {
		sumIncome = Max(sumIncome, nums[i]+robFoo(nums, i+2, memo))
	}
	// 为memo赋值
	(*memo)[start] = sumIncome
	return sumIncome
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
