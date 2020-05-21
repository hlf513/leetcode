package main

import "leetcode/common"

func main() {
	head := common.CreateListNode([]int{1, 2, 3, 4})
	reverseHead := reverseList(head)

	common.PrintListNode(reverseHead)

}

/**
解题思路：利用哨兵节点，每次把 cur.next -> p , p = cur , cur = cur.Next
时间复杂度：N
空间复杂度：1
 */
func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	// 声明一个哨兵节点，用于存储前一个节点
	var pre *common.ListNode
	// cur 表示当前节点
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}
