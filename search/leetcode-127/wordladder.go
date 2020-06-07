/*
词语阶梯1
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

    每次转换只能改变一个字母。
    转换过程中的中间单词必须是字典中的单词。

说明:

    如果不存在这样的转换序列，返回 0。
    所有单词具有相同的长度。
    所有单词只由小写字母组成。
    字典中不存在重复的单词。
    你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
*/

package main

import "fmt"


type searchNode struct {
	Word string
	Step int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 将beginWord加入wordList中
	wordList = append(wordList, beginWord)
	// 构建单词图
	graph := make(map[string][]string, len(wordList))
	// 为graph分配空间
	for _, word := range wordList {
		wordVec := make([]string, 0, len(wordList))
		graph[word] = wordVec
	}
	// 根据wordList构建单词图
	buildGraph(wordList, graph)
	return BFSGraph(beginWord, endWord, graph)
}


func buildGraph(wordList []string, graph map[string][]string) {
	// 先将beginWord添加进wordList中
	// 遍历wordList
	for _, word := range wordList {
		for _, w := range wordList {
			if isConnect(word, w) {
				graph[word] = append(graph[word], w)
			}
		}
	}
}

func BFSGraph(beginWord, endWord string, graph map[string][]string) (step int) {
	// 定义搜索队列
	searchQ := make([]searchNode, 0, 1024)
	// 定义访问队列
	visit := make(map[string]int)
	// 将beginWord入搜索队列
	searchQ = append(searchQ, searchNode{beginWord, 1})
	visit[beginWord] ++
	var curNode searchNode
	for len(searchQ) != 0 {
		curNode = searchQ[0]
		searchQ = searchQ[1:]
		if curNode.Word == endWord {
			return curNode.Step
		}
		for _, w := range graph[curNode.Word] {
			if _, ok := visit[w]; !ok {
				visit[w] ++
				searchQ = append(searchQ, searchNode{w, curNode.Step + 1})
			}
		}
	}
	return
}
// 判断两个word是否能相连
func isConnect(w1, w2 string) bool {
	flag := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			flag++
		}
	}
	return flag == 1
}

func main() {
	wordMap := make(map[string]int)
	wordMap["hello"] ++
	fmt.Println(wordMap)
}
