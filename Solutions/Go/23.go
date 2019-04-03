package main

import "leetcode/Solutions/Go/Common"

func main() {
	link11 := Common.ListNode{4, nil}
	link12 := Common.ListNode{5, nil}
	link11.Next = &link12

	link21 := Common.ListNode{3, nil}
	link22 := Common.ListNode{4, nil}
	link21.Next = &link22

	link31 := Common.ListNode{6, nil}

	lists := []*Common.ListNode{
		{1, &link11},
		{1, &link21},
		{2, &link31},
	}

	res := mergeKLists(lists)
	Common.PrintListNode(*res)
}

func mergeKLists(lists []*Common.ListNode) *Common.ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	for n > 1 {
		k := (n + 1) / 2 // 保证奇数
		//append(lists,mergeTwoLists(lists[]))
		for i := 0; i < n/2; i++ { // 分治
			lists[i] = mergeK2(lists[i], lists[i+k])
		}
		n = k
	}

	return lists[0]
}

func mergeK2(l1, l2 *Common.ListNode) *Common.ListNode {
	dummy := Common.ListNode{}
	cur := &dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next // 别忘记移动指针
	}
	// 处理剩余节点
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}

	return dummy.Next
}
