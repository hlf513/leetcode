package main

import (
	"leetcode/common"
)

func main() {
}

/**
解题思路：利用哈希表，把其中一个链表的节点地址放入 hash 表中，遍历另一个链表，查找是否有相同地址的节点
时间复杂度：m+n
空间复杂度：m/n
*/
func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	var hash = make(map[*common.ListNode]bool)
	for headA != nil {
		hash[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hash[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
