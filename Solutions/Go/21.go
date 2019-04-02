package main

import (
	"leetcode/Solutions/Go/Common"
)

func main() {
	link13 := Common.ListNode{4, nil}
	link12 := Common.ListNode{2, &link13}
	link11 := Common.ListNode{1, &link12}

	link23 := Common.ListNode{4, nil}
	link22 := Common.ListNode{3, &link23}
	link21 := Common.ListNode{1, &link22}

	//Common.PrintListNode(link11)
	//Common.PrintListNode(link21)

	dummy := mergeTwoLists(&link11, &link21)
	Common.PrintListNode(*dummy)
}

func mergeTwoLists(l1 *Common.ListNode, l2 *Common.ListNode) *Common.ListNode {
	dummy := Common.ListNode{}
	cur := &dummy

	// 循环比较链表的值
	for l1 != nil && l2 != nil { // 其中一个链表为空停止
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	// 处理剩余的链表
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
