package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	//removeHelper(&num, k)
	numStack := make([]uint8, 0, 1024)
	var result string
	for i := 0; i < len(num); i++ {
		x := num[i] - '0'
		if len(numStack) == 0 {
			//将第一个元素压入栈中
			numStack = append(numStack, x)
			continue
		} else {
			for k > 0 && len(numStack) != 0 && x < numStack[len(numStack)-1] {
				numStack = numStack[:len(numStack)-1]
				k--
			}
			if x != 0 || len(numStack) != 0 {
				numStack = append(numStack, x)
			}
		}
	}
	for k > 0 {
		numStack = numStack[:len(numStack)-1]
		k--
	}
	for _, x := range numStack {
		result += string(x + '0')
	}
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	var num string
	var k int
	_, _ = fmt.Scanf("%s", &num)
	_, _ = fmt.Scanf("%d", &k)
	fmt.Println(removeKdigits(num, k))
}
