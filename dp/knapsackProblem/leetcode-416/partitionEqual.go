package leetcode_416

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 || len(nums) == 0 {
		return false
	}
	partSum := sum / 2
	dp := make([]int, partSum+1)
	dp[0] = 1
	for i := 0; i <= partSum; i++ {
		if nums[0] == i {
			dp[i] = 1
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := partSum; j >= nums[i]; j++ {
			dp[j] = dp[j] | dp[j-nums[i]]
		}
		if dp[partSum] == 1 {
			return true
		}
	}
	return false
}
