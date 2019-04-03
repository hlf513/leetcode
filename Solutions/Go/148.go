package main

import (
	"leetcode/Solutions/Go/Common"
)

func main() {
	link13 := Common.ListNode{3, nil}
	link12 := Common.ListNode{4, &link13}
	link11 := Common.ListNode{2, &link12}
	link10 := Common.ListNode{5, &link11}

	res := sortList(&link10)
	Common.PrintListNode(*res)
}

// 归并排序
func sortList(head *Common.ListNode) *Common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		pre = slow
		// 移动指针
		slow = slow.Next      // 走一步
		fast = fast.Next.Next // 走两步
	}
	pre.Next = nil // 分割链表（很重要）

	return merger(sortList(head), sortList(slow))
}

// 合并两个有序链表
func merger(l1 *Common.ListNode, l2 *Common.ListNode) *Common.ListNode {
	d := new(Common.ListNode)
	cur := d

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next // 移动 l2 指针
		} else {
			cur.Next = l1
			l1 = l1.Next // 移动 l1 指针
		}

		cur = cur.Next // 移动 cur 指针
	}

	// 链表不等长，处理剩余节点
	if l1 == nil {
		cur.Next = l2
	} else if l2 == nil {
		cur.Next = l1
	}

	return d.Next
}
