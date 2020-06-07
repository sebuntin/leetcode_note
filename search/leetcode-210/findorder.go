package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 建立邻接表
	classGraph := make(map[int][]int)
	preGraph := make(map[int][]int)
	for _, pair := range prerequisites {
		classGraph[pair[1]] = append(classGraph[pair[1]], pair[0])
		preGraph[pair[0]] = append(preGraph[pair[0]], pair[1])
	}
	// 定义搜索队列
	Q := make([]int, 0)
	// 定义标记数组
	visit := make([]int, numCourses)
	// 确定搜索队列的首次入队元素
	for i := 0; i < numCourses; i++ {
		isOK := true
		for _, pair := range prerequisites {
			if i == pair[0] {
				isOK = false
				break
			}
		}
		if isOK {
			// 入队
			Q = append(Q, i)
			visit[i] = 1
		}
	}
	BFS(visit, &Q, classGraph, preGraph)
	if isComplete(visit) {
		return Q
	} else {
		return []int{}
	}
}

func BFS(visit []int, Q *[]int, classGraph, preGraph map[int][]int) {
	loc := 0
	for loc != len(*Q) {
		cur := (*Q)[loc]
		for _, next := range classGraph[cur] {
			isOK := true
			for _, pre := range preGraph[next] {
				if visit[pre] == 0 {
					isOK = false
					break
				}
			}
			if visit[next] == 1 || !isOK {
				continue
			}
			visit[next] = 1
			*Q = append(*Q, next)
		}
		loc ++
	}
}

func isComplete(visit []int) bool {
	for _, status := range visit {
		if status == 0 {
			return false
		}
	}
	return true
}
