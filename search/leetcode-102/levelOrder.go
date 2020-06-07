package leetcode_102

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type TreeNodeWithD struct {
	Node *TreeNode
	Depth int
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	BFS(root, &ret)
	return ret
}

func BFS(node *TreeNode, treeArr *[][]int) {
	// 定义搜索队列
	searchQ := make([]*TreeNodeWithD, 0)
	searchQ = append(searchQ, &TreeNodeWithD{Node:node, Depth:0})
	for len(searchQ) != 0 {
		cur:= searchQ[0].Node
		depth := searchQ[0].Depth
		searchQ = searchQ[1:]
		if len(*treeArr) < depth+1 {
			*treeArr = append(*treeArr, []int{cur.Val})
		} else {
			(*treeArr)[depth] = append((*treeArr)[depth], cur.Val)
		}
		if cur.Left != nil {
			searchQ = append(searchQ, &TreeNodeWithD{Node:cur.Left, Depth:depth+1})
		}
		if cur.Right != nil {
			searchQ = append(searchQ, &TreeNodeWithD{Node:cur.Right, Depth:depth+1})
		}
	}
}