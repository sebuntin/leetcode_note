package leetcode_78

import "sort"

func subsetsWithDup(nums []int) (ret [][]int){
	// 对输入的数组进行排序
	sort.Ints(nums)
	// 定义一个去重map
	dupMap := make(map[string]int)
	// 得到所有的子集
	all := 1 << len(nums)
	for i := 0; i < all; i++ {
		subSetByte := make([]byte, 0)
		subSet := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i & (1<<(len(nums)-1-j)) != 0 {
				subSetByte = append(subSetByte, byte(nums[j]))
				subSet = append(subSet, nums[j])
			}
		}
		if _ , ok:= dupMap[string(subSetByte)]; ok {
			continue
		} else {
			dupMap[string(subSetByte)] ++
			ret = append(ret, subSet)
		}
	}
	return ret
}
