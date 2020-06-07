package leetcode_78

func Subsets(nums []int) (ret [][]int) {
	all := 1<<len(nums)
	for i := 0; i < all; i++ {
		var subSet []int
		for j := 0; j < len(nums); j++ {
			if i & (1<<(len(nums)-1-j)) != 0 {
				subSet = append(subSet, nums[j])
			}
		}
		ret = append(ret, subSet)
	}
	return
}

