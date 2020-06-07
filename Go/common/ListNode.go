// 数据结构
package common

import (
	"fmt"
)

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListNode 创建链表
func CreateListNode(nodes []int) *ListNode {
	// 声明一个空节点，用于保存已初始化节点
	var dummy *ListNode

	l := len(nodes)
	// 从后往前遍历，最后一个元素就是头结点
	for i := l - 1; i >= 0; i-- {
		// 每次都创建一个新的节点
		var cur = new(ListNode)
		cur.Val = nodes[i]
		cur.Next = dummy // 第一个元素是尾节点指向 nil，其他指向之前创建的节点
		// 暂存当前节点
		dummy = cur
	}

	// dummy 最后保存的是头结点
	return dummy
}

// PrintLIstNode 打印链表
func PrintListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// 双向循环链表
type DoublyList struct {
	Val  int
	Pre  *DoublyList
	Next *DoublyList
}
