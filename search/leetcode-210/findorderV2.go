package main

import "fmt"

type vertext struct {
	in  int        //表示入度值
	val int        // 表示节点值
	adj []*vertext // 表示后继邻接点
}

func findOrderv2(numCourses int, prerequisites [][]int) []int {
	// 定义有向图图
	G := make([]*vertext, numCourses)
	for i := 0; i < numCourses; i++ {
		G[i] = &vertext{val: i}
	}
	for _, pair := range prerequisites {
		G[pair[1]].adj = append(G[pair[1]].adj, G[pair[0]])
		G[pair[0]].in++
	}
	ret := make([]int, 0)
	topologicalSort(G, &ret)
	if len(ret) < numCourses {
		return []int{}
	}
	return ret
}

// 拓扑排序
func topologicalSort(G []*vertext, ret *[]int) {
	S := make([]*vertext, 0)
	for i := 0; i < len(G); i++ {
		if G[i].in == 0 {
			S = append(S, G[i])
		}
	}
	for len(S) != 0 {
		v := S[0]
		S = S[1:]
		*ret = append(*ret, v.val)
		for i := range v.adj {
			v.adj[i].in--
			if v.adj[i].in == 0 {
				S = append(S, v.adj[i])
			}
		}
	}
}

func main() {
    numCouses := 4
    prerequisites := [][]int{{1,0}, {2,0}, {3,1}, {3,2}}
    ret := findOrderv2(numCouses,prerequisites)
    fmt.Println(ret)
}
