package main

import (
	"bufio"
	"fmt"
	"os"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (p *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if (l1.Val <= l2.Val) {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func createList(line []byte) *ListNode {
	head := new(ListNode)
	cur := head
	for i:=0;i<len(line); i++ {
		node := &ListNode{Val:int(line[i]-'0')}
		cur.Next = node
		cur = cur.Next
	}
	l1 := head.Next
	return l1
}

func main() {
	line1, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	line2 ,_, _ :=bufio.NewReader(os.Stdin).ReadLine()

	l1 := createList(line1)
	l2 := createList(line2)

	l := mergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}



