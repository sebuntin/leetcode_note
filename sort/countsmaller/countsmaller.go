package main

import "fmt"

type sortNode struct {
	val   int
	index int
}

func countSmaller(nums []int) []int {
	numsSort := make([]sortNode, len(nums))
	count := make([]int, len(nums))
	for i, num := range nums {
		numsSort[i] = sortNode{num, i}
	}
	mergeSort(numsSort, count)
	return count
}

func mergeSort(nums []sortNode, count []int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	leftNums := make([]sortNode, mid)
	RightNums := make([]sortNode, len(nums)-mid)
	copy(leftNums, nums[:mid])
	copy(RightNums, nums[mid:])
	mergeSort(leftNums, count)
	mergeSort(RightNums, count)
	nums = nums[0:0]
	mergeNums(leftNums, RightNums, nums, count)
}

func mergeNums(nums1, nums2, nums []sortNode, count []int) {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i].val <= nums2[j].val {
			count[nums1[i].index] += j
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	for ; i < len(nums1); i++ {
		count[nums1[i].index] += j
		nums = append(nums, nums1[i])
	}
	for ; j<len(nums2); j++ {
		nums = append(nums, nums2[j])
	}
}

func main() {
    nums := []int{1, 2, 6, 3, -1, 7, 4}
    fmt.Println(countSmaller(nums))
    s := "hello"
    fmt.Println(s[1:3])
}
