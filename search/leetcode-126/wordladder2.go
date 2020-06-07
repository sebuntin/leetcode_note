/*
给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

    每次转换只能改变一个字母。
    转换过程中的中间单词必须是字典中的单词。

说明:

    如果不存在这样的转换序列，返回一个空列表。
    所有单词具有相同的长度。
    所有单词只由小写字母组成。
    字典中不存在重复的单词。
    你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-ladder-ii
*/

package main

type searchNode struct {
	Word         string
	PreLoc, Step int
}

func findLadders(beginWord string, endWord string, wordList []string) (ret [][]string) {
	// 定义搜索队列
	Q := make([]searchNode, 0, 1024)
	// 定义结果列表
	end_pos := make([]int, 0, 10)
	// 判断wordList中是否包含beginWord
	var flag bool
	for _, word := range wordList {
		if word == beginWord {
			flag = true
			break
		}
	}
	if !flag {
		wordList = append(wordList, beginWord)
	}
	// 初始化单词图的邻接表
	wordGraph := make(map[string][]string)
	buildGraph(wordList, &wordGraph)
	// 进行宽度优先搜索
	BFSGraph(beginWord, endWord, wordGraph, &Q, &end_pos)
	for _, index := range end_pos {
		preLoc := Q[index].PreLoc
		path := make([]string, 0)
		path = append(path, endWord)
		for preLoc != -1 {
			temp := make([]string, len(path)+1)
			temp[0] = Q[preLoc].Word
			copy(temp[1:], path)
			path = temp
			preLoc = Q[preLoc].PreLoc
		}
		ret = append(ret, path)
	}
	return
}

// 对单词图进行宽度优先搜索
func BFSGraph(beginWord, endWord string, graph map[string][]string, searchQ *[]searchNode, endPos *[]int) {
	visit := make(map[string]int)
	visit[beginWord] = 1
	// 初始化搜索位置
	loc := 0
	// 将beginWord入队
	*searchQ = append(*searchQ, searchNode{beginWord, -1, 1})
	var wordNode searchNode
	// 定义最短路径
	var minStep int
	for loc != len(*searchQ) {
		// 获取队首元素
		wordNode = (*searchQ)[loc]
		step := wordNode.Step
		word := wordNode.Word
		// 由于宽度优先搜索的特点，下一个待搜节点的肯定不会小于当前节点
		// 所以,如果此时的step已经比minStep大了, 那么后续的所有节点肯定不会包含在最短路径内
		if minStep != 0 && minStep < step {
			break
		}
		if word == endWord {
			minStep = step
			*endPos = append(*endPos, loc)
		}
		// 将邻节点push进入搜索队列
		for _, w := range graph[word] {
			if _, ok := visit[w]; (!ok || visit[w] == step+1) {
				visit[w] = step + 1
				*searchQ = append(*searchQ, searchNode{w, loc, step + 1})
			}
		}
		loc ++
	}
}

// 构建单词图的邻接表
func buildGraph(wordList []string, graph *map[string][]string) {
	// 默认wordList中已经包含了beginWord
	// 遍历wordList
	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if isNeighbor(wordList[i], wordList[j]) {
				(*graph)[wordList[i]] = append((*graph)[wordList[i]], wordList[j])
				(*graph)[wordList[j]] = append((*graph)[wordList[j]], wordList[i])
			}
		}
	}
}

// 两个单词是否相邻
func isNeighbor(w1, w2 string) bool {
	flag := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			flag++
		}
	}
	return flag == 1
}
