package main

import (
	"fmt"
	"leetcode/common"
)

func main() {
	link14 := common.ListNode{-4, nil}
	link13 := common.ListNode{0, &link14}
	link12 := common.ListNode{2, &link13}
	link11 := common.ListNode{3, &link12}

	link14.Next = &link12

	fmt.Println(detectCycle(&link11).Val)
}

func detectCycle(head *common.ListNode) *common.ListNode {
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
