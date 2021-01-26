package main

import "leetcode/common"

func main() {
	link13 := common.ListNode{3, nil}
	link12 := common.ListNode{4, &link13}
	link11 := common.ListNode{2, &link12}

	link23 := common.ListNode{4, nil}
	link22 := common.ListNode{6, &link23}
	link21 := common.ListNode{5, &link22}

	res := addTwoNumbers(&link11, &link21)
	common.PrintListNode(*res)
}

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	dummy := common.ListNode{}
	cur := &dummy
	carry := 0

	for l1 != nil || l2 != nil {
		var add1, add2 int
		if l1 == nil {
			add1 = 0
		} else {
			add1 = l1.Val
		}
		if l2 == nil {
			add2 = 0
		} else {
			add2 = l2.Val
		}

		// 低位相加
		sum := add1 + add2 + carry

		// 处理进位
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		// 保存数据
		cur.Next = &common.ListNode{sum % 10, nil}
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 处理最高位
	if carry == 1 {
		cur.Next = &common.ListNode{1, nil}
	}

	return dummy.Next
}
