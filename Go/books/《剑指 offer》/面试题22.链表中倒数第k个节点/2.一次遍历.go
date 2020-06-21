package main

import (
	"leetcode/common"
)

/**
解题思路： 一次遍历，使用数组存储节点，直接返回下标 n-k 的节点
时间复杂度：N
空间复杂度：N
*/
func getKthFromEnd2(head *common.ListNode, k int) *common.ListNode {
	var a []*common.ListNode
	for head != nil {
		a = append(a, head)
		head = head.Next
	}

	l := len(a)
	if l >= k { // 防止越界
		return a[l-k]
	}

	return nil
}
