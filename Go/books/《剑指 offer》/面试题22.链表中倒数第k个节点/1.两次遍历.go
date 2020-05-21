package main

import (
	"leetcode/common"
)

func main() {
	head := common.CreateListNode([]int{1, 2, 3, 4, 5})

	common.PrintListNode(getKthFromEnd(head, 2))
}

/**
解题思路： 两次遍历，做减法，计算从前往后是第几个（n - k）
时间复杂度：2N
空间复杂度：1
*/
func getKthFromEnd(head *common.ListNode, k int) *common.ListNode {
	// 先遍历到最后一个节点，并记录长度
	var n int
	var cur = head 
	for cur != nil {
		n++
		cur = cur.Next
	}

	// 从前往后遍历，返回 n-k 
	for i := 1; i <= n-k; i++ {
		head = head.Next
	}

	return head
}
