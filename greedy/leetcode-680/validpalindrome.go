package main

func validPalindrome(s string) bool {
	//判断原字符串是否为回文串
	if isPalindrome(s) {
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	r := 0
	l := len(s) - 1
	for r < l {
		if s[r] != s[l] {
			return false
		}
		r++
		l--
	}
	return true
}
