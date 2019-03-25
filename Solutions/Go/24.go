package main

import "fmt"

func main() {
	tail := ListNode{5, nil}
	next3 := ListNode{4, &tail}
	next2 := ListNode{3, &next3}
	next1 := ListNode{2, &next2}
	head := ListNode{1, &next1}

	printListNode(*swapPairs(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{}
	pre := &dummy
	pre.Next = head

	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		post := cur.Next

		pre.Next, post.Next, cur.Next = post, cur, post.Next

		pre = cur
	}

	return dummy.Next
}
