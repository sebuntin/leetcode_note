package main

import (
	"fmt"
	"math"
	"sort"
)

/*
还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。
输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。

示例 1:


输入: [1,1,2,2,2]
输出: true

解释: 能拼成一个边长为2的正方形，每边两根火柴。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/matchsticks-to-square
*/

type intSlice []int

func (this intSlice) Len() int           { return len(this) }
func (this intSlice) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this intSlice) Less(i, j int) bool { return this[i] < this[j] }

func makesquare(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%4 != 0 || len(nums) == 0 {
		return false
	}
	bucket := make([]int, 4)
	// 对nums进行从大到小排序
	newNums := intSlice(nums)
	sort.Sort(sort.Reverse(newNums))
	return putstick(0, sum/4, newNums, bucket)

}

func putstick(i, target int, nums, bucket []int) bool {
	if i >= len(nums) {
		return bucket[0] == bucket[1] && bucket[0] == bucket[2] && bucket[0] == target
	}
	// 放置第i根火柴
	for j := 0; j < 4; j++ {
		// 逐个尝试bucket的四个位置能否放置火柴
		if bucket[j]+nums[i] > target {
			continue
		}
		bucket[j] += nums[i]
		if putstick(i+1, target, nums, bucket) {
			return true
		}
		// 回溯
		bucket[j] -= nums[i]
	}
	return false
}

// 方法二：位运算
func makesquare2(nums []int) bool {
	// 检查输入的合法性
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%4 != 0 || len(nums) == 0 {
		return false
	}
	target := sum / 4
	// 定义一个子集集合，子集中的元素之和为sum%4
	subSet := make([]int, 0)
	// 定义一个子集集合，集合中每个子集对应square的两条边
	halfSubSet := make([]int, 0)
	// 对数组进行编码
	all := int(math.Pow(float64(2), float64(len(nums))))
	for i := 0; i < all; i++ {
		sum = 0
		for j := 0; j < len(nums); j++ {
			if i & (1<<(len(nums)-1-j)) != 0 {
				sum += nums[j]
			}
		}
		if sum == target {
			subSet = append(subSet, i)
		}
	}
	for i := 0; i < len(subSet); i++ {
		for j := i + 1; j < len(subSet); j++ {
			if subSet[i]&subSet[j] == 0 {
				halfSubSet = append(halfSubSet, subSet[i]|subSet[j])
			}
		}
	}

	for i := 0; i < len(halfSubSet); i++ {
		for j := i + 1; j < len(halfSubSet); j++ {
			if halfSubSet[i]&halfSubSet[j] == 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	//nums := intSlice{1, 3, 2, 4, 5, 8, 6}
	//sort.Sort(sort.Reverse(nums))
	//fmt.Println(nums)
	fmt.Println(makesquare2([]int{1, 1, 2, 2, 2}))
}
