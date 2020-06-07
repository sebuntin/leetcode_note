package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 找到中位数在合并数组中的位置(left+right)/2
	n := len(nums1) + len(nums2)
	left, right := (n+1)/2, (n+2)/2
	return (float64(findHelper(nums1, nums2, left)) +
		 float64(findHelper(nums1, nums2, right))) / 2
}

// 从nums1和nums2中找到第k小的数
func findHelper(nums1 []int, nums2 []int, k int) int {
	if k == 1 {
		if len(nums1) == 0 {
			return nums2[k-1]
		} else if len(nums2) == 0 {
			return nums1[k-1]
		} else {
			return Min(nums1[k-1], nums2[k-1])
		}
	}
	var p1 int
	var p2 int
	if k/2 >= len(nums1) {
		p1 = len(nums1)
		p2 = p1
		k = k - 2*p1
		return findHelper(nums1[p1:], nums2[p2:], k)
	}
	if k/2 >= len(nums2) {
		p2 = len(nums2)
		p1 = p2
		k = k - 2*p2
		return findHelper(nums1[p1:], nums2[p2:], k)
	}
	if nums1[k/2-1] < nums2[k/2-1] {
		p1 = k / 2
		p2 = k/2 - 1
		k = k - (k/2) * 2 + 1
	} else {
		p2 = k / 2
		p1 = k/2 - 1
		k = k - (k/2) * 2 + 1
	}
	return findHelper(nums1[p1:], nums2[p2:], k)
}

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	nums1 := []int{1, 3, 5, 7}
	nums2 := []int{2, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
