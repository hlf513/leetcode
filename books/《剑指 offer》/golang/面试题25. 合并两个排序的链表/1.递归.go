package main

import "leetcode/common"

func main() {
	l1 := common.CreateListNode([]int{3, 5})
	l2 := common.CreateListNode([]int{1, 3})
	l := mergeTwoLists(l1, l2)

	common.PrintListNode(l)
}

/**
解题思路：递归，每次比较两个链表的值，每次连接最小的节点，直到遍历到其中一个的尾节点为止
时间复杂度：m+n
空间复杂度：m+n
*/
func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	// 退出条件
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		// 下沉，连接 l2
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		// 下沉，连接 l1
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}
