package main

/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。

示例:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8

说明:

    数组仅可以在 update 函数下进行修改。
    你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。

来源：https://leetcode-cn.com/problems/range-sum-query-mutable/
*/

import "fmt"

type NumArray struct {
	Value []int // 线段数组
	Len   int
}


func Constructor(nums []int) NumArray {
	value := make([]int, 4*len(nums))
	build_segment_tree(nums, value, 0, len(nums)-1, 0)
	return NumArray{value, len(nums)}
}

func (this *NumArray) SumRange(i int, j int) int {
	return sum_range_segment_tree(this.Value, 0, 0, this.Len-1, i, j)
}

func (this *NumArray) Update(i int, val int) {
	update_segment_tree(this.Value, val, i, 0, 0, this.Len-1)

}

func build_segment_tree(nums, value []int,left, right, pos int) {
	if len(nums) == 0 {
		return
	}
	if left == right {
		value[pos] = nums[left]
		return
	}
	mid := (left + right) / 2
	build_segment_tree(nums, value, left, mid, 2*pos+1)
	build_segment_tree(nums, value,mid+1, right, 2*pos+2)
	value[pos] = value[2*pos+1] + value[2*pos+2]
}

func sum_range_segment_tree(value []int, pos, left,
	right, qleft, qright int) int {

	if len(value) == 0 {
		return 0
	}
	if left > qright || right < qleft {
		return 0
	}
	if left >= qleft && right <= qright {
		return value[pos]
	}
	mid := (left + right) / 2
	return sum_range_segment_tree(value, 2*pos+1, left, mid, qleft, qright) +
		sum_range_segment_tree(value, 2*pos+2, mid+1, right, qleft, qright)
}

func update_segment_tree(value []int, val, index, pos, left, right int) {
	if len(value) == 0 {
		return
	}
	if left == right && left == index {
		value[pos] = val
		return
	}
	mid := (left + right) / 2
	if index <= mid {
		update_segment_tree(value, val, index,2*pos+1, left, mid)
	} else {
		update_segment_tree(value, val, index,2*pos+2, mid+1, right)
	}
	value[pos] = value[2*pos+1] + value[2*pos+2]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	sTree := Constructor(nums)
	fmt.Println(sTree)
	fmt.Println(sTree.SumRange(2, 4))
	sTree.Update(2, 3)
	fmt.Println(sTree.SumRange(2, 4))

}
