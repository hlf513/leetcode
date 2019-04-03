package main

import (
	"leetcode/Solutions/Go/Common"
)

func main() {
	tail := Common.ListNode{5, nil}
	next3 := Common.ListNode{4, &tail}
	next2 := Common.ListNode{3, &next3}
	next1 := Common.ListNode{2, &next2}
	head := Common.ListNode{1, &next1}

	res := reverseKGroup(&head, 3)
	Common.PrintListNode(*res)
}

// 根据链表长度计算交换次数
func reverseKGroup(head *Common.ListNode, k int) *Common.ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := Common.ListNode{-1, head}
	pre := &dummy
	cur := pre

	// 链表长度
	num := 0
	for cur.Next != nil {
		num++
		cur = cur.Next
	}

	// 交换多次
	for num >= k {
		cur = pre.Next
		for i := 1; i < k; i++ {
			tmp := cur.Next // 缓存
			cur.Next = tmp.Next
			tmp.Next = pre.Next // 必须用 pre.Next
			pre.Next = tmp      // 必须用 tmp
		}

		pre = cur
		num = num - k
	}

	return dummy.Next
}
