package main

import "leetcode/common"

func main() {

	listHead := common.CreateListNode([]int{4, 5, 1, 9})
	newListHead := deleteNode(listHead, 5)

	common.PrintListNode(newListHead)
}

/**
三种情况：删除头节点、删除中间节点、删除尾节点
解题思路：使用哨兵输出头节点，然后遍历找到对应节点删除
时间复杂度：N
空间复杂度：1
*/
func deleteNode(head *common.ListNode, val int) *common.ListNode {
	// 哨兵节点，用于输出头节点
	var dummy = &common.ListNode{
		Next: head,
	}
	// 用于存储前一个节点
	preNode := dummy
	for preNode.Next != nil {
		if preNode.Next.Val == val {
			preNode.Next = preNode.Next.Next
			break
		}
		preNode = preNode.Next
	}

	return dummy.Next
}
