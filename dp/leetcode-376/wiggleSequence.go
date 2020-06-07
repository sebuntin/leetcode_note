package leetcode_376

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums))
	dp[0] = []int{1, -1}
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i][0] = 1
		for j := i - 1; j >= 0; j-- {
			if dp[j][0] < 2 && nums[j] != nums[i] && dp[i][0] < dp[j][0]+1 {
				dp[i][0] = dp[j][0] + 1
				dp[i][1] = nums[j]
				continue
			}
			if nums[j] < nums[i] && nums[j] < dp[j][1] && dp[i][0] < dp[j][0]+1 {
				dp[i][0] = dp[j][0] + 1
				dp[i][1] = nums[j]
			} else if nums[j] > nums[i] && nums[j] > dp[j][1] && dp[i][0] < dp[j][0]+1 {
				dp[i][0] = dp[j][0] + 1
				dp[i][1] = nums[j]
			}
		}
		if dp[i][1] > maxLen {
			maxLen = dp[i][1]
		}
	}
	return maxLen
}
