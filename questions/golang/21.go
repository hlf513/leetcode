package main

import (
	"leetcode/common"
)

func main() {
	link13 := common.ListNode{4, nil}
	link12 := common.ListNode{2, &link13}
	link11 := common.ListNode{1, &link12}

	link23 := common.ListNode{4, nil}
	link22 := common.ListNode{3, &link23}
	link21 := common.ListNode{1, &link22}

	//common.PrintListNode(link11)
	//common.PrintListNode(link21)

	dummy := mergeTwoLists(&link11, &link21)
	common.PrintListNode(*dummy)
}

func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	dummy := common.ListNode{}
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
