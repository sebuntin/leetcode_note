package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	prefixSum := make(map[int]int)
	sum := 0
	count := 0
	prefixSum[0] = 1
	for i := range A {
		sum += A[i]
		mod := (sum%K + K) % K
		if k, ok := prefixSum[mod]; ok {
			count += k
		}
		prefixSum[mod] ++
	}
	return count
}

func main() {
	fmt.Println(-3 % 5)
}
