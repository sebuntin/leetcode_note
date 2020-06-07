package main

import "fmt"

// 归并排序
func mergeTwoSort(nums1, nums2 []int , nums []int) {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	for ; i < len(nums1); i++ {
		nums = append(nums, nums1[i])
	}
	for ; j < len(nums2); j++ {
		nums = append(nums, nums2[j])
	}
}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	nums1 := make([]int, mid)
	nums2 := make([]int, len(nums)-mid)
	copy(nums1, nums[:mid])
	copy(nums2, nums[mid:])
	sort(nums1)
	sort(nums2)
	nums = nums[0:0]
	mergeTwoSort(nums1, nums2, nums)
}

func main() {
	nums := []int{1, 3, 2, 6, 5, 8, 4, 7}
	sort(nums)
	fmt.Println(nums)
	nums = nums[0:0]
	nums1 := nums[1:4]
	fmt.Println(nums1)
}
