package main

import "sort"

type value [][]int

func (this value) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this value) Len() int {
	return len(this)
}

func (this value) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func findMinArrowShots(points value) (counts int) {
	if len(points) == 0 {
		return 0
	}
	counts = 1
	// 对所有气球按照左端点进行排序
	sort.Sort(points)
	// 初始化射击区间为第一个气球
	interval := points[0]
	// 遍历气球，并维护一个射击区间
	for _, each := range points {
		// 若当前的射击区间与当前遍历的气球位置没有重叠时，增加弓箭手数量
		// 同时更新射击区间
		if each[0] > interval[1] {
			counts++
			interval = each
		} else {
			// 更新射击区间
			interval[0] = each[0]
			if interval[1] > each[1] {
				interval[1] = each[1]
			}
		}
	}
	return
}
