package main

import "fmt"

/*
给定一个整数数组 nums，按要求返回一个新数组 counts。
数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
*/

type BSTNode struct {
	Count, Val int
	Left       *BSTNode
	Right      *BSTNode
}

func bstInsert(node *BSTNode, insert *BSTNode, result []int, i int) {
	if node.Val >= insert.Val {
		node.Count++
		if node.Left != nil {
			bstInsert(node.Left, insert, result, i)
		} else {
			node.Left = insert
		}
	} else {
		result[i] += node.Count + 1
		if node.Right != nil {
			bstInsert(node.Right, insert, result, i)
		} else {
			node.Right = insert
		}
	}
}

func countSmaller(nums []int) []int {
	result := make([]int, len(nums))
	// 将nums中的数变为BSTNode
	root := &BSTNode{Count: 0, Val: nums[len(nums)-1]}
	// 构建二叉搜索树
	for i := len(nums) - 2; i > 0; i-- {
		insertNode := &BSTNode{Count: 0, Val: nums[i]}
		bstInsert(root, insertNode, result, i)
	}
	return result
}

func moveElem(a []int) {
	head := a[0]
	a = append(a[:0], a[1:]...)
	a = append(a, head)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	moveElem(a)
	fmt.Println(a)
}
