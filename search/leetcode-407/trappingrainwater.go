package main

import (
	"container/heap"
)

type HeapNode struct {
	H, X, Y int
}

type HeightHeap []HeapNode

// 小顶堆
func (this HeightHeap) Less(i, j int) bool {
	return this[i].H < this[j].H
}

func (this HeightHeap) Len() int {
	return len(this)
}

func (this HeightHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *HeightHeap) Pop() interface{} {
	x := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return x
}

func (this *HeightHeap) Push(x interface{}) {
	*this = append(*this, x.(HeapNode))
}

func trapRainWater(heightMap [][]int) (ret int) {
	// 得到heightMap的形状
	row := len(heightMap)
	col := len(heightMap[0])
	if row < 3 || col < 3 {
		return 0
	}
	// 初始化标记数组，将入堆的元素标记为以访问
	mark := make([][]int, 0, len(heightMap))
	for i := 0; i < row; i++ {
		placeHolder := make([]int, col)
		mark = append(mark, placeHolder)
	}
	// 初始化一个堆
	heightHeap := &HeightHeap{}
	heap.Init(heightHeap)
	// 先将heightMap中的四周元素进入堆, 并将标记数组对应位置置1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				heap.Push(heightHeap, HeapNode{heightMap[i][j], i, j})
				mark[i][j] = 1
			}
		}
	}
	// 方向数组，表示上下左右
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	// 只要堆不空
	// 带堆的宽度搜索
	for heightHeap.Len() != 0 {
		node := heap.Pop(heightHeap).(HeapNode)
		// 判断其上下左右是否可以入堆
		for i := 0; i < 4; i++ {
			newX := node.X + dx[i]
			newY := node.Y + dy[i]
			// 边界判定和, 是否能够入堆
			if newX < 0 || newY < 0 || newX >= row || newY >= col ||
				mark[newX][newY] != 0 {
				continue
			}
			newH := heightMap[newX][newY]
			if newH < node.H {
				ret += node.H - newH
				newH = node.H
			}
			mark[newX][newY] = 1
			heap.Push(heightHeap, HeapNode{heightMap[newX][newY], newX, newY})
		}
	}
	return
}
