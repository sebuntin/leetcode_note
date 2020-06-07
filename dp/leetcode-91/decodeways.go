package leetcode_91

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		// 前一个字符为0时
		switch s[i-1] {
		case '0':
			if s[i] == '0' {
				return 0
			}
			dp[i] = dp[i-1]
			continue
		}
		if s[i] == '0' {
			if s[i-1]-'0' > 2 {
				return 0
			}
			if i-1 == 0 {
				dp[i] = dp[i-1]
			} else {
				dp[i] = dp[i-2]
			}
			continue
		}
		if s[i-1]-'0' < 2 || (s[i-1]-'0' == 2 && s[i]-'0' <= 6) {
			if i >= 2 {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1] + 1
			}
			continue
		}
		dp[i] = dp[i-1]
	}
	return dp[len(s)-1]
}
