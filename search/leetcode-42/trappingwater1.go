package main

import (
	"container/heap"
	"fmt"
)

type heapNode struct {
	Index, H int
}

type heightHeap []heapNode

func (this heightHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this heightHeap) Len() int {
	return len(this)
}

func (this heightHeap) Less(i, j int) bool {
	return this[i].H < this[j].H
}

func (this *heightHeap) Push(node interface{}) {
	*this = append(*this, node.(heapNode))
}

func (this *heightHeap) Pop() interface{} {
	node := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return node
}

func trap(height []int) int {
	res := 0
	BFS(height, &res)
	return res
}

func BFS(height []int, res *int) {
	// 创建一个最小堆
	hHeap := &heightHeap{}
	heap.Init(hHeap)
	// 创建标记数组
	mark := make([]int, len(height))
	mark[0] = 1
	mark[len(height)-1] = 1
	// 将height首尾push入堆中
	heap.Push(hHeap, heapNode{Index:0, H:height[0]})
	heap.Push(hHeap, heapNode{Index:len(height)-1, H:height[len(height)-1]})
	// 广度搜索
	dx := []int{1, -1}
	for hHeap.Len() != 0 {
		curNode := heap.Pop(hHeap).(heapNode)
		for i:=0; i<2; i++ {
			newIndex := curNode.Index + dx[i]
			if newIndex < 0 || newIndex >= len(height) || mark[newIndex] == 1 {
				continue
			}
			newH := height[newIndex]
			if newH < curNode.H {
				*res += curNode.H - newH
				newH = curNode.H   
			}
			heap.Push(hHeap, heapNode{Index:newIndex, H:newH})
			mark[newIndex] = 1
		}
	}
}

func main() {
    height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
    fmt.Println(trap(height))
}
