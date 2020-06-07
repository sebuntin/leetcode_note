package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (this NodeHeap) Len() int {
	return len(this)
}

func (this NodeHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this NodeHeap) Less(i, j int) bool {
	return this[i].Val < this[j].Val
}

func (this *NodeHeap) Push(node interface{}) {
	*this = append(*this, node.(*ListNode))
}

func (this *NodeHeap) Pop() interface{} {
	x := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return x
}

func mergeKList(nodeList []*ListNode) *ListNode {
	return heapMerge(nodeList)
}

func heapMerge(nodeList []*ListNode) *ListNode {
	// 定义小顶堆
	nHeap := NodeHeap{}
	heap.Init(&nHeap)
	// 将所有链表的头节点入堆
	for i:=0; i<len(nodeList); i++ {
		heap.Push(&nHeap, nodeList[i])
	}
	// 声明一个哑节点作为新的头节点
	head := &ListNode{}
	// tail指向合并过程中新链表的最后一个节点
	tail := head
	for nHeap.Len() != 0 {
		// 每次从堆顶中取出节点作为tail的后序
		curNode := heap.Pop(&nHeap).(*ListNode)
		tail.Next = curNode
		if curNode.Next != nil {
			heap.Push(&nHeap, curNode.Next)
		}
		// tail向后移动
		tail = tail.Next
	}
	return head.Next
}
