package leetcode_53

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp := make([]int, 0)
	//dp = append(dp, nums[0])
	//maxSum := dp[0]
	maxSum := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev < 0 {
			//dp = append(dp, nums[i])
			prev = nums[i]
		} else {
			//dp = append(dp, nums[i]+dp[i-1])
			prev += nums[i]
		}
		if maxSum < prev {
			maxSum = prev
		}
	}
	return maxSum
}


