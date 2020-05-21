package main

import (
	"leetcode/common"
)

func main() {
	head := common.CreateListNode([]int{1, 2, 3, 4, 5})

	common.PrintListNode(getKthFromEnd3(head, 2))

}

/**
解题思路：快慢指针，快指针先移动 K 个；然后快慢指针一起移动，当快指针到尾节点时，当前的慢指针就是倒数第 K 个节点.
时间复杂度：N
空间复杂度：1
*/
func getKthFromEnd3(head *common.ListNode, k int) *common.ListNode {
	if head == nil {
		return head
	}

	// 快指针先移动 K
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// 快慢一起移动
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
