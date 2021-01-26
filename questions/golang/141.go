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

	fmt.Println(hasCycle(&link11))
}

func hasCycle(head *common.ListNode) bool {
	slow, fast := head, head
	res := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			res = true
			break
		}
	}

	return res
}
