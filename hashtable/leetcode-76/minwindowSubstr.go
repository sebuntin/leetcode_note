package main

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minWindow(s string, t string) string {
	start := 0
	minLen := len(s)
	minSubStr := ""
	// 定义字符哈希表
	t_map := make([]int, 126)
	s_map := make([]int, 126)
	for _, char := range []byte(t) {
		t_map[int(char)] ++
	}
	for i, char := range []byte(s) {
		s_map[int(char)] ++
		for start < i {
			if t_map[int(s[start])] == 0 {
				s_map[int(s[start])] --
				start++
			} else if s_map[int(s[start])] > t_map[int(s[start])] {
				s_map[int(s[start])] --
				start++
			} else {
				break
			}
		}
		if isValidWiondw(s_map, t_map) && minLen >= i-start+1 {
			minLen = i - start + 1
			minSubStr = string([]byte(s)[start : i+1])
		}
	}
	return minSubStr
}

// 检查当前窗口是否包含t中的字符
func isValidWiondw(s_map, t_map []int) bool {
	for i := 0; i < 126; i++ {
		if s_map[i] < t_map[i] {
			return false
		}
	}
	return true
}
