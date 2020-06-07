package main

import "fmt"

var letterMap = map[byte][]byte{
	'2': []byte("abc"),
	'3': []byte("def"),
	'4': []byte("ghi"),
	'5': []byte("jkl"),
	'6': []byte("mno"),
	'7': []byte("pqrs"),
	'8': []byte("tuv"),
	'9': []byte("wxyz"),
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	findCombination([]byte(digits), "", &result)
	return result
}

func findCombination(digits []byte, s string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, s)
		return
	}
	for _, b := range letterMap[digits[0]] {
		findCombination(digits[1:], s+string(b), result)
	}
}

func main() {
    fmt.Println(letterCombinations("234"))
}
