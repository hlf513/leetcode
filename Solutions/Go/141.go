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

	fmt.Println(hasCycle(&link11))
}

func hasCycle(head *Common.ListNode) bool {
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
