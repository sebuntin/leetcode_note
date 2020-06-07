package leetcode_337

//TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	rootIncome := robHelper(root)
	return Max(rootIncome[0], rootIncome[1])
}

// 返回当前节点下的两种状态下所能偷到的最大值，
// nodeIncome[0]表示不偷当前节点得到的最大值，
// nodeIncome[1]表示偷当前节点得到的最大值。
func robHelper(node *TreeNode) []int {
	nodeIncome := make([]int, 2)
	if node == nil {
		return nodeIncome
	}
	left := robHelper(node.Left)
	right := robHelper(node.Right)
	nodeIncome[0] = Max(left[0], left[1]) + Max(right[0], right[1])
	nodeIncome[1] = left[0] + right[0] + node.Val
	return nodeIncome
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
