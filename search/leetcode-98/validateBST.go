package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// return DFS(root, ^int(^uint(0) >> 1), int(^uint(0) >> 1))
	return BFS(root)

}

// 深搜
func DFS(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	if !DFS(node.Left, lower, node.Val) {
		return false
	}
	if !DFS(node.Right, node.Val, upper) {
		return false
	}
	return true
}

// 宽搜
func BFS(node *TreeNode) bool {
	if node == nil {
		return true
	}
	type searchNode struct {
		Node         *TreeNode
		Upper, Lower int
	}
	// 定义搜索队列
	searchQ := make([]searchNode, 0)
	// 将根节点点放入搜索队列中
	searchQ = append(searchQ, searchNode{
		Node:  node,
		Upper: int(^uint(0) >> 1),
		Lower: ^int(^uint(0) >> 1)})
	for len(searchQ) != 0 {
		curNode := searchQ[0]
		searchQ = searchQ[1:]
		if curNode.Node.Val <= curNode.Lower || curNode.Node.Val >= curNode.Upper {
			return false
		}
		if curNode.Node.Left != nil {
			searchQ = append(searchQ, searchNode{
				Node:  curNode.Node.Left,
				Upper: curNode.Node.Val,
				Lower: curNode.Lower,
			})
		}
		if curNode.Node.Right != nil {
			searchQ = append(searchQ, searchNode{
				Node:  curNode.Node.Right,
				Upper: curNode.Upper,
				Lower: curNode.Node.Val,
			})
		}
	}
	return true
}

func main() {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)
}
