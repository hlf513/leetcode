package main

import "leetcode/common"

func main() {
}

/**
解题思路：两个指针每次移位一个，若到自己链表的尾节点，则重定向到另一个链表头，直到两个指针相遇(需记录各自队尾的节点，若两个尾节点不相同，则说明无相交节点)
时间复杂度：m+n
空间复杂度：1
*/
func getIntersectionNode2(headA, headB *common.ListNode) *common.ListNode {
	var headATailNode, headBTailNode *common.ListNode
	headAHead := headA
	headBHead := headB

	for headA != headB {
		if headA != nil {
			headA = headA.Next
		} else {
			headATailNode = headA
			headA = headBHead
		}

		if headB != nil {
			headB = headB.Next
		} else {
			headBTailNode = headB
			headB = headAHead
		}

		if headATailNode != nil && headBTailNode != nil && headATailNode != headBTailNode {
			return nil
		}
	}

	return headA

}
