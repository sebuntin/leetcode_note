package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// hashMap方法
func hasCycle(head *Node) bool {
	nodeMap := make(map[*Node]int, 10)
	for ; head != nil; head = head.Next {
		if nodeMap[head] == 1 {
			return true
		}
		nodeMap[head] = 1
	}
	return false
}

// 快慢指针法
func hasCyclePro(head *Node) bool {
	fastPtr := head
	slowPtr := head
	for fastPtr != nil {
		// 快指针走两步
		for i := 0; i < 2; i++ {
			if fastPtr == nil {
				return false
			}
			fastPtr = fastPtr.Next
		}
		// 慢指针走一步
		slowPtr = slowPtr.Next
		// 相遇则说明有环
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}

func main() {
	a := &Node{1, nil}
	b := &Node{3, nil}
	c := &Node{5, nil}
	d := &Node{6, nil}
	f := &Node{20, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = f
	f.Next = c

	fmt.Println(hasCycle(a))

}
