package main

import "fmt"

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
*/
//Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next:head}
	pre := hair
	for head != nil {
		tail := pre
		for i:=0; i<k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = reverseHelper(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func f2() (y int) {
	x := 5
	defer func() {
		x++
		fmt.Println(x)
	}()
	return x // 返回值为 5
}


func reverseHelper(head, tail *ListNode) (*ListNode, *ListNode){
	prev := tail.Next
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return tail, head
}

func main() {
    head := &ListNode{Val:  0, Next: nil, }
    n1 := &ListNode{ Val:  1, Next: nil, }
    n2 := &ListNode{Val:2, Next:nil}
    head.Next = n1
    n1.Next = n2
    cur := head
    cur.Next = n2
    cur = cur.Next
    cur.Next = n1
    n1.Next = nil
    for h:=head; h!=nil; h = h.Next {
    	fmt.Println(h.Next)
    }
    fmt.Println(f2())
}