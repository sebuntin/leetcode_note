package main

/*
所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。

示例：

输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]
*/


func findRepeatedDnaSequences(s string) []string {
	wordMap := make(map[string]int, 1024)
	word := make([]byte, 0)
	result := make([]string, 0, 1024)
	for _, char := range []byte(s) {
		if len(word) == 10 {
			wordMap[string(word)] += 1
			word = append(word[:len(word)-1], word[1:]...)
			word = append(word, char)
		} else {
			word = append(word, char)
		}
	}
	for subStr, count := range wordMap {
		if count > 1 {
			result = append(result, subStr)
		}
	}
	return result
}