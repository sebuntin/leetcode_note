package leetcode_209

func minSubArrayLen(s int, nums []int) int {
	start := 0
	windowSum := 0
	minLen := len(nums)
	maxSum := 0
	for i := range nums {
		maxSum += nums[i]
	}
	if maxSum < s {
		return 0
	}
	for i := range nums {
		windowSum += nums[i]
		for windowSum >= s && start <= i {
			//记录当前长度
			if minLen > i-start+1 {
				minLen = i - start + 1
			}
			windowSum -= nums[start]
			start++
		}
	}
	return minLen
}
