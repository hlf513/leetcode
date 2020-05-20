package main

import (
	"fmt"

	"leetcode/common"
)

func main() {
	head := common.CreateListNode([]int{1, 2, 4})
	ret := reversePrint(head)
	fmt.Println(ret)
}

/**
解题思路：把链表的数据存入「栈」，数据出栈时就是倒序
时间复杂度：N
空间复杂度：N
*/
func reversePrint(head *common.ListNode) []int {
	if head == nil {
		return []int{}
	}
	// 声明一个切片做栈
	var stack []int
	// 遍历链表，把 val 放入栈
	for head.Next != nil { // 注意到最后一个节点不会进入循环体
		stack = append(stack, head.Val)
		head = head.Next
	}
	// 保存最后一个节点数据
	stack = append(stack, head.Val)

	// 出栈（切片从后向前输出）
	var reverseList []int
	l := len(stack)
	for i := l - 1; i >= 0; i-- {
		reverseList = append(reverseList, stack[i])
	}

	return reverseList
}
