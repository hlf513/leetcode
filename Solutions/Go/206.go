package main

import "fmt"

func main() {

	tail := ListNode{5, nil}
	next3 := ListNode{4, &tail}
	next2 := ListNode{3, &next3}
	next1 := ListNode{2, &next2}
	head := ListNode{1, &next1}

	printListNode(*reverseList(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//  迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}

	return pre
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var node = reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

func printListNode(head ListNode) {
	for {
		fmt.Println(head.Val)
		if head.Next == nil {
			break
		}
		head = *head.Next
	}
}
