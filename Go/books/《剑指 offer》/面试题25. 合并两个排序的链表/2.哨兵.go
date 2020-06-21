package main

import "leetcode/common"

/**
解题思路： 利用哨兵，遍历两个链表；dummy 用于返回头节点，pre 用于构建新链表
时间复杂度：m+n
空间复杂度：1
*/
func mergeTwoLists2(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	// 声明一个哨兵节点，用于存储头节点
	var dummy = new(common.ListNode)
	// 声明一个 pre 节点，用于遍历时暂存前一个节点
	pre := dummy

	for l1 != nil && l2 != nil { // 两个链表有一个为空则退出
		if l1.Val > l2.Val {
			// 指向小的节点
			pre.Next = l2
			// l2 节点向后移位
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}
		// pre 同步向后移位
		pre = pre.Next
	}
	// 判断是否有非空链表，若有则直接指向
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}

	return dummy.Next
}
