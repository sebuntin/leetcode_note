package leetcode_128

import (
	"fmt"
	"testing"
)

func TestDisJoinSet_Union(t *testing.T) {
	nums := []int{1, 3, 4, 5, 2}
	maxLen := longestConsecutive(nums)
	fmt.Println(maxLen)

}
