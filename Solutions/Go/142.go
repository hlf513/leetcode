package main

import (
	"fmt"
	"leetcode/Solutions/Go/Common"
)

func main() {
	link14 := Common.ListNode{-4, nil}
	link13 := Common.ListNode{0, &link14}
	link12 := Common.ListNode{2, &link13}
	link11 := Common.ListNode{3, &link12}

	link14.Next = &link12

	fmt.Println(detectCycle(&link11).Val)
}

func detectCycle(head *Common.ListNode) *Common.ListNode {
	slow, fast := head, head

	// 第一次相遇
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	// 各自走一步
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
