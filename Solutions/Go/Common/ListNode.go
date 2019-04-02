package Common

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(head ListNode) {
	for {
		fmt.Println(head.Val)
		if head.Next == nil {
			break
		}
		head = *head.Next
	}
}
