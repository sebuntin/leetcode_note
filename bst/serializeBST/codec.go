package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	data := ""
	preorder(root, &data)
	return data
}

func preorder(root *TreeNode, strVal *string) {
	// 前序遍历二叉树
	if root == nil {
		return
	}
	*strVal += strconv.Itoa(root.Val)
	*strVal += "#"
	preorder(root.Left, strVal)
	preorder(root.Right, strVal)
}

func bstInsert(node *TreeNode, insert *TreeNode) {
	if insert.Val < node.Val {
		if node.Left != nil {
			bstInsert(node.Left, insert)
		} else {
			node.Left = insert
		}
	} else {
		if node.Right != nil {
			bstInsert(node.Right, insert)
		} else {
			node.Right = insert
		}
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strSlice := strings.Split(data, "#")
	strSlice = strSlice[:len(strSlice)-1]
	nodeStack := make([]*TreeNode, 0, 0124)
	for _, str := range strSlice {
		val, _ := strconv.Atoi(str)
		node := TreeNode{Val: val}
		nodeStack = append(nodeStack, &node)
	}
	root := nodeStack[0]
	for _, node := range nodeStack {
		if node == root {
			continue
		}
		bstInsert(root, node)
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
func main() {
    a := "1#2#3#"
    a_slice := strings.Split(a, "#")
    for _, each := range a_slice {
    	fmt.Println(strconv.Atoi(each))
    }
}