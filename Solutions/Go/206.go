package main

import "fmt"

func main() {

	tail := ListNode{5, nil}
	next3 := ListNode{4, &tail}
	next2 := ListNode{3, &next3}
	next1 := ListNode{2, &next2}
	head := ListNode{1, &next1}

	printListNode(*reverseList2(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
方法一：迭代
思路：从前往后反转，使用两个指针（pre，cur），每次迭代需要反转和移动指针
*/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		// cur.Next ： 翻转
		// cur,pre ： 移动指针
		cur.Next, cur, pre = pre, cur.Next, cur
	}

	return pre
}

/**
方法二：递归
思路：利用递归从后往前反转
*/
func reverseList2(head *ListNode) *ListNode {
	// 退出条件
	if head == nil || head.Next == nil { // head 用于 [] 边界判断，head.Next 用于递归判断
		return head
	}

	// 下沉
	var node = reverseList2(head.Next) // 最后一个节点

	// 处理逻辑
	head.Next.Next = head // 反转指针
	head.Next = nil       // 断开旧连接

	return node
}

func printListNode(head ListNode) {
	for {
		fmt.Println(head.Val)
		if head.Next == nil {
			break
		}
		head = *head.Next
	}
}
