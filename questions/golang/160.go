package main

import (
	"fmt"
	"leetcode/common"
)

func main() {
	link14 := common.ListNode{-4, nil}
	link13 := common.ListNode{0, &link14}
	link12 := common.ListNode{2, &link13}
	link11 := common.ListNode{3, &link12}

	link21 := common.ListNode{1, &link12}

	fmt.Println(getIntersectionNode(&link11, &link21).Val)
}

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	// 重置 head 指针
	lenA, lenB := getLength(headA), getLength(headB)
	if lenA >= lenB {
		diff := lenA - lenB
		for diff > 0 {
			headA = headA.Next
			diff--
		}
	} else {
		diff := lenB - lenA
		for diff > 0 {
			headB = headB.Next
			diff--
		}
	}
	// 开始比较
	for {
		if headA == headB {
			return headB
		}
		headA = headA.Next
		headB = headB.Next
		if headA == nil {
			return nil
		}
	}
}

func getLength(head *common.ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
